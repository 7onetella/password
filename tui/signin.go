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

// SignInPage returns signin page
func SignInPage() (title string, content tview.Primitive) {
	signInFlexView := tview.NewFlex()

	infoBarReset()

	signinForm.Clear(true)
	signInPageReset()

	rows := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(infoBar, 2, 1, false).
		AddItem(signinForm, 0, 2, true)

	signInFlexView.AddItem(rows, 0, 1, true)
	signInFlexView.SetBorder(true)

	return "Sign In", signInFlexView
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

	username := getInputValue(signinForm, "Username:")
	password := getInputValue(signinForm, "Password:")

	credentials := model.Credentials{
		Username: username,
		Password: password,
	}
	err := svc.Signin(credentials)
	if err != nil {
		isSignedIn = false
		infoBar.SetText("Authentication Error").SetTextColor(tcell.ColorRed)
		go func() {
			time.Sleep(3 * time.Second)
			infoBarReset()
			app.Draw()
		}()

		debug("error while authenticating: " + err.Error())
		return
	}

	isSignedIn = true

	WriteAuthToken(svc.Token, svc.Expiration)
	go RefreshTokenInBackground(svc)

	gotoPage(pageHome)

	clearMenu()
	unloadPages()

	pageitems = SignedInPages()
	loadPages()

	gotoPage(pageHome)
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
		// this shouldn't happen, if it did
		if !isSignedIn {
			return
		}

		var expired bool
		// read .keepass.json for session token
		authToken, err := ReadAuthToken()
		if err != nil {
			// file may not exist
			debug("error reading auth token: " + err.Error())
			continue
		}
		if authToken == nil {
			debug("auth token is nil")
			continue
		}

		expired = time.Now().Unix() > authToken.Expiration
		now := time.Now().Unix()
		expiration := authToken.Expiration

		if !expired {
			remaining := expiration - now
			// if less than 30 seconds is left then refresh token
			if remaining < 30 {
				svc.Token = authToken.Token
				err = svc.RefreshToken()
				if err != nil {
					debug("error while refreshing token")
					// perhaps the server is busy or temporarily unavailable
					continue
				}
				debug("refreshed token successful")
				WriteAuthToken(svc.Token, svc.Expiration)
				// after RefreshToken() svc has the new token
				continue
			}
		} else {
			// if expired then prompt for auth
			timeElapsed := now - authToken.Expiration
			debug("time elapsed since token expired:" + strconv.Itoa(int(timeElapsed)))
			if err := app.SetRoot(authPrompt("Session Expired"), false).Run(); err != nil {
				panic(err)
			}
			return
		}
	}
}

func authPrompt(msg string) *tview.Modal {
	modal := tview.NewModal().
		SetText(msg).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				signOut()
				gotoPage("Sign In")
				if err := app.SetRoot(flex, true).Run(); err != nil {
					panic(err)
				}
			}
		})
	return modal
}
