package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"net/smtp"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"v1.0/model"
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
func HasParam(params map[string]interface{}, key string) bool {
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

// Paginator 分页
func Paginator(data []map[string]string, page, limit int) map[string]interface{} {
	paginatorMap := make(map[string]interface{})
	totalItems := len(data)

	var before int
	var next int

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit))) //page总数

	if page <= 1 {
		page = 1
		before = 1
	} else {
		before = page - 1
	}
	next = page + 1

	start := (page - 1) * limit
	if start >= totalItems {
		start = totalItems
	}

	end := start + limit
	if end > totalItems {
		end = totalItems
	}

	if next > totalPages {
		next = totalPages
	}

	fmt.Println(start, end)

	paginatorMap["items"] = data[start:end]
	paginatorMap["current"] = page
	paginatorMap["limit"] = limit
	paginatorMap["totalItems"] = totalItems
	paginatorMap["totalPages"] = totalPages
	paginatorMap["first"] = 1
	paginatorMap["last"] = totalPages
	paginatorMap["before"] = before
	paginatorMap["next"] = next

	return paginatorMap
}

// ValidateToken 验证token
func ValidateToken(tokenStr string, w http.ResponseWriter, r *http.Request) (*model.Token, int) {

	sess := globalSessions.SessionStart(w, r)
	var token = model.NewToken()
	errorNo := 0

	sessionRes := sess.Get("token")
	fmt.Println("session:", sessionRes)
	if sessionRes == nil {

		if !token.Validate(tokenStr) {
			errorNo = 201
		}
	} else {
		token = sessionRes.(*model.Token)

		// token过期
		if time.Now().Unix()-token.CreatedAt > token.Expire {
			sess.Delete("token")
			errorNo = 201
		}
	}

	if errorNo == 0 {

		sess.Set("token", token)
		return token, 0
	}

	return nil, errorNo

}

// GetLimit return limit
func GetLimit() int {
	return 10
}

// Dial return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {

	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}

// SendToMail 发送邮件
func SendToMail(to, subject, body, mailtype string) error {
	config := model.NewConfig()

	user := config.Get("mail_user")
	password := config.Get("mail_pass")
	host := config.Get("mail_host")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := SendMailUsingTLS(host, auth, user, sendTo, msg)
	return err
}

// Grade 奖励元气(action:write,comment)
func Grade(userID int64, action string, num int) bool {
	score := model.NewScore()

	score.UserID = userID
	score.Action = action
	score.Num = num
	score.Type = 1

	return score.Add()
}
