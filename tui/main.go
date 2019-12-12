package main

import (
	"fmt"
	"strconv"

	"github.com/7onetella/password/api/client"
	"github.com/7onetella/password/api/model"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Slide is a function which returns the slide's main primitive and its title.
type Slide func() (title string, content tview.Primitive)

var pages = tview.NewPages()

var menubar = newMenuBar()

var app = tview.NewApplication()

var currentSlide = 0

var slides []Slide

var rows *tview.Flex

var svc *client.PasswordService

var credentials model.Credentials

var isDebugOn = false

var prevP tview.Primitive

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

func gotoPage(pageTitle string) {
	index, _ := strconv.Atoi(pageIndex[pageTitle])
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

	slides = signedOutSlides()
	
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
				notification.SetSelectable(true, false)
			} else {
				notification.SetSelectable(false, false)
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

var pageHome = "Home"

func homePage() (title string, content tview.Primitive) {
	homeView := tview.NewTextView().SetWordWrap(true)
	fmt.Fprint(homeView, `    Hello, my dear wife, love of my life!
	I created this app for you to manage your passwords.
	
	I love you.	`)
	homeView.SetBorder(true)
	homeView.SetBorderPadding(4, 0, 4, 0)
	return pageHome, homeView
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
