package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

// Page is a function which returns the page's main primitive and its title.
type Page func() (title string, content tview.Primitive)

var pages = tview.NewPages()

var currentSlide = 0

var pageitems []Page

var pageIndex = map[string]string{}

// SignedInPages returns signed-in pages
func SignedInPages() []Page {
	return []Page{
		Home,
		NewPage,
		SearchPage,
		SignOutPage,
	}
}

// SignedOutPages returns signed-out pages
func SignedOutPages() []Page {
	return []Page{
		Home,
		About,
		SignInPage,
	}
}

func gotoPage(title string) {
	debug("going to " + title)
	indexStr := pageIndex[title]
	index, _ := strconv.Atoi(indexStr)
	currentSlide = index
	menubar.Highlight(indexStr).ScrollToHighlight()
	pages.SwitchToPage(indexStr)
}

func previousPage() {
	currentSlide = (currentSlide - 1 + len(pageitems)) % len(pageitems)
	menubar.Highlight(strconv.Itoa(currentSlide)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(currentSlide))
}

func nextPage() {
	currentSlide = (currentSlide + 1) % len(pageitems)
	menubar.Highlight(strconv.Itoa(currentSlide)).ScrollToHighlight()
	pages.SwitchToPage(strconv.Itoa(currentSlide))
}

func loadPages() {
	for index, slide := range pageitems {
		title, primitive := slide()
		fmt.Fprintf(menubar, `["%d"][darkcyan]%s[white][""]  `, index, title)
		indexStr := strconv.Itoa(index)
		pages.AddPage(indexStr, primitive, true, index == currentSlide)
		pageIndex[title] = indexStr
	}
}

func unloadPages() {
	for index := range pageitems {
		if index == 0 {
			continue
		}
		pages.RemovePage(strconv.Itoa(index))
	}
}
