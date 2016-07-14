package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"golang.org/x/oauth2"
)

type oauthCallbackHandler struct {
	conf        *oauth2.Config
	cookieStore *sessions.CookieStore
}

func (cbh *oauthCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Get the code from the response
	code := r.FormValue("code")

	session, err := cbh.cookieStore.Get(r, "gofit")

	if err != nil {
		log.Print("Failed to get the session cookie value.")
		http.Error(w, "The session cookie is invalid: "+err.Error(), http.StatusInternalServerError)
	}

	// Handle the exchange code to initiate a transport.
	tok, err := cbh.conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	// Set the access token into a secure cookie and render a HTML page to pull the data from
	// the REST service for graphing.
	session.Values["token"] = tok.AccessToken
	session.Save(r, w)

	http.Redirect(w, r, "/dashboard/", http.StatusFound)
}
