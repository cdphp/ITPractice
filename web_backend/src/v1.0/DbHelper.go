package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DbHelper db工具
type DbHelper struct {
	db *sql.DB
}

// Init function
func (helper *DbHelper) Init() {
	helper.InstanceDb()
}

// InstanceDb 初始化 db
func (helper *DbHelper) InstanceDb() {
	if helper.db == nil {
		db, err := sql.Open("mysql", "root:hongker@/it_practice?charset=utf8")
		fmt.Println(err)

		if err != nil {
			panic(err)
		}
		helper.db = db
	}

}
