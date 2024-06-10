package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "The address of this app")
	flag.Parse()

	// routing
	setRouting()

	log.Println("Starting Web Server. PORT: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListernAndServe: ", err)
	}
}
