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

// CreatePassword creates password
func (ps *PasswordService) CreatePassword(p model.Password) (*model.CreatePasswordOutput, error) {

	url := fmt.Sprintf("http://%s/api/passwords", ps.serverAddr)

	response := &model.CreatePasswordOutput{}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	data, err := doPost(url, b)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ReadPassword creates password
func (ps *PasswordService) ReadPassword(ID string) (*model.Password, error) {

	url := fmt.Sprintf("http://%s/api/passwords/%s", ps.serverAddr, ID)

	response := &model.Password{}

	data, err := doGet(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdatePassword updates password
func (ps *PasswordService) UpdatePassword(p model.Password) error {

	url := fmt.Sprintf("http://%s/api/passwords", ps.serverAddr)

	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = doPatch(url, b)
	if err != nil {
		return err
	}

	return nil
}

// DeletePassword deletes password
func (ps *PasswordService) DeletePassword(ID string) error {

	url := fmt.Sprintf("http://%s/api/passwords/%s", ps.serverAddr, ID)

	_, err := doDelete(url)
	if err != nil {
		return err
	}

	return nil
}

// ListPasswords finds passwords by title
func (ps *PasswordService) ListPasswords(input model.ListPasswordsInput) (*model.ListPasswordsOutput, error) {

	url := fmt.Sprintf("http://%s/api/list?title=%s", ps.serverAddr, input.Title)

	response := &model.ListPasswordsOutput{}

	data, err := doGet(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func doGet(url string) ([]byte, error) {
	return doHTTPAction("GET", url, nil)
}

func doPost(url string, b []byte) ([]byte, error) {
	r := bytes.NewBuffer(b)
	return doHTTPAction("POST", url, r)
}

func doPut(url string, b []byte) ([]byte, error) {
	r := bytes.NewBuffer(b)
	return doHTTPAction("PUT", url, r)
}

func doPatch(url string, b []byte) ([]byte, error) {
	r := bytes.NewBuffer(b)
	return doHTTPAction("PATCH", url, r)
}

func doDelete(url string) ([]byte, error) {
	return doHTTPAction("DELETE", url, nil)
}

func doHTTPAction(action, url string, r io.Reader) ([]byte, error) {
	req, err := http.NewRequest(action, url, r)
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

func unmarshalPassword(data []byte) (*model.Password, error) {
	password := &model.Password{}
	err := json.Unmarshal(data, password)
	if err != nil {
		return nil, err
	}
	return password, nil
}
