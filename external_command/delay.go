package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if delay, err := strconv.Atoi(os.Args[1]); err == nil {
		// time.Sleep(time.Duration(delay) * time.Second)
		for count := 0; count < delay; count++ {
			fmt.Println("delay",count)
			fmt.Fprintln(os.Stderr,"Some debugging text on stderr")
			fmt.Fprintln(os.Stderr,"...could be many lines")
			time.Sleep(time.Second)
		}
	}
}
