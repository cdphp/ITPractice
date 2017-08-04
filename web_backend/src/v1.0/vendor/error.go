package vendor

import "fmt"

// Error 错误信息
type Error struct {
	Msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误信息：%s，", e.Msg)
}
