package main

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func initLogger() {
	serverConfig := GetServerConfig()
	logDir := serverConfig.Folders.Logs
	logPath := logDir + "/microdb.log"

	flag.Parse()

	file, err := os.Create(logPath)

	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	// Log.Println("LogFile : " + logPath)
}
