package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		os.Exit(1)
	}()

	go func() {
		for {
			fmt.Println("hello")
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(1 * time.Second)

	go func() {
		for {
			fmt.Fprintf(os.Stderr, "hello stderr\n")
			time.Sleep(2 * time.Second)
		}
	}()

	stdinScanner := bufio.NewScanner(os.Stdin)

	for stdinScanner.Scan() {
		fmt.Println(stdinScanner.Text())
	}
}
