package logger

import (
	"github.com/AlexeyBMSTU/shop_backend/src/consts"
	"log"
)

func Info(msg string, args ...interface{}) {
	log.Println(string(consts.Cyan), "[INFO] ", msg, string(consts.Reset))
}
func Warn(msg string, args ...interface{}) {
	log.Println(string(consts.Yellow), "[WARN] ", msg, string(consts.Reset))
}
func Error(msg string, args ...interface{}) {
	log.Println(string(consts.Red), "[ERROR]", msg, string(consts.Reset))
}
