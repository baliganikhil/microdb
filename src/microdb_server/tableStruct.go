package main

type Record struct {
	Data map[string]interface{} `json:"data"`
}

type RecordValidationErrorType string

const (
	DATA_TYPE_MISMATCH     RecordValidationErrorType = "DATA_TYPE_MISMATCH"
	DATA_LOWER_LIMIT_ERROR RecordValidationErrorType = "DATA_LOWER_LIMIT_ERROR"
	DATA_UPPER_LIMIT_ERROR RecordValidationErrorType = "DATA_UPPER_LIMIT_ERROR"
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
