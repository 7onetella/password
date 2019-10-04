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

	r.Methods("OPTIONS").HandlerFunc(PreflightOptionsHandler)

	r.Path("/api/passwords").Methods("POST").HandlerFunc(AuthRequired(CreatePasswordRequestHandler))
	r.Path("/api/passwords/{id}").Methods("GET").HandlerFunc(ReadPasswordRequestHandler)
	r.Path("/api/passwords").Methods("PATCH").HandlerFunc(UpdatePasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("DELETE").HandlerFunc(DeletePasswordRequestHandler)
	r.Path("/api/passwords").Methods("GET").HandlerFunc(ListPasswordsRequestHandler)
	r.Path("/api/passwords").Queries("title", "{title}").Queries("token", "{token}").Queries("ptoken", "{ptoken}").HandlerFunc(ListPasswordsRequestHandler)

	r.Path("/api/health").Methods("GET").HandlerFunc(AuthRequired(HealthCheckHandler))
	r.Path("/api/version").Methods("GET").HandlerFunc(VersionHandler)
	r.Path("/api/signin").Methods("POST").HandlerFunc(SignRequestHandler)
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
