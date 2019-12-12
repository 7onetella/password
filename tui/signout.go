package main

import (
	"github.com/rivo/tview"
)

// SignOutPage returns signout page
func SignOutPage() (title string, content tview.Primitive) {
	f := tview.NewForm().
		AddButton("Sign Out", signOut)
	f.SetBorder(true).SetTitle("Sign Out")
	f.SetBorderPadding(1, 1, 2, 1)
	return "Sign Out", f
}

func signOut() {
	isSignedIn = false
	debug("sign out initiated")
	gotoPage(pageHome)

	clearMenu()
	unloadPages()

	pageitems = SignedOutPages()
	loadPages()
	gotoPage(pageHome)
	app.Draw()
}
