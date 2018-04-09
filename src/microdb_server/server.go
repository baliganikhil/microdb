package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	Delimiter = '\r'
)

func main() {
	config := GetServerConfig()
	connHost := config.ConnectionInfo.Host
	connPort := config.ConnectionInfo.Port
	connType := config.ConnectionInfo.ConnType

	initDB()

	var l, err = net.Listen(connType, connHost+":"+connPort)

	if err != nil {
		serverStartError(err)
	}

	defer l.Close()
	fmt.Println("MicroDB Server started... Listening at " + connHost + ":" + connPort)

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
	for {
		fmt.Println("Handling request")
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("Command: " + message + "\n")

		if message == "show dbs\n" {
			listDbs(conn)
		} else {
			conn.Write([]byte("Unrecognised command\n"))
		}
	}
}

func listDbs(conn net.Conn) {
	dbs := getDBInfo().DBs

	for _, db := range dbs {
		conn.Write([]byte(db.Name + "\n"))
	}

	conn.Write([]byte(string(Delimiter)))

}
