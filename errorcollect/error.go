package errorcollect

import (
	"easy-crontab/help"
	"log"
)

func ErrorCollect(desc string) {
	var errMsg string
	if err := recover(); err != nil {
		errMsg = desc + "出现异常"
		errMsg += help.ErrorToString(err)
		trace := help.PrintStackTrace()
		log.Println(trace)
	}
}
