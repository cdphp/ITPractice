package vendor

import "fmt"

// Error 错误信息
type Error struct {
	No  int
	Msg string
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误信息：%s，编号：%d", e.Msg, e.No)
}
