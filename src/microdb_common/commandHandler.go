package microdbCommon

import (
	"encoding/json"
)

type Command struct {
	Command string      `json:"command"`
	Params  interface{} `json:"params"`
}

func (c Command) ToJson() string {
	j, _ := json.Marshal(c)
	return string(j)
}

func (c Command) FromJson(data []byte) Command {
	json.Unmarshal(data, &c)
	return c
}

func CreateCommand(cmd string, params interface{}) Command {
	return Command{cmd, params}
}
