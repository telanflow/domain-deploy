package config

import (
	"github.com/telanflow/domain-deploy/infra/logger"
)

func GetLog() *logger.LogConfig {
	return &logger.LogConfig{
		Level:      gViper.GetString("logger.level"),
		Output:     gViper.GetString("logger.output"),
		Encoder:    gViper.GetString("logger.encoder"),
		File:       gViper.GetString("logger.file"),
		MaxSize:    gViper.GetInt("logger.maxSize"),
		MaxBackups: gViper.GetInt("logger.maxBackups"),
		MaxAge:     gViper.GetInt("logger.maxAge"),
		Compress:   gViper.GetBool("logger.compress"),
	}
}
