package main

import (
	"github.com/mitchellh/go-ps"
	"log"
	"net/http"
	"os"
	"syscall"
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
	prcs, _ := ps.Processes()
	for _, p := range prcs {
		if p.Executable() == PS_NAME {
			pc, findErr := os.FindProcess(p.Pid())
			if findErr != nil {
				log.Fatalln(findErr)
			}
			killErr := pc.Signal(syscall.SIGTERM)
			if killErr != nil {
				log.Fatalln(killErr)
			}
		}
	}
	log.Println("killed chrome processes")
}
