module github.com/7onetella/password/cli

go 1.12

require (
	github.com/7onetella/password/api v0.0.0-20190924183701-189b4f6b667d
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/dgryski/go-sip13 v0.0.0-20190329191031-25c5027a8c7b // indirect
	github.com/fatih/color v1.7.0
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/olekukonko/tablewriter v0.0.1
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
)

replace github.com/7onetella/password/api => ../api
