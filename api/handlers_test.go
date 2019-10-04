package main

import (
	"testing"

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

	spec.When("svc.CreatePassword")

	response, err := svc.CreatePassword(p)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)

	spec.AssertAndFailNow(len(response.ID) > 0, "result should return non-empty id", len(response.ID))
}

func TestReadPasswordRequest(t *testing.T) {

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	expected := newPassword()
	createResponse, err := svc.CreatePassword(expected)
	ID := createResponse.ID

	spec.Given("id=" + ID)

	spec.When("svc.ReadPassword(ID)")

	password, err := svc.ReadPassword(ID)

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

	p := newPassword()
	createResponse, err := svc.CreatePassword(p)
	p.ID = createResponse.ID

	spec.Given(prettyJSON(p), "username="+p.Username)

	spec.When("svc.UpdatePassword")

	err = svc.UpdatePassword(p)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)
}

func TestDeletePasswordRequest(t *testing.T) {

	spec := GSpec{t}

	svc, err := client.NewPasswordService()
	if err != nil {
		t.Errorf("creating serivce failed: %v", err)
	}

	expected := newPassword()
	createResponse, err := svc.CreatePassword(expected)
	ID := createResponse.ID

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

	expected := newPassword()
	createResponse, err := svc.CreatePassword(expected)
	ID := createResponse.ID

	spec.Given("title=" + expected.Title)

	spec.When("svc.ListPasswords")

	input := model.ListPasswordsInput{
		Title: expected.Title,
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
