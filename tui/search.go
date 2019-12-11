package main

import (
	"fmt"
	"strconv"

	"github.com/7onetella/password/api/model"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var passwordsTable *tview.Table
var searchBar *tview.Form

func searchPage() (title string, content tview.Primitive) {
	flex := tview.NewFlex()
	searchBar = tview.NewForm()
	searchBarReset(searchBar)

	passwordsTable = tview.NewTable()
	passwordsTable.
		SetBorders(false).
		SetSeparator(' ').
		SetTitle("Results")
	passwordsTable.SetBorder(true)
	passwordsTable.SetBorderPadding(1, 1, 2, 2)

	rows := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(searchBar, 2, 1, true).
		AddItem(passwordsTable, 0, 10, true)

	flex.AddItem(rows, 0, 1, true)

	flex.SetBorder(true).SetBorderPadding(1, 1, 2, 2)

	return "Password", flex
}

func searchBarReset(f *tview.Form) {
	f.AddInputField("Title:", "", 0, nil, nil)
	f.SetBorderPadding(0, 1, 0, 0)
	item := f.GetFormItemByLabel("Title:")
	textField, ok := item.(*tview.InputField)

	if ok {
		textField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyEnter {
				showSearchResults()
				defer app.SetFocus(passwordsTable)
				defer app.Draw()
				return nil
			}
			return event
		})
	}
}

func showSearchResults() {

	input := model.ListPasswordsInput{}
	searchby := getInputValue(searchBar, "Title:")
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

	headerCell := func(val string) *tview.TableCell {
		return &tview.TableCell{
			Color:         tcell.ColorYellow,
			Align:         tview.AlignCenter,
			Text:          val,
			NotSelectable: true,
		}
	}

	passwordsTable.InsertRow(0)
	passwordsTable.SetCell(0, 0, headerCell("Title"))
	passwordsTable.SetCell(0, 1, headerCell("Username"))
	passwordsTable.SetCell(0, 2, headerCell("URL"))

	passwordsTable.SetSelectable(true, false)
	// passwordsTable.SetSeparator(tview.Borders.Vertical)
	passwordsTable.SetSeparator(' ')
	passwordsTable.SetSelectedStyle(tcell.ColorBlack, tcell.ColorGray, tcell.AttrNone)

	passwordsTable.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			notify("tab key pressed")
			searchBar.Clear(true)
			searchBarReset(searchBar)
			app.Draw()
			app.SetFocus(searchBar)
			return nil
		}

		if event.Key() == tcell.KeyEnter {
			notify("enter key pressed")
			row, _ := passwordsTable.GetSelection()
			notify(fmt.Sprintf("row = %d", row))

			ref := passwordsTable.GetCell(row, 0).GetReference()
			id, _ := ref.(string)
			notify("id = " + id)
			pi, err := svc.ReadPassword(id)
			if err != nil {
				notify("error while reading: " + err.Error())
			}
			d := pi.Data
			editUUID = d.ID
			_, p := editPage(d.Title, d.URL, d.Username, d.Password, d.Notes)
			pages.AddPage("Edit", p, true, true)

			return nil
		}

		return event
	})

	app.SetFocus(passwordsTable)
	app.Draw()
}
