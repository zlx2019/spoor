/**
  @author: Zero
  @date: 2023/3/30 12:36:55
  @desc:

**/

package tests

import (
	"github.com/zlx2019/spoor"
	"go.uber.org/zap"
	"testing"
)

func TestName(t *testing.T) {
	logger, err := spoor.newLogger(&spoor.Config{
		LogDir:             "./logs",
		FileName:           "test_app",
		LogLevel:           zap.DebugLevel,
		LogWriterFile:      true,
		LogWriterFromLevel: false,
		LogSplitTime:       0,
		MaxSaveTime:        0,
		MaxFileSize:        0,
		Style:              false,
		Plugins:            []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)},
		WrapSkip:           1,
	})
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugaredLogger := logger.Sugar()

	logger.Info("INFO")
	logger.Debug("DEBUG")
	logger.Error("ERROR")
	logger.Warn("Warn")
	sugaredLogger.Infof("username %s", "admin")
}
