package main

import (
	"net/http"
)

func setRouting() {
	http.HandleFunc("/sample", sampleHandler)
}
