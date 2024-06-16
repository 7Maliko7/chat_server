package log

import (
	"github.com/7Maliko7/chat_server/internal/config"
	logger "github.com/7Maliko7/chat_server/pkg/log"
	"github.com/creasty/defaults"
)
var Log logger.Logger

func New(cfg *config.LogConfig) logger.Logger {
	switch(cfg.Driver){
	case "zerolog":
		Log =
	default:
		Log =
	}
	return Log
}
