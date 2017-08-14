package controllers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"v1.0/vendor"

	"github.com/jinzhu/gorm"
)

// Database init func
func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:hongker@/it_practice2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}

const (
	//BASE64字符表,不要有重复
	base64Table        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	hashFunctionHeader = "hong"
	hashFunctionFooter = "ker"
)

// Substr 截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

// Md5 加密
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// UniqueID 生成Guid字串
func UniqueID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

var coder = base64.NewEncoding(base64Table)

// Base64Encode base64加密
func Base64Encode(str string) string {
	var src []byte = []byte(hashFunctionHeader + str + hashFunctionFooter)
	return string([]byte(coder.EncodeToString(src)))
}

// Base64Decode base64解密
func Base64Decode(str string) (string, error) {
	var src []byte = []byte(str)
	by, err := coder.DecodeString(string(src))
	return strings.Replace(strings.Replace(string(by), hashFunctionHeader, "", -1), hashFunctionFooter, "", -1), err
}

// Now return unix time
func Now() int64 {
	return time.Now().Unix()
}

// GetMsg 根据no获取对应msg
func GetMsg(no int) string {
	myConfig := new(vendor.Config)

	myConfig.InitConfig(getCurrentDir() + "/configs/configs.ini")

	msg := myConfig.Read("error", strconv.Itoa(no))

	return msg
}

// getCurrentDir 获取当前路径
func getCurrentDir() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// IsEmail 验证邮箱格式
func IsEmail(email string) bool {
	b, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email)
	return b
}

// GetLimit return limit
func GetLimit() int {
	return 2
}
