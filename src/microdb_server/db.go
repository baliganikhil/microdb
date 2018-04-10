package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func getDBFilename() string {
	config := GetServerConfig()
	baseFolder := config.Folders.Data
	dbFilename := baseFolder + "/db.json"
	return dbFilename
}

func initDB() {
	fmt.Println("Initialising DB")
	dbFilename := getDBFilename()

	if _, err := os.Stat(dbFilename); os.IsNotExist(err) {
		dbFile, err := os.Create(dbFilename)
		if err != nil {
			log.Fatal(err)
		}

		initDbFile := `{
    "dbs": [
        {
            "name": "test",
			"isActive": true,
            "tables": []
        }
    ]
}
`

		err = ioutil.WriteFile(dbFilename, []byte(initDbFile), 0666)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Creating DB File")
		dbFile.Close()
	}
}

func getDBInfo() DBInfo {
	fmt.Println("Getting DB Info")
	dbFile := getDBFilename()

	raw, err := ioutil.ReadFile(dbFile)
	if err != nil {
		fmt.Printf("Error while trying to read DB file: %s\n", dbFile)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c DBInfo
	json.Unmarshal(raw, &c)
	return c
}
