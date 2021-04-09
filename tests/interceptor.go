package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)

	signal.Notify(c)
	done := false
	fmt.Println("Ready for interception")
	for done != true {
		select {
		case sig := <-c:
			fmt.Printf("Received signal %s, bye\n", sig)
			done = true
		}
	}
}
