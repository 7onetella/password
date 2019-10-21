package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/7onetella/password/api/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// /password/api/username/scott?tags=bank,chase

func getUUID() string {
	return uuid.New().String()
}

// errorMessageHandler writes the given status code and message
func errorMessageHandler(message string, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}

// CreatePasswordRequestHandler handles creating of password request
func CreatePasswordRequestHandler(w http.ResponseWriter, req *http.Request) {

	CORSHeader(w)

	ContentTypeJSON(w)

	passwordRequest, err := DecodePasswordRequest(req)
	WriteServerError(err, "error while retrieving json body", w)
	if err != nil {
		return
	}

	ID, err := CreatePassword(passwordRequest.Data)
	if err != nil {
		errorMessageHandler("error while creating record", 500, w)
		return
	}

	response := model.PasswordInput{
		Data: passwordRequest.Data,
	}
	response.Data.ID = ID

	responseData, err := json.Marshal(&response)
	if err != nil {
		errorMessageHandler("error while marshalling record", 500, w)
		return
	}

	LogPasswordRequest(passwordRequest)

	w.WriteHeader(http.StatusOK)
	w.Write(responseData)

}

// ReadPasswordRequestHandler retrieves password by id
func ReadPasswordRequestHandler(w http.ResponseWriter, req *http.Request) {
	CORSHeader(w)
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(req)

	ID := vars["id"]

	log.Println("id:", ID)

	// try retieving the records
	password, err := ReadPassword(ID)
	if err != nil {
		errorMessageHandler("error while finding record", 500, w)
		return
	}

	response := model.PasswordInput{
		Data: password,
	}

	data, err := json.Marshal(&response)
	if err != nil {
		errorMessageHandler("error while marshalling", 500, w)
		return
	}

	w.Write(data)
}

// ListPasswordsRequestHandler retrieves password by given parameters
func ListPasswordsRequestHandler(w http.ResponseWriter, req *http.Request) {
	uuid := getUUID()

	CORSHeader(w)
	w.Header().Add("Content-Type", "application/json")
	// vars := mux.Vars(req)
	title := req.FormValue("filter[title]")
	token := req.FormValue("token")
	ptoken := req.FormValue("ptoken")
	size, _ := strconv.Atoi(req.FormValue("size"))
	if size == 0 {
		size = 20
	}

	fmt.Println("title:", title)

	if title == "" {
		title = "%"
	}

	log.Println("token:", token)

	// try retieving the records
	passwords, err := FindPasswordsByTitle(uuid, title, token, size)
	if err != nil {
		errorMessageHandler("error while finding records", 500, w)
		return
	}

	response := paginate(passwords, "/list", token, ptoken, size)
	response.Metadata.Next = response.Metadata.Next + "&title=" + title
	response.Metadata.Prev = response.Metadata.Prev + "&title=" + title

	data, err := json.Marshal(&response)
	if err != nil {
		errorMessageHandler("error while marshalling", 500, w)
		return
	}

	w.Write(data)

}

// AuthRequired requires auth for given handler
func AuthRequired(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		c, err := req.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := c.Value

		ID, err := DecodeToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fmt.Println("Auth Token Decoded. ID: ", ID)

		// allow access
		next(w, req)
	}
}

// SignRequestHandler signs user
func SignRequestHandler(w http.ResponseWriter, req *http.Request) {
	CORSHeader(w)

	w.Header().Add("Content-Type", "application/json")
	username := req.FormValue("username")
	password := req.FormValue("password")

	if username == "scott.seo@gmail.com" && password == "password" {
		tokenString, expiration, err := EncodeID("id1234")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("Sign-In successful dropping token cookie")
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expiration,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// CORSHeader adds cores header
func CORSHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
}

// ContentTypeJSON writes content type json
func ContentTypeJSON(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}

// PreflightOptionsHandler handles preflight OPTIONS
func PreflightOptionsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, OPTIONS, HEAD, POST, PUT, PATCH, DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusAccepted)
}

func paginate(passwords []model.Password, path, token, ptoken string, size int) model.ListPasswordsOutput {

	resultLength := len(passwords)

	log.Println("len(passwords)=", resultLength)
	log.Println("token:", token)
	log.Println("ptoken:", ptoken)

	var nextURL string
	var prevURL string

	// possibly more? records then resolve next url and prev url
	if resultLength == size {
		// nextID would be the last record id since query continues from last record then on
		nextID := passwords[resultLength-1].ID

		tokenLen := len(token)

		params := map[string]string{}
		params["size"] = strconv.Itoa(size)

		// first page
		if tokenLen == 0 {
			params["token"] = nextID
			nextURL = resolveURLWithParams(path, params)

			delete(params, "token")
			prevURL = resolveURLWithParams(path, params)
		}

		// second page and on
		if tokenLen > 0 {
			params["token"] = nextID
			params["ptoken"] = token
			nextURL = resolveURLWithParams(path, params)

			params["token"] = ptoken
			params["ptoken"] = ""
			prevURL = resolveURLWithParams(path, params)
		}
	}

	response := model.ListPasswordsOutput{
		Items: passwords,
		Metadata: model.Metadata{
			Size:  size,
			Count: resultLength,
			Next:  nextURL,
			Prev:  prevURL,
		},
	}

	return response
}

// UpdatePasswordRequestHandler handles creating of password request
func UpdatePasswordRequestHandler(w http.ResponseWriter, req *http.Request) {

	CORSHeader(w)

	ContentTypeJSON(w)

	dataPassword, err := DecodePasswordRequest(req)
	WriteServerError(err, "error while retrieving json body", w)
	if err != nil {
		return
	}

	err = UpdatePassword(dataPassword.Data)
	WriteServerError(err, "error while persisting to db", w)
	if err != nil {
		return
	}

	LogPasswordRequest(dataPassword)

	w.WriteHeader(http.StatusNoContent)
}

// DeletePasswordRequestHandler deletes password by id
func DeletePasswordRequestHandler(w http.ResponseWriter, req *http.Request) {
	CORSHeader(w)

	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(req)

	ID := vars["id"]

	log.Println("id:", ID)

	// try retieving the records
	_, err := ReadPassword(ID)
	if err != nil {
		errorMessageHandler("error while finding record", 500, w)
		return
	}

	err = DeletePassword(ID)
	if err != nil {
		errorMessageHandler("error while finding record", 500, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HealthCheckHandler handles health check
func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	log.Println("/api/health")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// VersionHandler handles health check
func VersionHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	log.Println("/api/version")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("0.0.2"))
}

func resolveURLWithParams(partialPath string, params map[string]string) string {
	if !strings.HasPrefix(partialPath, "/") {
		partialPath = "/" + partialPath
	}

	url := URLBase + partialPath + "?"

	index := 0
	for k, v := range params {
		if index == 0 {
			url += fmt.Sprintf("%s=%s", k, v)
			continue
		}

		url += fmt.Sprintf("&%s=%s", k, v)
	}

	return url
}

func resolveURL(endpointPrefix, token, ptoken string, size int) string {
	url := fmt.Sprintf("%s/%s?size=%d", URLBase, endpointPrefix, size)

	if len(token) > 0 {
		url += "&token=" + token
	}

	if len(ptoken) > 0 {
		url += "&ptoken=" + ptoken
	}

	return url
}

// DataCollectHandler data collect handler
func DataCollectHandler(w http.ResponseWriter, req *http.Request) {

	CORSHeader(w)
	w.Header().Add("Content-Type", "application/json")

	data := `{
    "statusCode": 200,
    "status": "success",
    "errorMessage": null
}
`
	time.Sleep(1 * time.Second)
	w.Write([]byte(data))
}

// DecodePasswordRequest decodes password request
func DecodePasswordRequest(req *http.Request) (model.PasswordInput, error) {
	var dataPassword model.PasswordInput
	vars := mux.Vars(req)
	ID := vars["id"]

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return dataPassword, err
	}

	json.Unmarshal(data, &dataPassword)
	dataPassword.Data.ID = ID

	return dataPassword, nil
}

// WriteServerError writes server error if err is not nuil
func WriteServerError(err error, message string, w http.ResponseWriter) {
	if err != nil {
		errorMessageHandler(message, 500, w)
	}
}

// LogPasswordRequest logs password request
func LogPasswordRequest(passwordRequest model.PasswordInput) {
	log.Println("rid:", getUUID(), "received:", passwordRequest.Data)
}

// /api/list/bank,chase

// /api/version

// /health
