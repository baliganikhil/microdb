package main

import (
	"microdb_common"
)

func ListDBs() string {
	return microdbCommon.CreateCommand(microdbCommon.SHOW_DBS, "").ToJson()
}

func setDB(db string) {
	curDBName = db
}

func getDB() string {
	return curDBName
}
