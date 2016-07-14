package main

import "net/http"

type dashboardHandler struct {
}

func (dh *dashboardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "dashboard.html", nil)
}
