package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/7onetella/password/api/client"
	"github.com/7onetella/password/api/model"
	homedir "github.com/mitchellh/go-homedir"
)

var infoBar = tview.NewTextView()
var signinForm = tview.NewForm()
var isSignedIn = false

func signInPage() (title string, content tview.Primitive) {
	signInFlexView := tview.NewFlex()

	infoBarReset()

	signinForm.Clear(true)
	signInPageReset()

	rows := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(infoBar, 2, 1, false).
		AddItem(signinForm, 0, 2, true)

	signInFlexView.AddItem(rows, 0, 1, true)
	signInFlexView.SetBorder(true)

	return "Sign in", signInFlexView
}

func infoBarReset() {
	infoBar.SetText("")
	infoBar.SetBorderPadding(1, 0, 2, 1)
}

func signInPageReset() {
	signinForm.
		AddInputField("Username:", "", 20, nil, nil).
		AddPasswordField("Password:", "", 20, '*', nil).
		AddButton("Sing In", signinAction)

	signinForm.SetBorder(false).SetBorderPadding(1, 1, 2, 1)
}

func signinAction() {
	var err error
	svc, err = client.NewPasswordService()
	if err != nil {
		notify(err.Error())
	}
	username := getInputValue(signinForm, "Username:")
	password := getInputValue(signinForm, "Password:")

	credentials := model.Credentials{
		Username: username,
		Password: password,
	}
	err = svc.Signin(credentials)
	if err != nil {
		isSignedIn = false
		infoBar.SetText("Authentication Error").SetTextColor(tcell.ColorRed)
		go func() {
			time.Sleep(3 * time.Second)
			infoBarReset()
			app.Draw()
		}()

		notify("error while authenticating: " + err.Error())
		return
	}

	isSignedIn = true

	WriteAuthToken(svc.Token, svc.Expiration)
	go RefreshTokenInBackground(svc)

	gotoPage(0)

	clearMenu()
	unloadPages()

	slides = signedInSlides()
	loadPages(signedInSlides())

	gotoPage(0)
	app.Draw()
}

func getInputValue(form *tview.Form, label string) string {
	fi := form.GetFormItemByLabel(label)
	input, ok := fi.(*tview.InputField)
	if ok {
		return input.GetText()
	}
	return ""
}

// ReadAuthToken reads auth token
func ReadAuthToken() (*model.AuthToken, error) {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	o := &model.AuthToken{}
	data, err := ioutil.ReadFile(home + "/.keepass.json")
	err = json.Unmarshal(data, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// WriteAuthToken reads auth token
func WriteAuthToken(token string, expiration int64) error {
	authToken := &model.AuthToken{
		Token:      token,
		Expiration: expiration,
	}

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := json.Marshal(authToken)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(home+"/.keepass.json", data, 0755)
	if err != nil {
		return err
	}

	return nil
}

// RefreshTokenInBackground ensures returned service client is authenticated
func RefreshTokenInBackground(svc *client.PasswordService) {
	for {
		time.Sleep(10 * time.Second)
		if !isSignedIn {
			break
		}

		var expired bool
		// read .keepass.json for session token
		authToken, err := ReadAuthToken()
		if err != nil {
			// file may not exist
			notify("error reading auth token: " + err.Error())
			continue
		}

		expired = time.Now().Unix() > authToken.Expiration
		if authToken != nil {
			now := time.Now().Unix()
			expiration := authToken.Expiration

			if !expired {
				remaining := expiration - now
				notify("time left: " + strconv.Itoa(int(remaining)))

				// if less than 30 seconds is left then refresh token
				if remaining < 110 {
					notify("refreshing token")
					svc.Token = authToken.Token
					err = svc.RefreshToken()
					if err != nil {
						notify("error while refreshing token")
						continue
					}
					notify("token refreshed")
					WriteAuthToken(svc.Token, svc.Expiration)
				}
			} else {
				// if expired then prompt for auth
				timeElapsed := now - authToken.Expiration
				notify("time elapsed:" + strconv.Itoa(int(timeElapsed)))
				continue
			}
		}

		// check expiration
		// if expired, then prompt for authentication and signin to populate Authentication
		if expired {
			notify("token exipred")
			continue
		}

		// if not expired, then populate service Authentication
		svc.Authorization = "Bearer " + authToken.Token
	}

}