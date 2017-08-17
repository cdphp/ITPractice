package lib

import (
	"fmt"
	"path/filepath"
	"strings"
)

// GetCurrentDir 获取当前路径
func GetCurrentDir() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
