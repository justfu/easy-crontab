package help

import (
	"fmt"
	"runtime"
)

// ErrorToString 错误类型转字符串
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		//error错误类型 发送钉钉发
		msg := v.Error()
		return msg
	default:
		return r.(string)
	}
}

// 打印堆栈信息
func PrintStackTrace() string {
	var str string
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		str += fmt.Sprintf("\n\n > %s:%d (0x%x)", file, line, pc)
	}
	return str
}
