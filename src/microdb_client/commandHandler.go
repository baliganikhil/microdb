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
		return parse_CREATE_TABLE(input, createTableRegex)
	}

	// Check if show dbs
	showDBsRegex, _ := regexp.Compile("show[ ]+dbs")
	if showDBsRegex.MatchString(input) {
		return parse_SHOW_DBS(input, createTableRegex)
	}

	// Check if show tables
	showTablesRegex, _ := regexp.Compile("show[ ]+tables")
	if showTablesRegex.MatchString(input) {
		return parse_SHOW_TABLES(input, showTablesRegex)
	}

	// Check if use db
	useDBRegex, _ := regexp.Compile("use[ ]+db[ ]+([a-zA-Z0-9]+)")
	if useDBRegex.MatchString(input) {
		return parse_DB_EXISTS_USE_DB(input, useDBRegex)
	}

	// Check if use db
	dBExistsRegex, _ := regexp.Compile("db[ ]+([a-zA-Z0-9]+)")
	if dBExistsRegex.MatchString(input) {
		return parse_DB_EXISTS(input, dBExistsRegex)
	}

	return microdbCommon.Command{}
}

func parse_CREATE_TABLE(input string, createTableRegex *regexp.Regexp) microdbCommon.Command {
	matches := createTableRegex.FindStringSubmatch(input)
	tableName := matches[1]
	tableInfo := microdbCommon.CmdCreateTable{TableName: tableName, DB: getDB()}

	return microdbCommon.CreateCommand(microdbCommon.CREATE_TABLE, tableInfo)
}

func parse_SHOW_DBS(input string, showDBsRegex *regexp.Regexp) microdbCommon.Command {
	return microdbCommon.CreateCommand(microdbCommon.SHOW_DBS, "")
}

func parse_SHOW_TABLES(input string, showTablesRegex *regexp.Regexp) microdbCommon.Command {
	tableInfo := microdbCommon.CmdCreateTable{DB: getDB()}

	return microdbCommon.CreateCommand(microdbCommon.SHOW_TABLES, tableInfo)
}

func parse_DB_EXISTS_USE_DB(input string, useDBRegex *regexp.Regexp) microdbCommon.Command {
	matches := useDBRegex.FindStringSubmatch(input)
	dbName := matches[1]
	dbExistsParams := microdbCommon.CmdDBExists{DB: dbName}
	return microdbCommon.CreateCommand(microdbCommon.DB_EXISTS_USE_DB, dbExistsParams)
}

func parse_DB_EXISTS(input string, dBExistsRegex *regexp.Regexp) microdbCommon.Command {
	matches := dBExistsRegex.FindStringSubmatch(input)
	dbName := matches[1]
	dbExistsParams := microdbCommon.CmdDBExists{DB: dbName}
	return microdbCommon.CreateCommand(microdbCommon.DB_EXISTS, dbExistsParams)
}

func setDB(db string) {
	curDBName = db
}

func getDB() string {
	return curDBName
}

func createTable(conn net.Conn) {

}
