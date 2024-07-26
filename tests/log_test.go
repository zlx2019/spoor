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
	"time"
)

func TestName(t *testing.T) {
	logger, err := spoor.NewSpoor(&spoor.Config{
		LogDir:       "./logs",
		FileName:     "app",
		Level:        zap.DebugLevel,
		WriteFile:    true,
		FileSeparate: false,
		JsonStyle:    false,
		Plugins:      []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)},
		WrapSkip:     1,
	}, spoor.WithTimeCutter(time.Hour, time.Hour, 1024*1024*1024))
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
