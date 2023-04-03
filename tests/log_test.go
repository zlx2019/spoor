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
	logger, sugaredLogger, err := spoor.DefaultSpoor()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("INFO")
	logger.Debug("DEBUG")
	logger.Error("ERROR")
	sugaredLogger.Infof("username %s", "admin")
}
