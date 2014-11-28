package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/kill", handler)
	log.Println("Listening on " + LISTEN_URL)
	err := http.ListenAndServe(LISTEN_URL, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	KillApp()
	w.Write([]byte("OK"))
}

func KillApp() {
	killApp()
}
