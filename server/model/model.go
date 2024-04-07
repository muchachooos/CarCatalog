package model

import "errors"

var ErrCarNotFound = errors.New("car not found")

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

type AddCarsReq struct {
	RegNums []string `json:"regNums"`
}

type Car struct {
	RegNum string  `db:"regNum" json:"regNum"`
	NewNum *string `db:"newNum" json:"newNum"`
	Mark   *string `db:"mark" json:"mark"`
	Model  *string `db:"model" json:"model"`
	Year   *int    `db:"year" json:"year"`
	Owner  *string `db:"owner" json:"owner"`
}
