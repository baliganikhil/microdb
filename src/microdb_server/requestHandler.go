package main

import (
	"bufio"
	"encoding/json"
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

		// if message == "show dbs" {
		if microdbCommon.SHOW_DBS == c.Command {
			listDbs(conn, microdbCommon.SHOW_DBS)
		} else if microdbCommon.SHOW_TABLES == c.Command {
			listTables(conn, c)
		} else if microdbCommon.DB_EXISTS_USE_DB == c.Command {
			dbExists(conn, c)
		} else if microdbCommon.DB_EXISTS == c.Command {
			dbExists(conn, c)
		} else if microdbCommon.CREATE_TABLE == c.Command {
			createTable(conn, c)
		} else {
			sendCommandResponse(conn, c.Command, "Unrecognised command")
		}
	}
}

func sendCommandResponse(conn net.Conn, cmd string, response string) {
	serverResponse, _ := json.Marshal(microdbCommon.ServerResponse{Command: cmd, Response: response})
	sendResponse(conn, string(serverResponse))
}

func sendResponse(conn net.Conn, response string) {
	conn.Write([]byte(response + string(Delimiter)))
}
