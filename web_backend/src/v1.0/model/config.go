package model

import "v1.0/vendor"

// Config 配置
type Config struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	Description  string `json:"description"`
	IsDelete     int    `json:"-"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
	vendor.Model `json:"-"`
}

// Configs array
type Configs []Config

// NewConfig 初始化
func NewConfig() *Config {
	c := new(Config)
	c.Init("configs") //设置表名

	return c
}

// Get 根据名称获取配置
func (c *Config) Get(name string) string {
	sql := "select id,value from " + c.Resource + " where name=?"

	err := c.ModelManager.QueryRow(sql, name).Scan(&c.ID, &c.Value)

	if err == nil {
		return c.Value
	}
	return ""
}
