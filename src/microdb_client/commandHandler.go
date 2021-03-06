package main

import (
	"microdb_common"
	"regexp"
)

func commandParser(input string) microdbCommon.Command {

	// Check if create db
	createDBRegex, _ := regexp.Compile("^create[ ]+db[ ]+([a-zA-Z0-9]+)$")
	if createDBRegex.MatchString(input) {
		return parse_CREATE_DB(input, createDBRegex)
	}

	// Check if create table
	createTableRegex, _ := regexp.Compile("^db\\.([a-zA-Z0-9]+)\\.create\\((.*)\\)$")
	if createTableRegex.MatchString(input) {
		return parse_CREATE_TABLE(input, createTableRegex)
	}

	// Check if show dbs
	showDBsRegex, _ := regexp.Compile("^show[ ]+dbs$")
	if showDBsRegex.MatchString(input) {
		return parse_SHOW_DBS(input, createTableRegex)
	}

	// Check if show tables
	showTablesRegex, _ := regexp.Compile("^show[ ]+tables$")
	if showTablesRegex.MatchString(input) {
		return parse_SHOW_TABLES(input, showTablesRegex)
	}

	// Check if use db
	useDBRegex, _ := regexp.Compile("^use[ ]+db[ ]+([a-zA-Z0-9]+)$")
	if useDBRegex.MatchString(input) {
		return parse_DB_EXISTS_USE_DB(input, useDBRegex)
	}

	// Check if use db
	dBExistsRegex, _ := regexp.Compile("^db[ ]+([a-zA-Z0-9]+)$")
	if dBExistsRegex.MatchString(input) {
		return parse_DB_EXISTS(input, dBExistsRegex)
	}

	// Check if drop db
	dropDBRegex, _ := regexp.Compile("^drop[ ]+db[ ]+([a-zA-Z0-9]+)$")
	if dropDBRegex.MatchString(input) {
		return parse_DROP_DB(input, dropDBRegex)
	}

	// Check if drop table
	dropTableRegex, _ := regexp.Compile("^db\\.([a-zA-Z0-9]+)\\.drop\\(\\)$")
	if dropTableRegex.MatchString(input) {
		return parse_DROP_TABLE(input, dropTableRegex)
	}

	// Describe table
	descTableRegex, _ := regexp.Compile("^db\\.([a-zA-Z0-9]+)\\.describe\\(\\)$")
	if descTableRegex.MatchString(input) {
		return parse_DESC_TABLE(input, descTableRegex)
	}

	// Save Record
	saveRecordRegex, _ := regexp.Compile("^db\\.([a-zA-Z0-9]+)\\.save\\((.*)\\)$")
	if saveRecordRegex.MatchString(input) {
		return parse_SAVE_RECORD(input, saveRecordRegex)
	}

	return microdbCommon.Command{}
}

func parse_CREATE_DB(input string, createDBRegex *regexp.Regexp) microdbCommon.Command {
	matches := createDBRegex.FindStringSubmatch(input)
	dbName := matches[1]
	dbInfo := microdbCommon.CmdCreateDB{DB: dbName}

	return microdbCommon.CreateCommand(microdbCommon.CREATE_DB, dbInfo)
}

func parse_CREATE_TABLE(input string, createTableRegex *regexp.Regexp) microdbCommon.Command {
	matches := createTableRegex.FindStringSubmatch(input)
	tableName := matches[1]
	tableSchema := matches[2]
	tableInfo := microdbCommon.CmdCreateTable{TableName: tableName, DB: getDB(), TableSchema: tableSchema}

	return microdbCommon.CreateCommand(microdbCommon.CREATE_TABLE, tableInfo)
}

func parse_SHOW_DBS(input string, showDBsRegex *regexp.Regexp) microdbCommon.Command {
	return microdbCommon.CreateCommand(microdbCommon.SHOW_DBS, microdbCommon.CmdShowDbs{})
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

func parse_DROP_DB(input string, dropDBRegex *regexp.Regexp) microdbCommon.Command {
	matches := dropDBRegex.FindStringSubmatch(input)
	dbName := matches[1]
	dropDbParams := microdbCommon.CmdDropDb{DB: dbName}
	return microdbCommon.CreateCommand(microdbCommon.DROP_DB, dropDbParams)
}

func parse_DROP_TABLE(input string, dropTableRegex *regexp.Regexp) microdbCommon.Command {
	matches := dropTableRegex.FindStringSubmatch(input)
	tableName := matches[1]
	dropTableParams := microdbCommon.CmdDropTable{DB: getDB(), TableName: tableName}
	return microdbCommon.CreateCommand(microdbCommon.DROP_TABLE, dropTableParams)
}

func parse_DESC_TABLE(input string, dropTableRegex *regexp.Regexp) microdbCommon.Command {
	matches := dropTableRegex.FindStringSubmatch(input)
	tableName := matches[1]
	descTableParams := microdbCommon.CmdDescTable{DB: getDB(), TableName: tableName}
	return microdbCommon.CreateCommand(microdbCommon.DESC_TABLE, descTableParams)
}

func parse_SAVE_RECORD(input string, saveRecordRegex *regexp.Regexp) microdbCommon.Command {
	matches := saveRecordRegex.FindStringSubmatch(input)
	tableName := matches[1]
	recordIn := matches[2]

	saveRecordParams := microdbCommon.CmdSaveRecord{DB: getDB(), TableName: tableName, Record: recordIn}
	return microdbCommon.CreateCommand(microdbCommon.SAVE_RECORD, saveRecordParams)
}

func setDB(db string) {
	curDBName = db
}

func getDB() string {
	return curDBName
}
