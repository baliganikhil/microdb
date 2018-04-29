package microdbCommon

// COMMANDS
type CmdShowDbs struct{}

type CmdCreateTable struct {
	DB          string      `json:"db"`
	TableName   string      `json:"tableName"`
	TableSchema interface{} `json:"tableSchema"`
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

type CmdDropTable struct {
	DB        string `json:"db"`
	TableName string `json:"tableName"`
}

type CmdDBExists struct {
	DB string `json:"db"`
}

type CmdDescTable struct {
	DB        string `json:"db"`
	TableName string `json:"tableName"`
}

type CmdSaveRecord struct {
	DB        string      `json:"db"`
	TableName string      `json:"tableName"`
	Record    interface{} `json:"record"`
}

// RESPONSES
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

type DropTableResponse struct {
	DB        string `json:"db"`
	TableName string `json:"tableName"`
	Dropped   bool   `json:"dropped"`
}

type DescTableResponse struct {
	DB        string                 `json:"db"`
	TableName string                 `json:"tableName"`
	Schema    map[string]interface{} `json:"schema"`
}

type SaveRecordResponse struct {
	HasError  bool                   `json:"error"`
	DB        string                 `json:"db"`
	TableName string                 `json:"tableName"`
	Record    map[string]interface{} `json:"record"`
}
