package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// TODO echo $UID check if you are root

	fmt.Println("Updating the system clock")
	sysclock := exec.Command("timedatectl", "set-ntp", "true")
	err := sysclock.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("System clock updated correctly")
}
