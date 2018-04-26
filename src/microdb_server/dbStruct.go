package main

type Table struct {
	Name string `json:"name"`
}

type Database struct {
	Name   string  `json:"name"`
	Tables []Table `json:"tables"`
}

type DBInfo struct {
	DBs []Database `json:"dbs"`
}
