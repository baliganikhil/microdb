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

	} else {
		fmt.Println(serverResponse.Response)
	}
}
