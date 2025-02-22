package initialize

import (
	"log"

	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("初始化logger失败....")
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	return logger
}
