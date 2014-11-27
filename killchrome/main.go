package main

import (
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/chrome", handler)
	log.Println("Listening on " + LISTEN_URL)
	err := http.ListenAndServe(LISTEN_URL, nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	killChrome()
	w.Write([]byte("OK"))
}

func killChrome() {
	log.Println("killing chrome processes")
	_, cmdErr := exec.Command("taskkill", "/im", "chrome.exe", "/f", "/t").Output()
	if cmdErr != nil {
		log.Fatalln(cmdErr)
	}
	log.Println("killed chrome processes")
}
