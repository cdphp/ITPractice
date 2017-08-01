package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"v1.0/vendor"
)

// Result struct
type Result struct {
	ErrorNo     int         `json:"errorNo"`
	ErrorMsg    string      `json:"errorMsg"`
	Data        interface{} `json:"data"`
	ResonseTime int64       `json:"responseTime"`
}

var globalSessions *vendor.Manager

func init() {
	globalSessions, _ = vendor.NewSessionManager("memory", "goSessionid", 3600)
	go globalSessions.GC()

}

// GetErrorMsg 根据no获取对应msg
func GetErrorMsg(no int) string {
	myConfig := new(vendor.Config)

	myConfig.InitConfig(getCurrentDirectory() + "/config/configs.ini")

	msg := myConfig.Read("error", strconv.Itoa(no))

	return msg
}

// HasParam 判断是否有key存在
func HasParam(params map[string]string, key string) bool {
	if _, ok := params[key]; ok {
		return true
	}
	return false
}

// JSONReturn func
func JSONReturn(w http.ResponseWriter, result *Result) {
	result.ErrorMsg = GetErrorMsg(result.ErrorNo)
	result.ResonseTime = time.Now().Unix()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

// getCurrentDirectory 获取当前路径
func getCurrentDirectory() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
