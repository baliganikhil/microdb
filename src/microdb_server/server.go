package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	config := GetServerConfig()
	CONN_HOST := config.ConnectionInfo.Host
	CONN_PORT := config.ConnectionInfo.Port
	CONN_TYPE := config.ConnectionInfo.ConnType

	var l, err = net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		serverStartError(err)
	}

	defer l.Close()
	fmt.Println("MicroDB Server started... Listening at " + CONN_HOST + ":" + CONN_PORT)

	setupServerInterruptHandler()
	setupRequestHandler(l)

}

func setupRequestHandler(l net.Listener) {
	for {
		conn, err := l.Accept()

		if err != nil {
			connAcceptError(err)
		}

		go handleRequest(conn)
	}
}

func serverStartError(err error) {
	fmt.Println("Could not start the server", err.Error())
	os.Exit(1)
}

func connAcceptError(err error) {
	fmt.Println("Could not accept connection", err.Error())
	os.Exit(1)
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	// buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	// message, err := conn.Read(buf)
	message, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		fmt.Println("Error reading:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Command: " + message + "\n")

	//
	if message == "show dbs\n" {
		listDbs(conn)
	} else {
		conn.Write([]byte("Unrecognised command\n"))
	}

	// Send a response back to person contacting us.

	// Close the connection when you're done with it.
	// conn.Close()
}

func listDbs(conn net.Conn) {
	conn.Write([]byte("nimbus\ntest\n"))
}
