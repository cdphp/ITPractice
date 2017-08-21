package lib

import (
	"encoding/base64"
	"errors"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
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

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
)

func SaveImageToDisk(fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}

	reader, err := base64.StdEncoding.DecodeString(data[idx+8:])

	ioutil.WriteFile(fileNameBase, reader, 0644)

	return fileNameBase, err
}
