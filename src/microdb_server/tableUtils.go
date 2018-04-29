package main

import (
	"fmt"
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

			val, recordHasValue := recordIn[key]
			if !recordHasValue {
				continue
			}

			valDataType := reflect.TypeOf(val).Kind()

			if valDataType.String() != reqdDatatype {
				// Raise error
				errorMessage := "Invalid data type for field: " + key + ". Required: " + reqdDatatype.(string) + ", Actual: " + valDataType.String()
				return Record{}, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_TYPE_MISMATCH, ErrMessage: errorMessage}
			}

		} else {
			// Complex information
			// Complex object has been provided
			var keyInfoMap map[string]interface{} = make(map[string]interface{})
			keyInfoMap = keyInfo.(map[string]interface{})

			val, recordHasValue := recordIn[key]
			if !recordHasValue {
				continue
			}

			// Check datatype
			if reqdDatatype, hasDataTypeDef := keyInfoMap["type"]; hasDataTypeDef {
				// reqdDatatype, hasDataTypeDef := keyInfoMap["type"]
				valDataType := reflect.TypeOf(val).Kind()

				if !hasDataTypeDef {
					fmt.Println("Could not find value type for key: " + key)
					continue
				}

				if valDataType.String() != reqdDatatype.(string) {
					// Raise error
					errorMessage := "Invalid data type for field: " + key + ". Required: " + reqdDatatype.(string) + ", Actual: " + valDataType.String()
					return Record{}, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_TYPE_MISMATCH, ErrMessage: errorMessage}
				}
			}

			// Check lower limit
			if lowerLimit, hasLowerLimitDef := keyInfoMap["lowerLimit"]; hasLowerLimitDef {
				if val.(float64) < lowerLimit.(float64) {
					// Raise error
					errorMessage := "Invalid value for field: " + key + ". Lower Limit: " + FloatToString(lowerLimit.(float64)) + ", Actual: " + FloatToString(val.(float64))
					return Record{}, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_LOWER_LIMIT_ERROR, ErrMessage: errorMessage}
				}
			}

			// Check lower limit
			if upperLimit, hasUpperLimitDef := keyInfoMap["upperLimit"]; hasUpperLimitDef {
				if val.(float64) > upperLimit.(float64) {
					// Raise error
					errorMessage := "Invalid value for field: " + key + ". Upper Limit: " + FloatToString(upperLimit.(float64)) + ", Actual: " + FloatToString(val.(float64))
					return Record{}, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_UPPER_LIMIT_ERROR, ErrMessage: errorMessage}
				}
			}
		}
	}

	return Record{Data: recordIn}, nil
}
