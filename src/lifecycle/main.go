package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	waitForPostStart()
	http.HandleFunc("/", home)
	http.HandleFunc("/shutdown", shutdown)
	log.Println("starting application in por 8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	log.Println("shutdown initiated!")
	log.Println("doing some cleanup work")
	for i := 0; i < 10; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println()
	log.Print("done!")
	os.Exit(1)
}

func waitForPostStart() {
	wait := os.Getenv("WAIT_FOR_POST_START")
	for wait == "true" {
		if _, err := os.Stat("/tmp/poststart"); err == nil {
			log.Println("file created. starting application.")
			return
		} else if os.IsNotExist(err) {
			log.Println("file creation has not completed yet...")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("post-start has not completed yet...")
			time.Sleep(5 * time.Second)
		}
	}
}
