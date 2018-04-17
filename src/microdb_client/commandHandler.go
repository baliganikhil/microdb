package main

import (
	"microdb_common"
	"net"
	"regexp"
)

func commandParser(input string) microdbCommon.Command {

	// Check if create table
	createTableRegex, _ := regexp.Compile("db\\.([a-zA-Z0-9]+)\\.create\\(\\)")
	if createTableRegex.MatchString(input) {
		matches := createTableRegex.FindStringSubmatch(input)
		tableName := matches[1]
		tableInfo := microdbCommon.CmdCreateTable{TableName: tableName, DB: getDB()}

		return microdbCommon.CreateCommand(microdbCommon.CREATE_TABLE, tableInfo)
	}

	// Check if show dbs
	showDBsRegex, _ := regexp.Compile("show[ ]+dbs")
	if showDBsRegex.MatchString(input) {
		return microdbCommon.CreateCommand(microdbCommon.SHOW_DBS, "")
	}

	// Check if show tables
	showTablesRegex, _ := regexp.Compile("show[ ]+tables")
	if showTablesRegex.MatchString(input) {
		tableInfo := microdbCommon.CmdCreateTable{DB: getDB()}

		return microdbCommon.CreateCommand(microdbCommon.SHOW_TABLES, tableInfo)
	}

	// Check if use db
	useDBRegex, _ := regexp.Compile("use[ ]+db[ ]+([a-zA-Z0-9]+)")
	if useDBRegex.MatchString(input) {
		matches := useDBRegex.FindStringSubmatch(input)
		dbName := matches[1]
		dbExistsParams := microdbCommon.CmdDBExists{DB: dbName}
		return microdbCommon.CreateCommand(microdbCommon.DB_EXISTS_USE_DB, dbExistsParams)
	}

	// Check if use db
	dBExistsRegex, _ := regexp.Compile("db[ ]+([a-zA-Z0-9]+)")
	if dBExistsRegex.MatchString(input) {
		matches := dBExistsRegex.FindStringSubmatch(input)
		dbName := matches[1]
		dbExistsParams := microdbCommon.CmdDBExists{DB: dbName}
		return microdbCommon.CreateCommand(microdbCommon.DB_EXISTS, dbExistsParams)
	}

	return microdbCommon.Command{}
}

func setDB(db string) {
	curDBName = db
}

func getDB() string {
	return curDBName
}

func createTable(conn net.Conn) {

}
