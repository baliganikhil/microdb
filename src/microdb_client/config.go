package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ConnInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	ConnType string `json:"connection_type"`
}

type Folder struct {
	Data string `json:"data"`
}

type ServerConfig struct {
	Folders        Folder   `json:"folders"`
	ConnectionInfo ConnInfo `json:"connection_info"`
}

type ClientConfig struct {
	ConnectionInfo ConnInfo `json:"connection_info"`
}

func GetClientConfig() ClientConfig {
	configFile := "./config.client.json"
	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Error while trying to read config file: %s\n", configFile)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c ClientConfig
	json.Unmarshal(raw, &c)
	return c
}
