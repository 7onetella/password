package main

import (
	"github.com/rivo/tview"

	"github.com/7onetella/password/api/model"
)

var newForm *tview.Form

func init() {
	newForm = tview.NewForm()
}

// NewSlide returns new password slide
func NewSlide() (title string, content tview.Primitive) {
	notify("loading new slide")

	newForm.AddInputField("Title:", "", 30, nil, nil).
		AddInputField("URL:", "", 60, nil, nil).
		AddInputField("Username:", "", 30, nil, nil).
		AddInputField("Password:", "", 30, nil, nil).
		AddInputField("Notes:", "", 60, nil, nil).
		AddButton("Submit", submitNewForm)

	newForm.SetBorderPadding(1, 1, 1, 1)
	newForm.SetBorder(true)

	return "New", newForm
}

func submitNewForm() {
	title := getInputValue(newForm, "Title:")
	url := getInputValue(newForm, "URL:")
	usernname := getInputValue(newForm, "Username:")
	password := getInputValue(newForm, "Password:")
	notes := getInputValue(newForm, "Notes:")
	notify("title: " + title)

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
		notify("error while creating password: " + err.Error())
	}

	notify("new password created: " + output.Data.ID)

	newForm.Clear(true)

	NewSlide()

	gotoSlide(2)

	app.SetFocus(searchbar)

	app.Draw()
}

var editUUID string
var updateform *tview.Form

// EditSlide returns new password slide
func EditSlide(title, url, username, password, notes string) (string, tview.Primitive) {
	notify("loading new slide")

	f := tview.NewForm().AddInputField("Title:", title, 30, nil, nil).
		AddInputField("URL:", url, 60, nil, nil).
		AddInputField("Username:", username, 30, nil, nil).
		AddInputField("Password:", password, 30, nil, nil).
		AddInputField("Notes:", notes, 60, nil, nil).
		AddButton("Update", submitUpdateForm).
		AddButton("Delete", deleteAction)

	f.SetBorderPadding(1, 1, 1, 1)
	f.SetBorder(true)
	updateform = f
	return "Edit", f
}

func submitUpdateForm() {
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
		notify("error while updating password: " + err.Error())
	}

	newForm.Clear(true)

	NewSlide()

	gotoSlide(2)

	app.SetFocus(searchbar)

	app.Draw()
}

func deleteAction() {

	err := svc.DeletePassword(editUUID)
	if err != nil {
		notify("error while updating password: " + err.Error())
	}

	newForm.Clear(true)

	NewSlide()

	gotoSlide(2)

	app.SetFocus(searchbar)

	app.Draw()

}
