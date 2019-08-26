package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/7onetella/password/api/model"
	"github.com/google/uuid"
)

const jsonprefix = "                "

func TestCreatePassword(t *testing.T) {
	spec := GSpec{t}

	p := newPassword()

	spec.Given(prettyJSON(p))

	spec.When("CreatePassword(p)")

	id, err := CreatePassword(p)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)

	spec.AssertAndFailNow(len(id) > 0, "id should not be empty", len(id))
}

func TestReadPassword(t *testing.T) {
	spec := GSpec{t}

	e := newPassword()
	ID, err := CreatePassword(e)

	spec.Given("ID=" + ID)

	spec.When("ReadPassword()")

	a, err := ReadPassword(ID)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)

	spec.AssertAndFailNow(ID == a.ID, "ID should be "+ID, a.ID)

	spec.AssertAndFailNow(e.Title == a.Title, "Username should be "+e.Title, a.Title)

	spec.AssertAndFailNow(e.Username == a.Username, "Username should be "+e.Username, a.Username)

	spec.AssertAndFailNow(e.URL == a.URL, "URL should be "+e.URL, a.URL)

	spec.AssertAndFailNow(e.Password == a.Password, "Password should be "+e.Password, a.Password)

	spec.AssertAndFailNow(e.Notes == a.Notes, "Notes should be "+e.Notes, a.Notes)

	spec.AssertAndFailNow(e.Tags[0] == a.Tags[0], "First tag should be "+e.Tags[0], a.Tags[0])

	spec.AssertAndFailNow(e.Tags[1] == a.Tags[1], "Second tag should be "+e.Tags[1], a.Tags[1])
}

func TestUpdatePassword(t *testing.T) {
	spec := GSpec{t}

	e := newPassword()
	ID, _ := CreatePassword(e)

	e.ID = ID
	e.Title = "new title"
	e.URL = "new URL"
	e.Username = "new username"
	e.Password = "new password"
	e.Notes = "new notes"
	e.Tags[0] = "new first tag"
	e.Tags[1] = "new second tag"

	spec.Given(prettyJSON(e))

	spec.When("UpdatePassword(e)")

	err := UpdatePassword(e)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return err", err)

	a, err := ReadPassword(ID)

	spec.AssertAndFailNow(ID == a.ID, "ID should be "+ID, a.ID)

	spec.AssertAndFailNow(e.Title == a.Title, "Title should be "+e.Title, a.Title)

	spec.AssertAndFailNow(e.Username == a.Username, "Username should be "+e.Username, a.Username)

	spec.AssertAndFailNow(e.URL == a.URL, "URL should be "+e.URL, a.URL)

	spec.AssertAndFailNow(e.Password == a.Password, "Password should be "+e.Password, a.Password)

	spec.AssertAndFailNow(e.Notes == a.Notes, "Notes should be "+e.Notes, a.Notes)

	spec.AssertAndFailNow(e.Tags[0] == a.Tags[0], "First tag should be "+e.Tags[0], a.Tags[0])

	spec.AssertAndFailNow(e.Tags[1] == a.Tags[1], "Second tag should be "+e.Tags[1], a.Tags[1])
}

func TestDeletePassword(t *testing.T) {
	spec := GSpec{t}

	p := newPassword()
	ID, _ := CreatePassword(p)

	spec.Given(prettyJSON(p))

	spec.When("DeletePassword(p)")

	err := DeletePassword(ID)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return err", err)

	_, err = ReadPassword(ID)

	spec.AssertAndFailNow(err != nil, "read should return error after delete", err)

}

func TestListAllPasswords(t *testing.T) {

	spec := GSpec{t}

	token := ""
	pageSize := 10

	spec.Given("token = ''"+token, fmt.Sprintf("size  = %d", pageSize))

	spec.When("ListAllPasswords('', 10)")

	passwords, err := ListAllPasswords("", 10)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)

	spec.AssertAndFailNow(len(passwords) > 0, "result should return non-zero count", len(passwords))

}

func TestFindPasswordsByTitle(t *testing.T) {
	spec := GSpec{t}

	p := newPassword()
	p.Title = "foo zulu bar"
	_, err := CreatePassword(p)

	spec.Given(prettyJSON(p), "rid = ''", "nextToken = ''", "size = 10")

	spec.When("FindPasswordsByUsername(p)")

	passwords, err := FindPasswordsByTitle("", "%zulu%", "", 10)

	spec.Then()

	spec.AssertAndFailNow(err == nil, "result should not return error", err)

	spec.AssertAndFailNow(len(passwords) > 1, "result should return at least 1 record", len(passwords))
}

func newPassword() model.Password {
	p := model.Password{
		ID:       "",
		Title:    "title1",
		URL:      "http://www.example.com",
		Username: "user-" + uuid.New().String(),
		Password: "password",
		Notes:    "Lorem ipsum dolor sit amet, has fabulas percipit consequat id",
		Tags:     []string{"bank", "bank of mars"},
	}

	return p
}

func prettyJSON(p model.Password) string {
	data, _ := json.MarshalIndent(&p, jsonprefix, "    ")
	return string(data)
}
