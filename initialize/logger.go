package initialize

import (
	"pool/global"
	"pool/internal/logx"
	"pool/log"
)

func InitLogger() error {
	logx.SetLogConfig(global.Config.Logger)

	log.SystemLogger = logx.GetLoggerByService("[System]", "System")

	log.SystemLog().Info("init log success")

	return nil
}
