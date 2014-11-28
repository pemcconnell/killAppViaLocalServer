// +build windows

package main

import (
	"log"
	"os/exec"
)

const PS_NAME string = "chrome.exe"

func killApp() {
	log.Println("killing " + PS_NAME + " processes")
	_, cmd := exec.Command("taskkill", "/im", PS_NAME).Output()
	if cmd != nil {
		log.Fatalln(cmd)
	}
	log.Println("killed " + PS_NAME + " processes")
}
