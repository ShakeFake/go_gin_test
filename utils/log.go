package utils

import (
	log "github.com/cihub/seelog"
	"os"
)

func InitLog() {
	logger, err := log.LoggerFromWriterWithMinLevel(os.Stdout, log.InfoLvl)
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
}
