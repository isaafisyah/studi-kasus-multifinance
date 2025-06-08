package log

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)


var (
	loggers   = map[string]*logrus.Logger{}
	loggersMu sync.Mutex
)

func GetLogger(module string) *logrus.Logger {
	if module == "" {
		module = "app"
	}

	loggersMu.Lock()
	defer loggersMu.Unlock()

	// Cek cache
	if logger, ok := loggers[module]; ok {
		return logger
	}

	logPath := "storage/logs"
	_ = os.MkdirAll(logPath, os.ModePerm)

	filename := "app-" + time.Now().Format("2006-01-02") + ".log"
	logFile := filepath.Join(logPath, filename)

	logWriter := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    10,   // MB
		MaxBackups: 5,
		MaxAge:     7,    // days
		Compress:   true,
	}

	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(os.Stdout, logWriter))
	logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logger.SetLevel(logrus.InfoLevel)

	logger = logger.WithField("module", module).Logger
	return logger
}
