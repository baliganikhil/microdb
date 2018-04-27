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

func handle_CREATE_DB(conn net.Conn, command microdbCommon.Command) {
	dbInfo := getDBInfo()

	var cmdCreateDB microdbCommon.CmdCreateDB
	mapstructure.Decode(command.Params, &cmdCreateDB)

	dbName := cmdCreateDB.DB
	dbFound := false

	for dbIndex := range dbInfo.DBs {
		db := &dbInfo.DBs[dbIndex]

		if db.Name == dbName {
			dbFound = true
		}
	}

	if !dbFound {
		dbInfo.DBs = append(dbInfo.DBs, Database{Name: dbName})
		setDBInfo(dbInfo)
		sendCommandResponse(conn, command.Command, "Database "+dbName+" has been created")
	} else {
		sendCommandResponse(conn, command.Command, "Database "+dbName+" already exists")
	}

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
	tableSchemaStr := tableInfo.TableSchema

	var tableSchema map[string]interface{}
	errTableSchema := json.Unmarshal([]byte(tableSchemaStr.(string)), &tableSchema)

	// Check table schema
	if errTableSchema != nil {
		sendCommandResponse(conn, command.Command, "The table schema appears to be broken")
		return
	}

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
				db.Tables = append(db.Tables, Table{Name: tableName, Schema: tableSchema})
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

func handle_DROP_DB(conn net.Conn, command microdbCommon.Command) {
	var cmdDropDb microdbCommon.CmdDropDb
	mapstructure.Decode(command.Params, &cmdDropDb)

	dbName := cmdDropDb.DB
	dbInfo := getDBInfo()
	dbDropped := false

	for dbIndex := range dbInfo.DBs {
		db := &dbInfo.DBs[dbIndex]

		if db.Name == dbName {
			dbInfo.DBs = append(dbInfo.DBs[0:dbIndex], dbInfo.DBs[dbIndex+1:]...)
			dbDropped = true
			setDBInfo(dbInfo)
			break
		}
	}

	dbDropResponse := microdbCommon.DropDBResponse{DB: dbName, Dropped: dbDropped}
	dbDropResponseJson, _ := json.Marshal(dbDropResponse)

	sendCommandResponse(conn, command.Command, string(dbDropResponseJson))
}

func handle_DROP_TABLE(conn net.Conn, command microdbCommon.Command) {
	var cmdDropTable microdbCommon.CmdDropTable
	mapstructure.Decode(command.Params, &cmdDropTable)

	dbName := cmdDropTable.DB
	tableName := cmdDropTable.TableName
	dbInfo := getDBInfo()
	tableDropped := false

	for dbIndex := range dbInfo.DBs {
		db := &dbInfo.DBs[dbIndex]

		if db.Name == dbName {
			for tableIndex := range db.Tables {
				table := &db.Tables[tableIndex]

				if table.Name == tableName {
					db.Tables = append(db.Tables[0:tableIndex], db.Tables[tableIndex+1:]...)
					tableDropped = true
					setDBInfo(dbInfo)
					break
				}

			}
		}
	}

	tableDropResponse := microdbCommon.DropTableResponse{DB: dbName, TableName: tableName, Dropped: tableDropped}
	tableDropResponseJson, _ := json.Marshal(tableDropResponse)

	sendCommandResponse(conn, command.Command, string(tableDropResponseJson))
}
