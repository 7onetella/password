package main

import (
	"github.com/rivo/tview"

	"github.com/7onetella/password/api/model"
)

var editUUID string
var updateform *tview.Form

func editPage(title, url, username, password, notes string) (string, tview.Primitive) {
	debug("loading edit page")

	f := tview.NewForm().AddInputField("Title:", title, 60, nil, nil).
		AddInputField("URL:", url, 60, nil, nil).
		AddInputField("Username:", username, 60, nil, nil).
		AddInputField("Password:", password, 60, nil, nil).
		AddInputField("Notes:", notes, 60, nil, nil).
		AddButton("Update", updateAction).
		AddButton("Delete", deleteAction).
		AddButton("Cancel", cancelAction)

	f.SetBorderPadding(1, 1, 2, 2)
	f.SetBorder(true)
	updateform = f
	return "Edit", f
}

func updateAction() {
	title := getInputValue(updateform, "Title:")
	url := getInputValue(updateform, "URL:")
	usernname := getInputValue(updateform, "Username:")
	password := getInputValue(updateform, "Password:")
	notes := getInputValue(updateform, "Notes:")

	input := model.PasswordInput{
		Data: model.Password{
			ID:       editUUID,
			Title:    title,
			URL:      url,
			Username: usernname,
			Password: password,
			Notes:    notes,
		},
	}

	err := svc.UpdatePassword(input)
	if err != nil {
		debug("error while updating password: " + err.Error())
	}

	gotoPage(pageSearch)
	showSearchResults()
	app.SetFocus(searchBar)
	editUUID = ""

	app.Draw()
}

func deleteAction() {

	err := svc.DeletePassword(editUUID)
	if err != nil {
		debug("error while updating password: " + err.Error())
	}

	gotoPage(pageSearch)
	showSearchResults()
	app.SetFocus(searchBar)
	editUUID = ""

	app.Draw()

}

func cancelAction() {

	gotoPage(pageSearch)
	showSearchResults()
	app.SetFocus(searchBar)
	editUUID = ""

	app.Draw()

}
