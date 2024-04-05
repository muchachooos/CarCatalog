package model

type Err struct {
	Error string `json:"error"`
}

type Config struct {
	Port   int    `json:"port"`
	DBConf DBConf `json:"DataBase"`
}

type DBConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dataBaseName"`
	Sslmode  string `json:"sslmode"`
}
