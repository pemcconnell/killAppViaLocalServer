// +build linux

package main

import (
	"log"
	"os/exec"
)

const LISTEN_URL string = "127.0.0.1:6732"
const PS_NAME string = "chrome"

func killApp() {
	log.Println("killing " + PS_NAME + " processes")
	_, cmd := exec.Command("killall", "-v", PS_NAME).Output()
	if cmd != nil {
		log.Fatalln(cmd)
	}
	log.Println("killed " + PS_NAME + " processes")
}
