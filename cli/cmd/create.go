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
	"fmt"

	"github.com/7onetella/password/api/client"
	"github.com/7onetella/password/api/model"
	"github.com/spf13/cobra"
)

var createCmdTitle string
var createCmdURL string
var createCmdNotes string

var createCmd = &cobra.Command{
	Use:     "create <username> <password>",
	Short:   "Creates password",
	Long:    ``,
	Example: ``,
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		svc, err := client.NewPasswordService()
		ExitOnError(err, "initializing client")

		username, pwd := credentials()
		err = svc.Signin(model.Credentials{Username: username, Password: pwd})
		ExitOnError(err, "authenticating")

		p := model.Password{}
		p.Title = createCmdTitle
		p.URL = createCmdURL
		p.Username = args[0]
		p.Password = args[1]
		p.Notes = createCmdNotes

		output, err := svc.CreatePassword(model.PasswordInput{Data: p})
		ExitOnError(err, "creating password")

		Success("creating password")

		fmt.Println("ID: " + output.Data.ID)

	},
}

func init() {

	rootCmd.AddCommand(createCmd)

	flags := createCmd.Flags()

	flags.StringVar(&createCmdTitle, "title", "", "required: title")
	flags.StringVarP(&createCmdURL, "url", "u", "", "required: url")
	flags.StringVarP(&createCmdNotes, "notes", "n", "", "required: notes")

}
