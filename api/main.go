package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/password/api/scott", PasswordGetRequestHandler)
	http.Handle("/", r)

	log.Println("starting password")

	srv := &http.Server{
		Addr: "0.0.0.0:8088",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Fatal(srv.ListenAndServe())
}

// PasswordGetRequestHandler retrieves password
func PasswordGetRequestHandler(w http.ResponseWriter, req *http.Request) {
	json := `
{
	"items": [
		{
			"url"         : "http://example.com",
			"username"    : "user-foo",
			"password"    : "1234",
			"description" : "Bank of Example",
			"tags": ["example", "bank"]
		}
	],
	"metadata": {
		"count": 20,
		"more": true,
		"page": 0,
		"size": 20,
		"next": ""
	}
}	
	`

	w.Header()

	w.Write([]byte(json))
}
