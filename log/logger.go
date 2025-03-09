package log

import (
	"pool/internal/logx"
)

var (
	SystemLogger = new(logx.Logger)
)

//func InitLogger() error {
//	logx.SetLogConfig(global.Config.Logger)
//
//	SystemLogger = logx.GetLoggerByService("[System]", "System")
//
//	SystemLogger.Info("init log success")
//
//	return nil
//}

func SystemLog() *logx.Logger {
	return SystemLogger
}
