package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	_client_version = "1.0.0"
	Delimiter       = '\r'
)

var curDBName = "test"

func main() {
	conn := connectToServer()
	printWelcome()
	setupClientInterruptHandler(conn)
	runRepl(conn)
}

func runRepl(conn net.Conn) {
	printPrompt()
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" || input == "quit" {
			cleanUpClientAndExit(conn)
		} else if input == "version" {
			fmt.Println(_client_version)
		} else if len(strings.TrimSpace(input)) == 0 {
			// Do nothing
		} else {
			sendToServer(conn, commandParser(input).ToJson())
		}

		printPrompt()
	}
}

func printPrompt() {
	fmt.Print("μDB > ")
}

func printWelcome() {
	fmt.Println("Welcome to MicroDB - By Nikhil Baliga")
}

func sendToServer(conn net.Conn, cmd string) {
	// fmt.Println("Listing DBs")
	// fmt.Println("---------------")
	conn.Write([]byte(cmd + "\n"))

	message, _ := bufio.NewReader(conn).ReadString(Delimiter)
	handleResponse(message)
}

func connectToServer() net.Conn {
	config := GetClientConfig()

	conn, err := net.Dial(config.ConnectionInfo.ConnType, config.ConnectionInfo.Host+":"+config.ConnectionInfo.Port)
	if err != nil {
		fmt.Printf("Could not connect to the database server at %s:%s\n", config.ConnectionInfo.Host, config.ConnectionInfo.Port)
		os.Exit(1)
	}

	return conn
}
