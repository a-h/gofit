package main

import (
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type loginHandler struct {
	conf *oauth2.Config
}

func (lh *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := lh.conf.AuthCodeURL("state")
	log.Print("Redirecting to login URL ", url)
	http.Redirect(w, r, url, http.StatusFound)
}
