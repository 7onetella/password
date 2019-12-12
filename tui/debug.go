package main

import (
	"time"
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

var debugView *tview.Flex
var notification = newDebug()

func newDebug() *tview.Table {
	table := tview.NewTable().
		SetBorders(false).
		InsertColumn(0).
		InsertRow(0).
		InsertColumn(0).
		InsertColumn(0).
		SetSelectable(false, false).
		SetSelectedStyle(tcell.ColorBlack, tcell.ColorWhite, tcell.AttrNone)
	return table
}

func debug(message string) {
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
