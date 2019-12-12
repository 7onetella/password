package main

import (
	"github.com/rivo/tview"

	"github.com/7onetella/password/api/model"
)

var newPageForm = tview.NewForm()

// NewPage returns new page
func NewPage() (title string, content tview.Primitive) {
	newPageReset()
	return "New", newPageForm
}

func newPageReset() {
	newPageForm.Clear(true)
	newPageForm.AddInputField("Title:", "", 30, nil, nil).
		AddInputField("URL:", "", 60, nil, nil).
		AddInputField("Username:", "", 30, nil, nil).
		AddInputField("Password:", "", 30, nil, nil).
		AddInputField("Notes:", "", 60, nil, nil).
		AddButton("Save", saveAction)
	newPageForm.SetBorderPadding(1, 1, 2, 1).SetBorder(true)
}

func saveAction() {
	title := getInputValue(newPageForm, "Title:")
	url := getInputValue(newPageForm, "URL:")
	usernname := getInputValue(newPageForm, "Username:")
	password := getInputValue(newPageForm, "Password:")
	notes := getInputValue(newPageForm, "Notes:")

	input := model.PasswordInput{
		Data: model.Password{
			Title:    title,
			URL:      url,
			Username: usernname,
			Password: password,
			Notes:    notes,
		},
	}

	output, err := svc.CreatePassword(input)
	if err != nil {
		debug("error while creating password: " + err.Error())
	}

	debug("new password created: " + output.Data.ID)

	newPageReset()

	gotoPage(pageSearch)
	showSearchResults()
	app.SetFocus(searchBar)

	app.Draw()
}
