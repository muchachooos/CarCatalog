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

type AddParamInCar struct {
	NewNum *string `json:"newNum"`
	RegNum string  `db:"reg_num" json:"regNum"`
	Mark   *string `db:"mark" json:"mark"`
	Model  *string `db:"model" json:"model"`
	Year   *int    `db:"year" json:"year"`
	Owner  *string `db:"owner" json:"owner"`
}

type Car struct {
	RegNum string  `db:"reg_num" json:"regNum"`
	Mark   *string `db:"mark" json:"mark"`
	Model  *string `db:"model" json:"model"`
	Year   *int    `db:"year" json:"year"`
	Owner  *string `db:"owner" json:"owner"`
}

type CarsFilter struct {
	RegNum *string `db:"reg_num"`
	Mark   *string `db:"mark"`
	Model  *string `db:"model"`
	Year   *int    `db:"year"`
	Owner  *string `db:"owner"`
}
