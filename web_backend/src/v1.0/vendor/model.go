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
