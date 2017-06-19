package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values
// into another goroutine.

func main() {
	messages := make(chan string)

	go func() {
		// Send a value into a channel using the channel <- syntax.
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

/*

When we run the program the "ping" message is successfully
passed from one goroutine to another via our channel.
*/

/*
By default sends and receives block until both the sender
and receiver are ready. This property allowed us to wait at
the end of our program for the "ping" message without having to
use any other synchronization.
*/