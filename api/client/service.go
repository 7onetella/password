package client

import (
	"bytes"
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
	serverAddr string
}

// NewPasswordService returns new instance of password service
func NewPasswordService() (*PasswordService, error) {
	serverAddr, exists := os.LookupEnv("SERVER_ADDR")
	if !exists {
		return nil, errors.New("SERVER_ADDR environment variable not set")
	}

	return &PasswordService{serverAddr: serverAddr}, nil
}

// RestfulService restful service
type RestfulService interface {
	GetEndpoint() string
}

// GetEndpoint returns endpoint base url
func (ps *PasswordService) GetEndpoint() string {
	return fmt.Sprintf("http://%s/api/passwords", ps.serverAddr)
}

// CallEndpoint calls endpoint
func CallEndpoint(url, method string, v interface{}, o interface{}) (interface{}, error) {
	data, err := httpAction(url, method, v)
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

// CreatePassword creates password
func (ps *PasswordService) CreatePassword(p model.Password) (*model.CreatePasswordOutput, error) {
	o, err := CallEndpoint(ps.GetEndpoint(), "POST", &p, &model.CreatePasswordOutput{})
	if err != nil {
		return nil, err
	}

	return o.(*model.CreatePasswordOutput), nil
}

// ReadPassword creates password
func (ps *PasswordService) ReadPassword(ID string) (*model.Password, error) {

	o, err := CallEndpoint(ps.GetEndpoint()+"/"+ID, "GET", nil, &model.Password{})
	if err != nil {
		return nil, err
	}

	return o.(*model.Password), nil

}

// UpdatePassword updates password
func (ps *PasswordService) UpdatePassword(p model.Password) error {

	_, err := CallEndpoint(ps.GetEndpoint(), "PATCH", &p, nil)
	if err != nil {
		return err
	}

	return nil

}

// DeletePassword deletes password
func (ps *PasswordService) DeletePassword(ID string) error {

	_, err := CallEndpoint(ps.GetEndpoint()+"/"+ID, "DELETE", nil, nil)
	if err != nil {
		return err
	}

	return nil

}

// ListPasswords finds passwords by title
func (ps *PasswordService) ListPasswords(input model.ListPasswordsInput) (*model.ListPasswordsOutput, error) {

	o, err := CallEndpoint(ps.GetEndpoint()+"?title="+input.Title, "GET", nil, &model.ListPasswordsOutput{})
	if err != nil {
		return nil, err
	}

	return o.(*model.ListPasswordsOutput), nil

}

func httpAction(url, method string, v interface{}) ([]byte, error) {
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
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
