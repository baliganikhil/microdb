package main

type Record struct {
	Data map[string]interface{} `json:"data"`
}

type RecordValidationErrorType string

const (
	DATA_TYPE_MISMATCH RecordValidationErrorType = "DATA_TYPE_MISMATCH"
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
