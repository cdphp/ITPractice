package vendor

import (
	"database/sql"
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

// CloseDb func
func (m *Model) CloseDb() {
	defer m.ModelManager.Close()
}

// SetConditions set conditions
func (m *Model) SetConditions(conditions map[string]string) string {
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

	return sql
}

func (m *Model) FindFirst(conditions map[string]string) map[string]string {
	sql := m.SetConditions(conditions)
	stmt, _ := m.ModelManager.Prepare(sql)
	defer stmt.Close()

	rows, _ := stmt.Query()
	defer rows.Close()

	if rows == nil {
		return nil
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil
	}

	rawResult := make([][]byte, len(cols))
	result := make(map[string]string, len(cols))
	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	if rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[cols[i]] = ""
			} else {
				result[cols[i]] = string(raw)
			}
		}

	} else {
		return nil
	}

	return result
}

func (m *Model) Find(conditions map[string]string) []map[string]string {
	sql := m.SetConditions(conditions)
	stmt, _ := m.ModelManager.Prepare(sql)
	defer stmt.Close()

	rows, _ := stmt.Query()
	defer rows.Close()

	if rows == nil {
		return nil
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil
	}

	rawResult := make([][]byte, len(cols))
	result := []map[string]string{}

	dest := make([]interface{}, len(cols))
	for i, _ := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			continue
		}
		item := make(map[string]string)
		for i, raw := range rawResult {

			if raw == nil {
				item[cols[i]] = ""
			} else {
				item[cols[i]] = string(raw)
			}
		}
		result = append(result, item)

	}
	//defer m.CloseDb()

	return result
}
