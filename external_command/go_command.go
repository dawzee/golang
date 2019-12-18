// A sample program to demonstrate running and interfacing with an external command using go routines
package main

import (
	"fmt"
	// "log"
	// "io/ioutil"
	// "os"
	"os/exec"
	// "time"
)

func main() {
	cmd := exec.Command("delay", "5")
	done := make(chan bool)
	stdout := make(chan string)
	stderr := make(chan string)

	// Connect pipes to stdin, stdout and stderr
	// stdin, _ := cmd.StdinPipe()
	// stderr, _ := cmd.StderrPipe()
	// defer stdin.Close()
	// defer stderr.Close()

	// A go routine to read the external commmand stdout and feed into a channel
	go func() {
		stdoutPipe, _ := cmd.StdoutPipe()
		defer stdoutPipe.Close()

		buffer := make([]byte, 100)

	LOOP:
		for {
			numread, readerr := stdoutPipe.Read(buffer)
			if readerr != nil {
				// fmt.Println("rd thd: stdout err=", readerr)
				break LOOP

			} else {
				// fmt.Println("rd thd: stdout", numread, "bytes")
				stdout <- string(buffer[0:numread])
			}

			// fmt.Println("rd thd: loop")
		}
		// fmt.Println("rd thd: done")
	}()

	// A go routine to read the external commmand stderr and feed into a channel
	go func() {
		stderrPipe, _ := cmd.StderrPipe()
		defer stderrPipe.Close()

		buffer := make([]byte, 100)

	LOOP:
		for {
			numread, readerr := stderrPipe.Read(buffer)
			if readerr != nil {
				// fmt.Println("rd thd: stderr err=", readerr)
				break LOOP

			} else {
				// fmt.Println("rd thd: stderr", numread, "bytes")
				stderr <- string(buffer[0:numread])
			}

			// fmt.Println("rd thd: loop")
		}
		// fmt.Println("rd thd: done")
	}()

	go func() {
		// fmt.Println("go thd: running")
		if err := cmd.Run(); err != nil {
			// fmt.Println("go thd: err=", err)
		}
		// fmt.Println("go thd: done")
		done <- true
	}()

	fmt.Println("mn thd: waiting")

LOOP:
	for {
		select {
		case <-done:
			fmt.Println("mn thd: done!")
			break LOOP
		case out := <-stdout:
			fmt.Print("mn thd: out=", out)
		case err := <-stderr:
			fmt.Print("mn thd: err=", err)
		}
	}
}
