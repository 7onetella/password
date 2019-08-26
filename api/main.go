package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var httpPort string
var host string

func main() {
	r := mux.NewRouter()

	r.Path("/api/password").Methods("POST").HandlerFunc(CreatePasswordRequestHandler)
	r.Path("/api/password/{id}").Methods("GET").HandlerFunc(ReadPasswordRequestHandler)
	r.Path("/api/password").Methods("PUT").HandlerFunc(UpdatePasswordRequestHandler)
	r.Path("/api/password/{id}").Methods("DELETE").HandlerFunc(DeletePasswordRequestHandler)
	r.Path("/api/list").Queries("title", "{title}").Queries("token", "{token}").Queries("ptoken", "{ptoken}").HandlerFunc(ListPasswordsRequestHandler)
	r.Path("/api/list").Methods("GET").HandlerFunc(ListPasswordsRequestHandler)

	r.Path("/api/health").Methods("GET").HandlerFunc(HealthCheckHandler)
	r.Path("/api/version").Methods("GET").HandlerFunc(VersionHandler)
	http.Handle("/", r)

	log.Println("starting server on port " + httpPort)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", httpPort),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Fatal(srv.ListenAndServe())
}
