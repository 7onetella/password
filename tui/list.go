package main

import (
	"fmt"
	"strconv"

	"github.com/7onetella/password/api/model"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var passwordsTable *tview.Table
var searchbar *tview.Form

// NewPasswordSlide returns new password slide
func NewPasswordSlide() (title string, content tview.Primitive) {
	notify("loading password slide")

	flex := tview.NewFlex()
	searchbar = tview.NewForm()
	updatesearchbar(searchbar)

	passwordsTable = tview.NewTable()
	passwordsTable.SetBorders(true).SetTitle("Results")

	rows := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(searchbar, 0, 3, true).
		AddItem(passwordsTable, 0, 10, true)
	flex.AddItem(rows, 0, 1, true)

	flex.SetBorder(true).SetBorderPadding(1, 1, 2, 2)

	return "Password", flex
}

func updatesearchbar(f *tview.Form) {
	f.AddInputField("Title:", "", 0, nil, nil)
	f.SetBorderPadding(0, 0, 0, 0)
	item := f.GetFormItemByLabel("Title:")
	textField, ok := item.(*tview.InputField)

	if ok {
		textField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEnter {
				notify("enter key pressed")
				showPasswords()
				defer app.SetFocus(passwordsTable)
				defer app.Draw()
				return nil
			}
			return event
		})
	}
}

func showPasswords() {
	notify("table received focus")

	input := model.ListPasswordsInput{}
	searchby := getInputValue(searchbar, "Title:")
	input.Title = "%%"
	if len(searchby) > 0 {
		input.Title = searchby
	}
	notify("search by = " + searchby)

	result, err := svc.ListPasswords(input)
	if err != nil {
		notify("error: " + err.Error())
		return
	}
	if result == nil {
		notify("result is empty")
		return
	}

	notify("result size is " + strconv.Itoa(len(result.Items)))
	passwordsTable.Clear()
	passwordsTable.InsertColumn(0)
	passwordsTable.InsertColumn(0)
	passwordsTable.InsertColumn(0)

	for _, password := range result.Items {
		passwordsTable.InsertRow(0)
		passwordsTable.SetCell(0, 0, &tview.TableCell{
			Reference:       password.ID,
			Text:            password.Title,
			Align:           tview.AlignLeft,
			Color:           tview.Styles.PrimaryTextColor,
			BackgroundColor: tcell.ColorDefault,
		})
		// passwordsTable.SetCellSimple(0, 0, password.Title)
		passwordsTable.SetCellSimple(0, 1, password.Username)
		passwordsTable.SetCellSimple(0, 2, password.URL)
	}

	passwordsTable.InsertRow(0)
	passwordsTable.SetCellSimple(0, 0, "Title")
	passwordsTable.SetCellSimple(0, 1, "Username")
	passwordsTable.SetCellSimple(0, 2, "URL")
	passwordsTable.SetSelectable(true, false)

	passwordsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			notify("tab key pressed")
			searchbar.Clear(true)
			updatesearchbar(searchbar)
			app.Draw()
			app.SetFocus(searchbar)
			return nil
		}

		if event.Key() == tcell.KeyEnter {
			notify("enter key pressed")
			row, _ := passwordsTable.GetSelection()
			notify(fmt.Sprintf("row = %d", row))

			ref := passwordsTable.GetCell(row, 0).GetReference()
			id, _ := ref.(string)
			col1 := passwordsTable.GetCell(row, 0).Text
			col2 := passwordsTable.GetCell(row, 1).Text
			col3 := passwordsTable.GetCell(row, 2).Text

			notify("id = " + id)
			notify("col1 = " + col1)
			notify("col2 = " + col2)
			notify("col3 = " + col3)

			pi, err := svc.ReadPassword(id)
			if err != nil {
				notify("error while reading: " + err.Error())
			}
			d := pi.Data
			editUUID = d.ID
			_, p := EditSlide(d.Title, d.URL, d.Username, d.Password, d.Notes)
			// fmt.Fprintf(menubar, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
			pages.AddPage("Edit", p, true, true)

			return nil
		}

		return event
	})

	app.SetFocus(passwordsTable)
	app.Draw()
}
