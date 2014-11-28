// +build linux

package main

import (
	"log"
	"os/exec"
)

const PS_NAME string = "chrome"

// killApp kills any processes it can find with the name
// defined in PS_NAME
func killApp() {
	log.Println("killing " + PS_NAME + " processes")
	_, cmd := exec.Command("killall", "-v", PS_NAME).Output()
	if cmd != nil {
		log.Fatalln(cmd)
	}
	log.Println("killed " + PS_NAME + " processes")
}
