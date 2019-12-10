module github.com/7onetella/password/tui

go 1.12

require (
	github.com/7onetella/password/api v0.0.0-20190926005830-4348a249e868
	github.com/7onetella/password/cli v0.0.0-20191208085244-df3ef8f0e3b8 // indirect
	github.com/gdamore/tcell v1.3.0
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rivo/tview v0.0.0-20191129065140-82b05c9fb329
)

replace github.com/7onetella/password/api => ../api

replace github.com/7onetella/password/cli => ../cli
