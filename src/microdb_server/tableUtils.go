package main

import (
	"reflect"

	"github.com/rs/xid"
)

func generatePrimaryKey() string {
	guid := xid.New()
	return guid.String()
}

func validateRecordBeforeSave(recordIn map[string]interface{}, dbName string, tableName string) (Record, error) {
	dbs := getDBInfo().DBs
	var tableSchema map[string]interface{} = nil

	for _, db := range dbs {
		if db.Name == dbName {
			for _, table := range db.Tables {
				if table.Name == tableName {
					tableSchema = table.Schema
				}
			}
		}
	}

	if tableSchema == nil {
		// Raise error
	}

	for key, keyInfo := range tableSchema {
		// Simple information - Just the type has been provided
		if reflect.TypeOf(keyInfo).Kind() == reflect.String {
			reqdDatatype := keyInfo

			val := recordIn[key]
			valDataType := reflect.TypeOf(val).Kind()

			if valDataType.String() != reqdDatatype {
				// Raise error
				errorMessage := "Invalid data type for field: " + key + ". Required: " + reqdDatatype.(string) + ", Actual: " + valDataType.String()
				return Record{}, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_TYPE_MISMATCH, ErrMessage: errorMessage}
			}

		} else {
			// Complex information
			// Complex object has been provided
		}
	}

	return Record{Data: recordIn}, nil
}
