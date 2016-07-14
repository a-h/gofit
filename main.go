package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/api/fitness/v1"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func loadGoogleAPIConf(jsonFilename string) *oauth2.Config {
	json, err := ioutil.ReadFile(jsonFilename)

	if err != nil {
		log.Fatal("Failed to read the client_id.json file from disk. ", err)
	}

	conf, err := google.ConfigFromJSON(json)
	conf.Scopes = append(conf.Scopes, fitness.FitnessActivityReadScope)
	conf.Scopes = append(conf.Scopes, fitness.FitnessBodyReadScope)
	conf.Scopes = append(conf.Scopes, fitness.FitnessLocationReadScope)

	if err != nil {
		log.Fatal("Failed to get the API configuration from the client_id.json file. ", err)
	}

	return conf
}

func main() {
	// Initialise the Gorilla Router.
	r := mux.NewRouter()

	skey := []byte{22, 141, 7, 140, 241, 190, 89, 157, 164, 226, 50, 157, 104, 57, 245, 124, 15, 47, 178, 49, 70, 124, 165, 51, 64, 197, 122, 71, 201, 226, 132, 161}

	cs := createCookieStore(skey, false)

	// Handle incoming login requests.
	conf := loadGoogleAPIConf("client_id.json")

	loginHandler := &loginHandler{
		conf: conf,
	}
	r.Handle("/", loginHandler)

	// Handle incoming /oauth2callback requests and set the token cookie.
	cbh := &oauthCallbackHandler{
		conf:        conf,
		cookieStore: cs,
	}
	r.Handle("/oauth2callback", cbh)

	// The data JSON endpoint.
	sh := &dataHandler{
		conf:        conf,
		cookieStore: cs,
	}
	r.Handle("/data/", sh)

	// The dashboard.
	dh := &dashboardHandler{}
	r.Handle("/dashboard/", dh)

	// Serve static content.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Start the server with the routes.
	http.ListenAndServe(":8080", r)
}

func createCookieStore(encryptionKey []byte, setSecureFlag bool) *sessions.CookieStore {
	store := sessions.NewCookieStore(encryptionKey)
	store.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   setSecureFlag,
	}
	return store
}

func createSessionEncryptionKey() []byte {
	key := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	r.Read(key)
	return key
}
