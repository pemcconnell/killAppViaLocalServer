package main

import (
	"log"
	"net/http"
)

const LISTEN_URL string = "127.0.0.1:6732"

func main() {
	runServer()
}

// runServer starts a local webserver
// the address is defined by LISTEN_URL
func runServer() {
	http.HandleFunc("/kill", handlerKill)
	log.Println("Listening on " + LISTEN_URL)
	err := http.ListenAndServe(LISTEN_URL, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

// handlerKill is the handler set up for the ./kill endpoint.
func handlerKill(w http.ResponseWriter, r *http.Request) {
	killApp()
	w.Write([]byte("OK"))
}
