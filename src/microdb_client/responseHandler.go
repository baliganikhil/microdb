package main

import (
	"encoding/json"
	"fmt"
	"microdb_common"
)

func handleResponse(responseIn string) {
	// fmt.Println(responseIn)
	var serverResponse microdbCommon.ServerResponse
	json.Unmarshal([]byte(responseIn), &serverResponse)

	if serverResponse.Command == microdbCommon.SHOW_DBS {
		show_SHOW_DBS_response(serverResponse)
	} else if serverResponse.Command == microdbCommon.SHOW_TABLES {
		show_SHOW_TABLES_response(serverResponse)
	} else if serverResponse.Command == microdbCommon.DB_EXISTS {
		show_DB_EXISTS_response(serverResponse)
	} else if serverResponse.Command == microdbCommon.DB_EXISTS_USE_DB {
		show_DB_EXISTS_USE_DB_response(serverResponse)
	} else if serverResponse.Command == microdbCommon.DROP_DB {
		show_DROP_DB_response(serverResponse)
	} else {
		fmt.Println(serverResponse.Response)
	}
}

func show_SHOW_DBS_response(serverResponse microdbCommon.ServerResponse) {
	var responseJSON microdbCommon.DBListResponse
	json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJSON)

	for _, db := range responseJSON.DBs {
		if db == getDB() {
			fmt.Println(db + "    [active]")
		} else {
			fmt.Println(db)
		}
	}
}

func show_SHOW_TABLES_response(serverResponse microdbCommon.ServerResponse) {
	var responseJSON microdbCommon.TableListResponse
	json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJSON)

	for _, table := range responseJSON.Tables {
		fmt.Println(table)
	}
}

func show_DB_EXISTS_response(serverResponse microdbCommon.ServerResponse) {
	var responseJSON microdbCommon.DBExistsResponse
	json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJSON)

	if responseJSON.Exists {
		fmt.Println("DB '" + responseJSON.DB + "' exists")
	} else {
		fmt.Println("DB '" + responseJSON.DB + "' does not exist")
	}
}

func show_DB_EXISTS_USE_DB_response(serverResponse microdbCommon.ServerResponse) {
	var responseJSON microdbCommon.DBExistsResponse
	json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJSON)

	if responseJSON.Exists {
		setDB(responseJSON.DB)
		fmt.Println("DB '" + responseJSON.DB + "' is now active")
	} else {
		fmt.Println("DB '" + responseJSON.DB + "' does not exist")
	}
}

func show_DROP_DB_response(serverResponse microdbCommon.ServerResponse) {
	var responseJSON microdbCommon.DropDBResponse
	json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJSON)

	if responseJSON.Dropped {
		fmt.Println("DB '" + responseJSON.DB + "' has been dropped")
	} else {
		fmt.Println("DB '" + responseJSON.DB + "' could not be dropped")
	}
}
