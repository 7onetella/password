package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/markbates/refresh/refresh/web"
)

var stage string
var port string
var db *sql.DB
var hmacSecret = []byte("")
var Username string
var Password string

// URLBase url base for next link
var URLBase string

// Version api version
var Version string

func init() {
	stage = GetEnvWithDefault("STAGE", "localhost")

	port = GetEnvWithDefault("HTTP_PORT", "4242")

	URLBase = fmt.Sprintf("http://%s:%s/api", stage, port)

	connStr := GetEnvWithDefault("DB_CONNSTR", "postgres://dev:dev114@localhost/devdb")
	dbinst, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("error while opening db connection")
	}
	db = dbinst

	cryptoToken := GetEnvWithDefault("CRYPTO_TOKEN", "crypto_token_12345")
	hmacSecret = []byte(cryptoToken)

	credential := GetEnvWithDefault("CREDENTIAL", "admin:password")
	terms := strings.Split(credential, ":")
	Username = terms[0]
	Password = terms[1]
}

// GetEnvWithDefault attemps to retrieve from env. default calculated based on stage if env value empty.
func GetEnvWithDefault(env, defaultV string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultV
	}
	return v
}

func main() {
	log.Println("stage:", stage)

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(PreflightOptionsHandler)

	r.Path("/api/passwords").Methods("POST").HandlerFunc(CreatePasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("GET").HandlerFunc(ReadPasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("PATCH").HandlerFunc(UpdatePasswordRequestHandler)
	r.Path("/api/passwords/{id}").Methods("DELETE").HandlerFunc(DeletePasswordRequestHandler)
	r.Path("/api/passwords").Methods("GET").HandlerFunc(AuthRequired(ListPasswordsRequestHandler))
	r.Path("/api/passwords").Queries("filter[title]", "{title}").Queries("token", "{token}").Queries("ptoken", "{ptoken}").HandlerFunc(AuthRequired(ListPasswordsRequestHandler))

	r.Path("/favicon.ico").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("/favicon")
		w.Write([]byte{})
	})
	r.Path("/api/health").Methods("GET").HandlerFunc(HealthCheckHandler)
	r.Path("/api/version").Methods("GET").HandlerFunc(VersionHandler)
	r.Path("/api/token-refresh").Methods("POST").HandlerFunc(TokenRefreshHandler)
	r.Path("/api/signin").Methods("POST").HandlerFunc(SigninRequestHandler)
	r.Path("/data/collect/v1/").Methods("POST").HandlerFunc(DataCollectHandler)
	r.PathPrefix("/ui/").Handler(http.StripPrefix("/ui/", http.FileServer(assetFS())))
	r.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("/ redirecting to /ui/")
		http.Redirect(w, r, "/ui/", http.StatusMovedPermanently)
	})

	http.Handle("/", r)
	log.Println("starting https server on port " + port)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      web.ErrorChecker(r), // Pass our instance of gorilla/mux in.
	}

	if stage == "keepass" {
		go log.Fatal(http.ListenAndServe(":80", Port80toHTTPSRedirectHandler()))
	}

	log.Fatal(srv.ListenAndServeTLS(GetCertAndKey()))
}

// GetCertAndKey return cert and key locations
func GetCertAndKey() (string, string) {
	// this will resolve to refresh or api
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("workfolder:", dir)

	return dir + "/" + stage + "-crt.pem", dir + "/" + stage + "-key.pem"
}
