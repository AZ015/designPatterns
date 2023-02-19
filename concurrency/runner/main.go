package main

import (
	"log"
	"os"
	"time"
)

const timeout = 10 * time.Second

func main() {
	r := New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Println("Terminated due timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Println("Terminated due interrupt")
			os.Exit(2)
		}
	}

	log.Println("Process ended")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
