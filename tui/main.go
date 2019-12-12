package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/7onetella/password/api/client"
	"github.com/7onetella/password/api/model"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Slide is a function which returns the slide's main primitive and its title.
type Slide func() (title string, content tview.Primitive)

var pages *tview.Pages

var menubar *tview.TextView

var app = tview.NewApplication()

var debugView *tview.Flex
var notification *tview.Table

var currentSlide = 0

var slides []Slide

var rows *tview.Flex

var svc *client.PasswordService

var credentials model.Credentials

var isDebugOn = true

var prevP tview.Primitive

func init() {

	pages = tview.NewPages()

	menubar = newMenuBar()

	notification = newtable()

	slides = signedOutSlides()

}

func newDebugBox() *tview.TextView {
	box := tview.NewTextView().SetWordWrap(true).SetChangedFunc(func() {
		app.Draw()
	})
	box.SetBorderPadding(0, 0, 0, 0)
	box.SetBorder(true).SetTitle("Debug")
	return box
}

func signedInSlides() []Slide {
	return []Slide{
		homePage,
		newPage,
		searchPage,
		signOutPage,
	}
}

func signedOutSlides() []Slide {
	return []Slide{
		homePage,
		aboutPage,
		signInPage,
	}
}

func debug(message string) {
	if isDebugOn {
		// lastRow := notification.GetRowCount()
		lastRow := 0
		notification.InsertRow(lastRow)
		notification.SetCell(lastRow, 0, &tview.TableCell{
			Text:            time.Now().Format(time.RFC3339),
			Align:           tview.AlignLeft,
			Color:           tcell.ColorDarkCyan,
			BackgroundColor: tcell.ColorDefault,
		})
		notification.SetCellSimple(lastRow, 1, message)
		app.Draw()
		notification.ScrollToEnd()
	}
}

func gotoPage(index int) {
	currentSlide = index
	menubar.Highlight(strconv.Itoa(index)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(index))
}

func previousPage() {
	currentSlide = (currentSlide - 1 + len(slides)) % len(slides)
	menubar.Highlight(strconv.Itoa(currentSlide)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(currentSlide))
}

func nextPage() {
	currentSlide = (currentSlide + 1) % len(slides)
	menubar.Highlight(strconv.Itoa(currentSlide)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(currentSlide))
}

func main() {

	flex := tview.NewFlex()

	menubar.Highlight(strconv.Itoa(currentSlide))

	loadPages(slides)

	rows = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(menubar, 1, 1, false).
		AddItem(pages, 0, 9, true)

	flex.AddItem(rows, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyCtrlL:
			nextPage()
		case tcell.KeyCtrlH:
			previousPage()
		case tcell.KeyCtrlO:
			if isDebugOn {
				isDebugOn = false
				rows.RemoveItem(debugView)
			} else {
				isDebugOn = true
				if debugView == nil {
					debugView = tview.NewFlex().AddItem(notification, 0, 1, false)
					debugView.SetBorder(true).SetBorderPadding(0, 0, 0, 0)
					debugView.SetTitle("Debug")
				}
				rows.AddItem(debugView, 10, 2, false)
			}
			app.Draw()
		case tcell.KeyCtrlD:
			if !notification.HasFocus() {
				prevP = app.GetFocus()
				app.SetFocus(notification)
			} else {
				app.SetFocus(prevP)
			}
			app.Draw()
		default:
			// do nothing
		}

		return event
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

}

func clearMenu() {
	menubar.Clear()
	app.Draw()
}

var pageIndex = map[string]string{}

func loadPages(newslides []Slide) {
	for index, slide := range newslides {
		title, primitive := slide()
		fmt.Fprintf(menubar, `["%d"][darkcyan]%s[white][""]  `, index, title)
		indexStr := strconv.Itoa(index)
		pages.AddPage(indexStr, primitive, true, index == currentSlide)
		pageIndex[title] = indexStr
	}
}

func unloadPages() {
	for index := range slides {
		if index == 0 {
			continue
		}
		pages.RemovePage(strconv.Itoa(index))
	}
}

func newMenuBar() *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)
}

func newtable() *tview.Table {
	table := tview.NewTable().
		SetBorders(false).
		InsertColumn(0).
		InsertRow(0).
		InsertColumn(0).
		InsertColumn(0).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.ColorBlack, tcell.ColorWhite, tcell.AttrNone)
	return table
}

func homePage() (title string, content tview.Primitive) {
	homeView := tview.NewTextView().SetWordWrap(true)
	fmt.Fprint(homeView, `    Hello, my dear wife, love of my life!
	I created this app for you to manage your passwords.
	
	I love you.	`)
	homeView.SetBorder(true)
	homeView.SetBorderPadding(4, 0, 4, 0)
	return "Home", homeView
}

func aboutPage() (title string, content tview.Primitive) {
	aboutView := tview.NewTextView().SetWordWrap(true)
	fmt.Fprint(aboutView, `    Motivation

	It's hard to remember so many accounts and their login credentials. 
	I decided to learn rivo/tview and create a tool in the process.
	`)
	aboutView.SetBorder(true)
	aboutView.SetBorderPadding(4, 0, 4, 0)
	return "About", aboutView
}
