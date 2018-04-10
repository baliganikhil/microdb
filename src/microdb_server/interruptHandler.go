package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func setupServerInterruptHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanUpServerAndExit()
	}()
}

func cleanUpServerAndExit() {
	fmt.Println()
	fmt.Println("Thank you for using MicroDB")
	fmt.Println()
	os.Exit(0)
}
