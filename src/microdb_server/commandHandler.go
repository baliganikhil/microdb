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
		conn.Write([]byte(db.Name + "\n"))
	}

	conn.Write([]byte(string(Delimiter)))
}
