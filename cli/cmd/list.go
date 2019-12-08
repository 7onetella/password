// Copyright © 2019 Seven OneTella<7onetella@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/7onetella/password/api/model"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list <title>",
	Short:   "Lists passwords",
	Long:    ``,
	Example: ``,
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		svc, err := GetAuthenticatedService()
		ExitOnError(err, "initializing client")

		input := model.ListPasswordsInput{}
		input.Title = "%%" // %% will list everything
		if len(args) > 0 {
			input.Title = args[0]
		}
		response, err := svc.ListPasswords(input)
		ExitOnError(err, "listing password")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Title", "Username"})

		for _, p := range response.Items {
			table.Append([]string{p.ID, p.Title, p.Username})
		}
		Newline()
		table.Render()
	},
}

func init() {

	rootCmd.AddCommand(listCmd)

}
