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
	"github.com/7onetella/password/api/client"
	"github.com/spf13/cobra"
)

var updateCmdTitle string
var updateCmdURL string
var updateCmdNotes string
var updateCmdTags []string
var updateCmdPassword string

var updateCmd = &cobra.Command{
	Use:     "update <ID>",
	Short:   "Updates password",
	Long:    ``,
	Example: ``,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		svc, err := client.NewPasswordService()
		ExitOnError(err, "initializing client")

		ID := args[0]

		record, err := svc.ReadPassword(ID)
		ExitOnError(err, "reading existing record")
		var dirty bool

		notEmpty := func(s string) bool {
			dirty = true
			return len(s) > 0
		}

		if notEmpty(updateCmdTitle) {
			record.Title = updateCmdTitle
		}

		if notEmpty(updateCmdURL) {
			record.URL = updateCmdURL
		}

		if notEmpty(updateCmdPassword) {
			record.Password = updateCmdPassword
		}

		if notEmpty(updateCmdNotes) {
			record.Notes = updateCmdNotes
		}

		if len(updateCmdTags) > 0 {
			record.Tags = updateCmdTags
		}

		if dirty {
			err = svc.UpdatePassword(*record)
			ExitOnError(err, "updating password")

			Success("updating password")
		}
	},
}

func init() {

	rootCmd.AddCommand(updateCmd)

	flags := updateCmd.Flags()

	flags.StringVar(&updateCmdTitle, "title", "", "required: title")
	flags.StringVarP(&updateCmdPassword, "password", "p", "", "optional: password")
	flags.StringVarP(&updateCmdURL, "url", "u", "", "optional: url")
	flags.StringVarP(&updateCmdNotes, "notes", "n", "", "optional: notes")
	flags.StringArrayVarP(&updateCmdTags, "tags", "t", []string{}, "optional: tags")

}
