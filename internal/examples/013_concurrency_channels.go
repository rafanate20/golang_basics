package examples

import (
	"fmt"
	"math/rand"
	"time"
)

func sayAgain(text string, c chan<- string) {
	c <- text
}

func UsingConcurrentChannels() {
	fmt.Println("Chapter 013 - Using Go Rutines and channels to conmunicate between Go Rutines")

	//channel
	c := make(chan string, 1)

	fmt.Println("Hello")

	go sayAgain("Bye\n----------------------------------------------------------------", c)

	fmt.Println(<-c)
}

func message(text string, c chan string) {
	var secondsWaiting = rand.Intn(5)
	time.Sleep(time.Second * time.Duration(secondsWaiting))
	c <- text
}

func UsingMoreChannels() {
	fmt.Println("Using Go Rutines and channels to conmunicate with main function")
	c := make(chan string, 2)
	fmt.Println("before send channel C", len(c), cap(c))
	message("msg 1", c) //c <- "msg 1"
	message("msg 2", c) //c <- "msg 2"
	fmt.Println("after wait channel C", len(c), cap(c))
	close(c)

	for msg := range c {
		fmt.Println(msg)
	}

	// Select and verify if channel is completed
	email1 := make(chan string)
	email2 := make(chan string)

	go message("user001@email.com", email1)
	go message("user002@email.com", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de canal 1", m1)
			defer close(email1)
		case m2 := <-email2:
			fmt.Println("Email recibido de canal 2", m2)
			defer close(email2)
		}
	}
}
