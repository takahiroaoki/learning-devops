package main

import "net/http"

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, r, http.StatusOK, "Hello World")
}
