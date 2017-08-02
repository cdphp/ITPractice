package vendor

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	// comment
	_ "github.com/go-sql-driver/mysql"
)

// ModelInterface interface
type ModelInterface interface {
	Get()
	Add()
	Delete()
	Update()
}

// Model struct
type Model struct {
	Resource     string
	ModelManager *sql.DB
	QueryStr     string
}

// Init func
func (m *Model) Init(resource string) {
	m.Resource = resource
	m.InstanceDb()
}

// InstanceDb 初始化 db
func (m *Model) InstanceDb() {
	if m.ModelManager == nil {
		db, err := sql.Open("mysql", "root:hongker@/it_practice?charset=utf8")
		fmt.Println(err)

		if err != nil {
			panic(err)
		}

		m.ModelManager = db
	}

}

// Md5 加密
func (m *Model) Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// CloseDb func
func (m *Model) CloseDb() {
	defer m.ModelManager.Close()
}

// Find func
func (m *Model) Find() {

	fmt.Println("sql:", m.QueryStr)
	rows, _ := m.ModelManager.Query(m.QueryStr)
	cols, _ := rows.Columns()
	fmt.Println(cols)

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			fmt.Println(err)
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		fmt.Print(m)
	}
}

// SetConditions set conditions
func (m *Model) SetConditions(conditions map[string]string) *Model {
	sql := "select "

	// columns
	if _, ok := conditions["columns"]; ok {
		sql += conditions["columns"]
	} else {
		sql += " * "
	}

	sql += " from " + m.Resource

	if _, ok := conditions["conditions"]; ok {
		sql += " where " + conditions["conditions"]
	}

	if _, ok := conditions["order"]; ok {
		sql += " order by " + conditions["order"]
	}

	if _, ok := conditions["limit"]; ok {
		sql += " limit " + conditions["limit"]
	}

	m.QueryStr = sql
	return m
}
