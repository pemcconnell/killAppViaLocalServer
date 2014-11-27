// +build windows

package main

import (
    "log"
    "os/exec"
)

const LISTEN_URL string = "127.0.0.1:6732"
const PS_NAME string = "chrome.exe"

func killChrome() {
    log.Println("killing chrome processes")
    _, cmd := exec.Command("taskkill", "/im", PS_NAME).Output()
    if cmd != nil {
        log.Fatalln(cmd)
    }
    log.Println("killed chrome processes")
}