package logger

import (
	"github.com/Zhousiru/NyaaHub/scheduler/lib/db"
)

func Error(collection, msg string) {
	db.AddLog(collection, db.LogTypeErr, msg)
}

func Info(collection, msg string) {
	db.AddLog(collection, db.LogTypeInfo, msg)
}
