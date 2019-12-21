package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/7onetella/password/api/client"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var menubar = newMenuBar()

var flex = tview.NewFlex()

var rows *tview.Flex

var svc *client.PasswordService

var prevFocused tview.Primitive

func init() {
	stage := os.Getenv("STAGE")
	if stage == "" {
		panic("set STAGE environment variable")
	}

	serverAddr := lookupSRV("password-" + stage + "-app")

	if stage == "localhost" {
		serverAddr = "localhost:4242"
	}
	var err error
	svc, err = client.NewPasswordServiceWithServerAddress(serverAddr)
	if err != nil {
		debug(err.Error())
	}
}

func main() {

	menubar.Highlight(strconv.Itoa(currentSlide))

	pageitems = SignedOutPages()
	loadPages()

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
				prevFocused = app.GetFocus()
				app.SetFocus(notification)
				notification.SetSelectable(true, false)
			} else {
				notification.SetSelectable(false, false)
				app.SetFocus(prevFocused)
			}
			app.Draw()
		case tcell.KeyCtrlC:
			if err := app.SetRoot(confirmQuit(), false).Run(); err != nil {
				panic(err)
			}
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

func newMenuBar() *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)
}

func confirmQuit() *tview.Modal {
	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
				os.Exit(0)
			} else {
				if err := app.SetRoot(flex, true).Run(); err != nil {
					panic(err)
				}
			}
		})
	return modal
}

func lookupSRV(serviceName string) string {

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "tcp", net.JoinHostPort("127.0.0.1", "8600"))
		},
	}

	name := serviceName + ".service.dc1.consul"
	_, srvs, err := resolver.LookupSRV(context.Background(), "", "", name)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for i, srv := range srvs {
		ips, err := resolver.LookupHost(context.Background(), srv.Target)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		// fmt.Printf("target: %s, port: %d, priority: %d, weight: %d", ips[0], srv.Port, srv.Priority, srv.Weight)
		if i == 0 {
			return fmt.Sprintf("%s:%d", ips[0], srv.Port)
		}
	}

	return ""
}
