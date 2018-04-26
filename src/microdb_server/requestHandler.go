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
			sendResponse(conn, "ERROR 501: Something went wrong while trying to read request\n"+err.Error())
			return
		}

		fmt.Println("Command: " + message + "\n")
		message = strings.Trim(message, "\n")

		c := getCommand(message)

		Log.Println(c.ToJson())

		if microdbCommon.SHOW_DBS == c.Command {
			handle_SHOW_DBS(conn, c)
		} else if microdbCommon.SHOW_TABLES == c.Command {
			handle_SHOW_TABLES(conn, c)
		} else if microdbCommon.DB_EXISTS_USE_DB == c.Command {
			handle_DB_EXISTS(conn, c)
		} else if microdbCommon.DB_EXISTS == c.Command {
			handle_DB_EXISTS(conn, c)
		} else if microdbCommon.CREATE_DB == c.Command {
			handle_CREATE_DB(conn, c)
		} else if microdbCommon.CREATE_TABLE == c.Command {
			handle_CREATE_TABLE(conn, c)
		} else if microdbCommon.DROP_DB == c.Command {
			handle_DROP_DB(conn, c)
		} else if microdbCommon.DROP_TABLE == c.Command {
			handle_DROP_TABLE(conn, c)
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
