package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func setupClientInterruptHandler(conn net.Conn) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanUpClientAndExit(conn)
	}()
}

func cleanUpClientAndExit(conn net.Conn) {
	conn.Close()

	fmt.Println()
	fmt.Println("Thank you for using MicroDB")
	fmt.Println()
	os.Exit(0)
}
