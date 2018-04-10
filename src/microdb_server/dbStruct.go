package main

type Table struct {
	Name string `json:"name"`
}

type Database struct {
	Name     string  `json:"name"`
	IsActive bool    `json:"isActive"`
	Tables   []Table `json:"tables"`
}

type DBInfo struct {
	DBs []Database `json:"dbs"`
}
