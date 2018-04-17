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

		var responseJson microdbCommon.DBListResponse
		json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJson)

		for _, db := range responseJson.DBs {
			if db == getDB() {
				fmt.Println(db + "    [active]")
			} else {
				fmt.Println(db)
			}
		}

	} else if serverResponse.Command == microdbCommon.SHOW_TABLES {
		// fmt.Println("serverResponse.Command: " + serverResponse.Command)

		var responseJson microdbCommon.TableListResponse
		json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJson)

		for _, table := range responseJson.Tables {
			fmt.Println(table)
		}

	} else if serverResponse.Command == microdbCommon.DB_EXISTS {
		var responseJson microdbCommon.DBExistsResponse
		json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJson)

		if responseJson.Exists {
			fmt.Println("DB '" + responseJson.DB + "' exists")
		} else {
			fmt.Println("DB '" + responseJson.DB + "' does not exist")
		}

	} else if serverResponse.Command == microdbCommon.DB_EXISTS_USE_DB {
		var responseJson microdbCommon.DBExistsResponse
		json.Unmarshal([]byte(serverResponse.Response.(string)), &responseJson)

		if responseJson.Exists {
			setDB(responseJson.DB)
			fmt.Println("DB '" + responseJson.DB + "' is now active")
		} else {
			fmt.Println("DB '" + responseJson.DB + "' does not exist")
		}

	} else {
		fmt.Println(serverResponse.Response)
	}
}
