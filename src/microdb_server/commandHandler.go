package main

import (
	"microdb_common"
	"net"
)

func getCommand(cmdJSON string) microdbCommon.Command {
	var c microdbCommon.Command
	c = c.FromJson([]byte(cmdJSON))
	return c
}

func listDbs(conn net.Conn) {
	dbs := getDBInfo().DBs

	for _, db := range dbs {
		dbName := db.Name
		if db.IsActive {
			dbName += " [active]"
		}

		sendResponse(conn, dbName+"\n")
	}

	sendResponse(conn, string(Delimiter))
}

func listTables(conn net.Conn) {}

func useDB(conn net.Conn, params microdbCommon.Command) {
	dbName := params.Params.(string)
	sendResponse(conn, dbName)
}
