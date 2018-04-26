package microdbCommon

type CmdShowDbs struct{}

type CmdCreateTable struct {
	DB        string `json:"db"`
	TableName string `json:"tableName"`
}

type CmdCreateDB struct {
	DB string `json:"db"`
}

type CmdListTables struct {
	DB string `json:"db"`
}

type CmdDropDb struct {
	DB string `json:"db"`
}

type CmdDBExists struct {
	DB string `json:"db"`
}

type ServerResponse struct {
	Command  string      `json:"command"`
	Response interface{} `json:"response"`
}

type DBListResponse struct {
	DBs []string `json:"dbs"`
}

type TableListResponse struct {
	Tables []string `json:"tables"`
}

type DBExistsResponse struct {
	DB     string `json:"db"`
	Exists bool   `json:"exists"`
}

type DropDBResponse struct {
	DB      string `json:"db"`
	Dropped bool   `json:"dropped"`
}
