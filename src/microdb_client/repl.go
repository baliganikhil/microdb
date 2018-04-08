package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "7188"
	CONN_TYPE = "tcp"
)

func main() {
	conn := connectToServer()
	printWelcome()
	runRepl(conn)
}

func runRepl(conn net.Conn) {
	printPrompt()
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		// fmt.Println(input)

		if input == "exit" || input == "quit" {
			cleanUpClientAndExit()
		} else if input == "" {
			// printPrompt()
		} else if input == "show dbs" {
			showDbs(conn)
		} else if input == "show tables" || input == "show collections" {
			showTables(conn)
		} else {
			fmt.Println("Unrecognised command")
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

func showDbs(conn net.Conn) {
	fmt.Println("Listing DBs")
	fmt.Println("---------------")
	conn.Write([]byte("show dbs\n"))

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print(message)
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

func cleanUpClientAndExit() {
	fmt.Println()
	fmt.Println("Thank you for using MicroDB")
	fmt.Println()
	os.Exit(0)
}
