package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Starting main.")
	sig := make(chan os.Signal, 1)
	waiter := make(chan struct{})
	signal.Notify(sig, syscall.SIGTERM)
	fmt.Println("Awaiting SIGTERM Signal from Docker Stop")
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("Just loopin' away!")
		}
	}()
	s := <-sig
	go func() {
		fmt.Println("Going to sleep a couple seconds...")
		time.Sleep(2 * time.Second)
		close(waiter)
	}()
	fmt.Println("Signal received: ", s)
	<-waiter
	fmt.Println("Buh bye!")
}
