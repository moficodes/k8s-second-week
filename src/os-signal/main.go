package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Doing all sorts of cleanup work!")
		time.Sleep(10 * time.Second)
		done <- true
		// uncomment to get to the rouge version
		// log.Println("going rouge")
		// log.Println("i am invincible")
	}()

	log.Println("awaiting signal")
	<-done
	log.Println("exiting")
}
