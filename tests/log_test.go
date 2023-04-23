/**
  @author: Zero
  @date: 2023/3/30 12:36:55
  @desc:

**/

package tests

import (
	"github.com/zlx2019/spoor"
	"testing"
)

func TestName(t *testing.T) {
	//option := spoor.DefaultConfig()
	//option.LogWriterFile = true
	//logger, err := spoor.NewDefaultSpoor()
	logger, err := spoor.NewSpoor(spoor.WithWriterFile(), spoor.WithWriterFileFromLevel())
	logger, err = spoor.NewLogger(&spoor.Config{
		LogDir:             "",
		FileName:           "",
		LogLevel:           0,
		LogPrefix:          "",
		LogWriterFile:      false,
		LogWriterFromLevel: false,
		LogSplitTime:       0,
		MaxSaveTime:        0,
		MaxFileSize:        0,
		Style:              false,
		Color:              false,
		RootPath:           "",
		Plugins:            nil,
	})
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugaredLogger := logger.Sugar()

	logger.Info("INFO")
	logger.Debug("DEBUG")
	logger.Error("ERROR")
	sugaredLogger.Infof("username %s", "admin")
}
