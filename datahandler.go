package main

import (
	"log"
	"net/http"

	"github.com/a-h/gofit/fitnessdata"
	"github.com/gorilla/sessions"

	"encoding/json"

	"golang.org/x/oauth2"
)

type dataHandler struct {
	conf        *oauth2.Config
	cookieStore *sessions.CookieStore
}

func (sh *dataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := sh.cookieStore.Get(r, "gofit")

	if err != nil {
		log.Print("Failed to get the session cookie value.")
		http.Error(w, "The session cookie is invalid: "+err.Error(), http.StatusInternalServerError)
	}

	tok := session.Values["token"].(string)

	log.Print("token:", tok)

	t := &oauth2.Token{
		AccessToken: tok,
	}

	fd, err := fitnessdata.Get(sh.conf, t)

	if err != nil {
		log.Print("Failed to get the fitness data: ", err)
		http.Error(w, "Failed to retrieve fitness data.", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fd)
}
