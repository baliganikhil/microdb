package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	Delimiter = '\r'
)

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
		} else if input == "show dbs" {
			sendToServer(conn, ListDBs())
		} else {
			// fmt.Println("Unrecognised command")
			sendToServer(conn, input)
		}

		printPrompt()
	}
}

func printPrompt() {
	fmt.Print("> ")
}

func printWelcome() {
	fmt.Println("Welcome to MicroDB - By Nikhil Baliga")
}

func sendToServer(conn net.Conn, cmd string) {
	// fmt.Println("Listing DBs")
	// fmt.Println("---------------")
	conn.Write([]byte(cmd + "\n"))

	message, _ := bufio.NewReader(conn).ReadString(Delimiter)
	fmt.Println(message)
}

func showTables(conn net.Conn) {
	fmt.Println("Listing Tables")
	fmt.Println("---------------")

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
