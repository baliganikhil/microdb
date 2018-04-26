package main

import (
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

	initLogger()
	initDB()

	var l, err = net.Listen(connType, connHost+":"+connPort)

	if err != nil {
		serverStartError(err)
	}

	defer l.Close()

	startMessage := "MicroDB Server started... Listening at " + connHost + ":" + connPort
	Log.Println(startMessage)
	fmt.Println(startMessage)

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
