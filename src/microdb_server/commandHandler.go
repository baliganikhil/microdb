package main

import (
	"encoding/json"
	"microdb_common"
	"net"

	"github.com/mitchellh/mapstructure"
)

func getCommand(cmdJSON string) microdbCommon.Command {
	var c microdbCommon.Command
	c = c.FromJson([]byte(cmdJSON))
	return c
}

func handle_SHOW_DBS(conn net.Conn, command microdbCommon.Command) {
	dbs := getDBInfo().DBs
	var dbList []string

	for _, db := range dbs {
		dbName := db.Name
		dbList = append(dbList, dbName)
	}

	dbResponse := microdbCommon.DBListResponse{DBs: dbList}
	dbListJson, _ := json.Marshal(dbResponse)

	sendCommandResponse(conn, command.Command, string(dbListJson))
}

func handle_SHOW_TABLES(conn net.Conn, command microdbCommon.Command) {
	dbInfo := getDBInfo()
	var tableList []string

	var cmdListTables microdbCommon.CmdListTables
	mapstructure.Decode(command.Params, &cmdListTables)

	dbName := cmdListTables.DB

	for dbIndex := range dbInfo.DBs {
		db := &dbInfo.DBs[dbIndex]

		if db.Name == dbName {

			for tableIndex := range db.Tables {
				table := &db.Tables[tableIndex]
				tableList = append(tableList, table.Name)
			}
		}
	}

	tableResponse := microdbCommon.TableListResponse{Tables: tableList}
	tableListJson, _ := json.Marshal(tableResponse)

	// sendResponse(conn, strings.Join(tableList, "\n"))
	sendCommandResponse(conn, command.Command, string(tableListJson))
}

func useDB(conn net.Conn, params microdbCommon.Command) {
	dbName := params.Params.(string)
	sendResponse(conn, dbName)
}

func handle_CREATE_TABLE(conn net.Conn, command microdbCommon.Command) {
	var tableInfo microdbCommon.CmdCreateTable
	mapstructure.Decode(command.Params, &tableInfo)

	dbName := tableInfo.DB
	tableName := tableInfo.TableName

	dbInfo := getDBInfo()
	tableFound := false

	for dbIndex := range dbInfo.DBs {
		db := &dbInfo.DBs[dbIndex]

		if db.Name == dbName {

			for tableIndex := range db.Tables {
				table := &db.Tables[tableIndex]

				if table.Name == tableName {
					tableFound = true
					break
				}
			}

			if !tableFound {
				db.Tables = append(db.Tables, Table{Name: tableName})
				setDBInfo(dbInfo)
			}

		}

		if tableFound {
			break
		}
	}

	if !tableFound {
		sendCommandResponse(conn, command.Command, "Table "+tableName+" has been successfully created")
	} else {
		sendCommandResponse(conn, command.Command, "Table "+tableName+" already exists")
	}
}

func handle_DB_EXISTS(conn net.Conn, command microdbCommon.Command) {
	var cmdDbExists microdbCommon.CmdDBExists
	mapstructure.Decode(command.Params, &cmdDbExists)

	dbName := cmdDbExists.DB
	dbInfo := getDBInfo()
	dbFound := false

	for dbIndex := range dbInfo.DBs {
		db := &dbInfo.DBs[dbIndex]

		if db.Name == dbName {
			dbFound = true
			break
		}
	}

	dbExistsResponse := microdbCommon.DBExistsResponse{DB: dbName, Exists: dbFound}
	dbExistsResponseJson, _ := json.Marshal(dbExistsResponse)

	sendCommandResponse(conn, command.Command, string(dbExistsResponseJson))

}
