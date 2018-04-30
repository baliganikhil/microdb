package main

import (
	"reflect"
	"regexp"
	"strconv"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func ValidateEmail(emailIn string) bool {
	emailRegex, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
	return emailRegex.Match([]byte(emailIn))
}

func ValidateNonStdDatatypes(datatypeInSchema string, key string, val interface{}) (bool, error, error) {
	customToNativeDatatypeMap := make(map[string]string)
	customToNativeDatatypeMap["email"] = "string"

	reqdDatatype, foundValue := customToNativeDatatypeMap[datatypeInSchema]
	if !foundValue {
		// error
		errorMessage := "You seem to have used an unknown datatype for key: " + key + ". Datatype used: " + datatypeInSchema
		return false, nil, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_TYPE_UNKNOWN, ErrMessage: errorMessage}
	}

	valDataType := reflect.TypeOf(val).Kind().String()
	isValid := valDataType == reqdDatatype

	if isValid {
		return true, nil, nil
	} else {
		errorMessage := "Invalid data type for field: " + key + ". Required: " + reqdDatatype + ", Actual: " + valDataType
		return false, &RecordValidationError{HasError: true, Field: key, ErrType: DATA_TYPE_MISMATCH, ErrMessage: errorMessage}, nil
	}
}
