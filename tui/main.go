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

var pages *tview.Pages

var menubar *tview.TextView

var app = tview.NewApplication()

var notification *tview.TextView

var currentSlide = 0

var slides []Slide

var rows *tview.Flex

var signoutform *tview.Form

var svc *client.PasswordService

var credentials model.Credentials

func init() {

	pages = tview.NewPages()

	menubar = newmenubar()

	notification = tview.NewTextView().SetWordWrap(true).SetChangedFunc(func() {
		app.Draw()
	})
	notification.SetBorder(true)
	notification.SetTitle("Debug")

	slides = signedOutSlides()

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

func notify(message string) {
	// notification.Clear()
	fmt.Fprint(notification, message+"\n")
	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	notification.Clear()
	// 	app.Draw()
	// }()
}

func gotoSlide(index int) {
	currentSlide = index
	menubar.Highlight(strconv.Itoa(index)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(index))
}

func previousSlide() {
	currentSlide = (currentSlide - 1 + len(slides)) % len(slides)
	menubar.Highlight(strconv.Itoa(currentSlide)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(currentSlide))
}

func nextSlide() {
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
		AddItem(pages, 0, 9, true).
		AddItem(notification, 0, 4, false)

	flex.AddItem(rows, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlL {
			nextSlide()
		} else if event.Key() == tcell.KeyCtrlH {
			previousSlide()
		} else if event.Key() == tcell.KeyCtrlL {
			notification.Clear()
			app.Draw()
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

func loadPages(newslides []Slide) {
	for index, slide := range newslides {
		title, primitive := slide()
		fmt.Fprintf(menubar, `["%d"][darkcyan]%s[white][""]  `, index, title)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == currentSlide)
	}
}

func unloadPages() {
	for index := range signedOutSlides() {
		if index == 0 {
			continue
		}
		pages.RemovePage(strconv.Itoa(index))
	}
}

func signOutAction() {
	notify("sign out initiated")
	gotoSlide(0)

	clearMenu()
	unloadPages()

	loadPages(signedOutSlides())
	gotoSlide(0)
	app.Draw()
}

func newmenubar() *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)
}

func newbox(title string) *tview.Box {
	return tview.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle(title)
}

func newtable(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Table {
	table := tview.NewTable()
	table.SetBorders(true)
	table.InsertColumn(0)
	table.InsertRow(0)
	table.SetCellSimple(0, 0, "cell")
	table.SetBorder(true)

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			notify("arrow key pressed")
			return nil
		}
		return event
	})

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

func signOutPage() (title string, content tview.Primitive) {
	f := tview.NewForm().
		AddButton("Sign Out", signOutAction)
	f.SetBorder(true).SetTitle("Sign Out")
	return "Sign Out", f
}
