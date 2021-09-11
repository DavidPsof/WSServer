package logger

import (
	"WSServer/config"
	"github.com/subchen/go-log"
	"github.com/subchen/go-log/writers"
	"strings"
)

// Init - initializes the logger
func Init() {
	cfg := config.Get()

	lvl, err := log.ParseLevel(cfg.Log.Level)
	if err != nil {
		lvl = log.INFO
	}

	log.Default.Level = lvl

	if strings.TrimSpace(cfg.Log.FileName) == "" {
		return
	}

	if cfg.Log.MaxCountFile == 0 {
		cfg.Log.MaxCountFile = 10
	}

	log.Default.Out = &writers.DailyFileWriter{
		Name:     cfg.Log.FileName,
		MaxCount: cfg.Log.MaxCountFile,
	}

	log.Infof("log level: %s", log.Default.Level.String())
}
