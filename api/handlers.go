package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/7onetella/password/api/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// URLBase url base for next link
var URLBase string

// Version api version
var Version string

func init() {
	httpPort = os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "9000"
	}

	host = os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	URLBase = fmt.Sprintf("http://%s:%s/api", host, httpPort)
}

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
	rid := getUUID()

	w.Header().Add("Content-Type", "application/json")
	//vars := mux.Vars(req)

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorMessageHandler("error while retrieving json body", 500, w)
		return
	}

	var password model.Password
	json.Unmarshal(data, &password)

	id, err := CreatePassword(password)
	if err != nil {
		errorMessageHandler("error while creating record", 500, w)
		return
	}

	var response model.CreatePasswordOutput
	response.ID = id
	response.RID = rid

	jsonData, err := json.Marshal(response)
	if err != nil {
		errorMessageHandler("error while marshalling record", 500, w)
		return
	}

	log.Println(string(jsonData))

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

// ReadPasswordRequestHandler retrieves password by id
func ReadPasswordRequestHandler(w http.ResponseWriter, req *http.Request) {
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

	data, err := json.Marshal(&password)
	if err != nil {
		errorMessageHandler("error while marshalling", 500, w)
		return
	}

	w.Write(data)
}

// ListPasswordsRequestHandler retrieves password by given parameters
func ListPasswordsRequestHandler(w http.ResponseWriter, req *http.Request) {
	uuid := getUUID()

	addCORSHeader(w)
	w.Header().Add("Content-Type", "application/json")
	// vars := mux.Vars(req)
	title := req.FormValue("title")
	token := req.FormValue("token")
	ptoken := req.FormValue("ptoken")
	size, _ := strconv.Atoi(req.FormValue("size"))
	if size == 0 {
		size = 20
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

func addCORSHeader(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
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
	rid := getUUID()

	w.Header().Add("Content-Type", "application/json")
	//vars := mux.Vars(req)

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorMessageHandler("error while retrieving json body", 500, w)
		return
	}

	var password model.Password
	json.Unmarshal(data, &password)

	err = UpdatePassword(password)
	if err != nil {
		errorMessageHandler("error while updating record", 500, w)
		return
	}

	log.Println(rid, "received:", password, "rid:", rid)

	w.WriteHeader(http.StatusNoContent)
}

// DeletePasswordRequestHandler deletes password by id
func DeletePasswordRequestHandler(w http.ResponseWriter, req *http.Request) {
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

// /api/list/bank,chase

// /api/version

// /health
