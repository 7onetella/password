package main

import (
	"github.com/rivo/tview"
)

func signOutPage() (title string, content tview.Primitive) {
	f := tview.NewForm().
		AddButton("Sign Out", signOutAction)
	f.SetBorder(true).SetTitle("Sign Out")
	f.SetBorderPadding(1, 1, 2, 1)
	return "Sign Out", f
}

func signOutAction() {
	isSignedIn = false
	notify("sign out initiated")
	gotoPage(0)

	clearMenu()
	unloadPages()

	loadPages(signedOutSlides())
	gotoPage(0)
	app.Draw()
}
