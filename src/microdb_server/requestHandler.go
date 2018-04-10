package main

import (
	"bufio"
	"fmt"
	"microdb_common"
	"net"
	"strings"
)

func handleRequest(conn net.Conn) {
	for {
		fmt.Println("Handling request")
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("Command: " + message + "\n")
		message = strings.Trim(message, "\n")

		c := getCommand(message)
		fmt.Println("Command: " + c.Command + "\n")

		// if message == "show dbs" {
		if microdbCommon.SHOW_DBS == c.Command {
			listDbs(conn)
		} else if microdbCommon.SHOW_TABLES == c.Command {
			listTables(conn)
		} else if microdbCommon.USE_DB == c.Command {
			useDB(conn, c)
		} else {
			conn.Write([]byte("Unrecognised command\n"))
		}
	}
}

func sendResponse(conn net.Conn, response string) {
	conn.Write([]byte(string(Delimiter)))
}
