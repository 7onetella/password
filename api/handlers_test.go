package main

import (
	"testing"
	"time"

	"github.com/7onetella/password/api/client"
	"github.com/7onetella/password/api/model"
)

func TestCreatePasswordRequest(t *testing.T) {

	spec := GSpec{t}

	p := newPassword()

	spec.Given(prettyJSON(p), "username="+p.Username)

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	err = svc.Signin(model.Credentials{Username: "admin", Password: "password"})
	if err != nil {
		t.Errorf("authenticating failed: %v", err)
		return
	}

	spec.When("svc.CreatePassword")

	response, err := svc.CreatePassword(model.PasswordInput{Data: p})

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)

	spec.AssertAndFailNow(len(response.Data.ID) > 0, "result should return non-empty id", len(response.Data.ID))
}

func TestReadPasswordRequest(t *testing.T) {

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	err = svc.Signin(model.Credentials{Username: "admin", Password: "password"})
	if err != nil {
		t.Errorf("authenticating failed: %v", err)
		return
	}

	expected := newPassword()
	createResponse, err := svc.CreatePassword(model.PasswordInput{Data: expected})
	ID := createResponse.Data.ID

	spec.Given("id=" + ID)

	spec.When("svc.ReadPassword(ID)")

	o, err := svc.ReadPassword(ID)
	password := o.Data

	spec.Then()

	spec.AssertAndFailNow(err == nil, "calling the endpoint should not return error", err)

	spec.AssertAndFailNow(len(password.ID) > 0, "result should return non-empty id", len(password.ID))
	spec.AssertAndFailNow(password.URL == "http://www.example.com", "url should be http://www.example.com", password.URL)
	spec.AssertAndFailNow(password.Username == expected.Username, "username should be "+expected.Username, password.Username)
	spec.AssertAndFailNow(password.Password == expected.Password, "password should be "+expected.Password, password.Password)
	spec.AssertAndFailNow(password.Notes == expected.Notes, "notes should be "+expected.Notes, password.Notes)
	spec.AssertAndFailNow(password.Tags[0] == expected.Tags[0], "first tag should be "+expected.Tags[0], password.Tags[0])
	spec.AssertAndFailNow(password.Tags[1] == expected.Tags[1], "second tag should be "+expected.Tags[1], password.Tags[1])

}

func TestUpdatePasswordRequest(t *testing.T) {

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	err = svc.Signin(model.Credentials{Username: "admin", Password: "password"})
	if err != nil {
		t.Errorf("authenticating failed: %v", err)
		return
	}

	p := newPassword()
	input := model.PasswordInput{Data: p}
	output, err := svc.CreatePassword(input)
	p.ID = output.Data.ID

	spec.Given(prettyJSON(p), "username="+p.Username)

	spec.When("svc.UpdatePassword")

	err = svc.UpdatePassword(model.PasswordInput{Data: p})

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)
}

func TestDeletePasswordRequest(t *testing.T) {

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	err = svc.Signin(model.Credentials{Username: "admin", Password: "password"})
	if err != nil {
		t.Errorf("authenticating failed: %v", err)
		return
	}

	expected := newPassword()
	createResponse, err := svc.CreatePassword(model.PasswordInput{Data: expected})
	ID := createResponse.Data.ID

	spec.Given("ID=" + ID)

	spec.When("svc.DeletePassword")

	err = svc.DeletePassword(ID)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)
}

func TestListPasswordsRequest(t *testing.T) {
	DeleteAllPasswords()

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	err = svc.Signin(model.Credentials{Username: "admin", Password: "password"})
	if err != nil {
		t.Errorf("authenticating failed: %v", err)
		return
	}

	expected := newPassword()
	createResponse, err := svc.CreatePassword(model.PasswordInput{Data: expected})
	ID := createResponse.Data.ID

	spec.Given("title=" + expected.Title)

	spec.When("svc.ListPasswords")

	input := model.ListPasswordsInput{
		Title:   expected.Title,
		AdminID: expected.AdminID,
	}

	response, err := svc.ListPasswords(input)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)
	spec.AssertAndFailNow(len(response.Items) > 0, "result should return 1 item", len(response.Items))

	item := response.Items[0]
	spec.AssertAndFailNow(item.ID == ID, "id should be "+ID, item.ID)
	spec.AssertAndFailNow(item.URL == expected.URL, "url should be "+expected.URL, item.URL)
	spec.AssertAndFailNow(item.Username == expected.Username, "username should be "+expected.Username, item.Username)
	spec.AssertAndFailNow(item.Password == expected.Password, "password should be "+expected.Password, item.Password)
	spec.AssertAndFailNow(item.Notes == expected.Notes, "notes should be "+expected.Notes, item.Notes)
	spec.AssertAndFailNow(item.Tags[0] == expected.Tags[0], "first tag should be "+expected.Tags[0], item.Tags[0])
	spec.AssertAndFailNow(item.Tags[1] == expected.Tags[1], "second tag should be "+expected.Tags[1], item.Tags[1])

}

func TestRefreshToken(t *testing.T) {

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	err = svc.Signin(model.Credentials{Username: "admin", Password: "password"})
	if err != nil {
		t.Errorf("authenticating failed: %v", err)
		return
	}

	authToken := svc.Token

	spec.Given("Auth Token=" + svc.Token)

	spec.When("svc.RefreshToken()")

	time.Sleep(5 * time.Second)

	err = svc.RefreshToken()

	spec.Then()

	refreshToken := svc.Token

	spec.AssertAndFailNow(err == nil, "result should not return error", err)
	spec.AssertAndFailNow(len(authToken) > 0, "length of auth token is not 0", len(authToken))
	spec.AssertAndFailNow(len(refreshToken) > 0, "length of refresh token is not 0", len(refreshToken))
	spec.AssertAndFailNow(authToken != refreshToken, "authToken != refreshToken", "authToken == refreshToken")

}
