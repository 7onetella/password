package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/markbates/refresh/refresh/web"
)

var httpPort string
var host string

// URLBase url base for next link
var URLBase string

// Version api version
var Version string

func init() {
	serverAddr := os.Getenv("SERVER_ADDR")
	terms := strings.Split(serverAddr, ":")

	host = terms[0]
	httpPort = terms[1]

	URLBase = fmt.Sprintf("http://%s:%s/api", host, httpPort)
}

func main() {
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(PreflightOptionsHandler)

	r.Path("/api/passwords").Methods("POST").HandlerFunc(CreatePasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("GET").HandlerFunc(ReadPasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("PATCH").HandlerFunc(UpdatePasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("DELETE").HandlerFunc(DeletePasswordRequestHandler)
	r.Path("/api/passwords").Methods("GET").HandlerFunc(ListPasswordsRequestHandler)
	r.Path("/api/passwords").Queries("filter[title]", "{title}").Queries("token", "{token}").Queries("ptoken", "{ptoken}").HandlerFunc(ListPasswordsRequestHandler)

	r.Path("/api/health").Methods("GET").HandlerFunc(AuthRequired(HealthCheckHandler))
	r.Path("/api/version").Methods("GET").HandlerFunc(VersionHandler)
	r.Path("/api/signin").Methods("POST").HandlerFunc(SignRequestHandler)
	r.Path("/data/collect/v1/").Methods("POST").HandlerFunc(DataCollectHandler)
	http.Handle("/", r)

	log.Println("starting server on port " + httpPort)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", httpPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      web.ErrorChecker(r), // Pass our instance of gorilla/mux in.
	}

	log.Fatal(srv.ListenAndServe())
}
