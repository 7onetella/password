package main

import (
	"fmt"

	"github.com/rivo/tview"
)

const pageHome = "Home"

const pageAbout = "About"

// Home returns home title and page item
func Home() (title string, content tview.Primitive) {
	homeView := tview.NewTextView().SetWordWrap(true)
	fmt.Fprint(homeView, `    Hello, my dear wife, love of my life!
	I created this app for you to manage your passwords.
	
	I love you.	`)
	homeView.SetBorder(true)
	homeView.SetBorderPadding(4, 0, 4, 0)
	return "Home", homeView
}

// About returns about title and page item
func About() (title string, content tview.Primitive) {
	aboutView := tview.NewTextView().SetWordWrap(true)
	fmt.Fprint(aboutView, `    Motivation

	It's hard to remember so many accounts and their login credentials. 
	I decided to learn rivo/tview and create a tool in the process.
	`)
	aboutView.SetBorder(true)
	aboutView.SetBorderPadding(4, 0, 4, 0)
	return "About", aboutView
}
