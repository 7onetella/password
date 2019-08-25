package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/7onetella/password/api/model"
)

// PasswordService password service
type PasswordService struct {
	serverAddr         string
	Authorization      string
	InsecureSkipVerify bool
}

// NewPasswordService returns new instance of password service
func NewPasswordService() (*PasswordService, error) {
	ps := &PasswordService{}

	serverAddr, exists := os.LookupEnv("SERVER_ADDR")
	if !exists {
		return nil, errors.New("SERVER_ADDR environment variable not set")
	}
	ps.serverAddr = serverAddr

	insecure := os.Getenv("INSECURE")
	if insecure == "true" {
		ps.InsecureSkipVerify = true
	}

	return ps, nil
}

// RestfulService restful service
type RestfulService interface {
	GetEndpoint() string
}

// GetEndpoint returns endpoint base url
func (ps *PasswordService) GetEndpoint() string {
	return fmt.Sprintf("https://%s/api/passwords", ps.serverAddr)
}

// CallEndpoint calls endpoint
func CallEndpoint(url, method, authorization string, insecureSkipVerify bool, v interface{}, o interface{}) (interface{}, error) {
	data, err := httpAction(url, method, authorization, insecureSkipVerify, v)
	if err != nil {
		return nil, err
	}

	if o != nil {
		err = json.Unmarshal(data, o)
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}

// Signin authenticates
func (ps *PasswordService) Signin(input model.Credentials) error {
	o, err := CallEndpoint("https://"+ps.serverAddr+"/api/signin", "POST", "", ps.InsecureSkipVerify, &input, &model.AuthToken{})
	if err != nil {
		return err
	}

	authToken := o.(*model.AuthToken)
	ps.Authorization = "Bearer " + authToken.Token

	return nil
}

// CreatePassword creates password
func (ps *PasswordService) CreatePassword(input model.PasswordInput) (*model.PasswordOutput, error) {
	o, err := CallEndpoint(ps.GetEndpoint(), "POST", ps.Authorization, ps.InsecureSkipVerify, &input, &model.PasswordOutput{})
	if err != nil {
		return nil, err
	}

	return o.(*model.PasswordOutput), nil
}

// ReadPassword creates password
func (ps *PasswordService) ReadPassword(ID string) (*model.PasswordOutput, error) {

	o, err := CallEndpoint(ps.GetEndpoint()+"/"+ID, "GET", ps.Authorization, ps.InsecureSkipVerify, nil, &model.PasswordOutput{})
	if err != nil {
		return nil, err
	}

	return o.(*model.PasswordOutput), nil

}

// UpdatePassword updates password
func (ps *PasswordService) UpdatePassword(input model.PasswordInput) error {

	_, err := CallEndpoint(ps.GetEndpoint()+"/"+input.Data.ID, "PATCH", ps.Authorization, ps.InsecureSkipVerify, &input, nil)
	if err != nil {
		return err
	}

	return nil

}

// DeletePassword deletes password
func (ps *PasswordService) DeletePassword(ID string) error {

	_, err := CallEndpoint(ps.GetEndpoint()+"/"+ID, "DELETE", ps.Authorization, ps.InsecureSkipVerify, nil, nil)
	if err != nil {
		return err
	}

	return nil

}

// ListPasswords finds passwords by title
func (ps *PasswordService) ListPasswords(input model.ListPasswordsInput) (*model.ListPasswordsOutput, error) {

	o, err := CallEndpoint(ps.GetEndpoint()+"?title="+input.Title+"&admin_id="+input.AdminID, "GET", ps.Authorization, ps.InsecureSkipVerify, nil, &model.ListPasswordsOutput{})
	if err != nil {
		return nil, err
	}

	return o.(*model.ListPasswordsOutput), nil

}

func httpAction(url, method string, authorization string, insecureSkipVerify bool, v interface{}) ([]byte, error) {
	var r io.Reader
	if v != nil {
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", authorization)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 300 {
		return nil, errors.New("api returned status code " + res.Status)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
