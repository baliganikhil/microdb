package main

type Record struct {
	Data map[string]interface{} `json:"data"`
}

type RecordValidationErrorType string
type RecordDatatype string

const (
	DATA_TYPE_MISMATCH        RecordValidationErrorType = "DATA_TYPE_MISMATCH"
	DATA_LOWER_LIMIT_ERROR    RecordValidationErrorType = "DATA_LOWER_LIMIT_ERROR"
	DATA_UPPER_LIMIT_ERROR    RecordValidationErrorType = "DATA_UPPER_LIMIT_ERROR"
	DATA_EMAIL_REGEX_MISMATCH RecordValidationErrorType = "DATA_EMAIL_REGEX_MISMATCH"
	DATA_TYPE_UNKNOWN         RecordValidationErrorType = "DATA_TYPE_UNKNOWN"
)

const (
	DATATYPE_EMAIL RecordDatatype = "email"
)

type RecordValidationError struct {
	HasError   bool                      `json:"error"`
	Field      string                    `json:"field"`
	ErrType    RecordValidationErrorType `json:"errType"`
	ErrMessage string                    `json:"errMessage"`
}

func (e *RecordValidationError) Error() string {
	return e.ErrMessage
}

func isStdDatatype(datatypeIn string) bool {
	switch datatypeIn {
	case "bool":
	case "string":
	case "int":
	case "int8":
	case "int16":
	case "int32":
	case "int64":
	case "uint":
	case "uint8":
	case "uint16":
	case "uint32":
	case "uint64":
	case "byte":
	case "float32":
	case "float64":
	case "complex64":
	case "complex128":
		return true
	}

	return false
}
