/**
  @author: Zero
  @date: 2023/4/2 13:36:19
  @desc: 创建日志组件实例

**/

package spoor

import "go.uber.org/zap"

type Spoor struct {
	*zap.Logger
}

func (sp *Spoor) Info(s string) {
	sp.Logger.Info(s)
}

func (sp *Spoor) InfoSf(s string) {
	sp.Sugar().Infof(s)
}

func (sp *Spoor) Debug(s string) {
	sp.Logger.Debug(s)
}

func (sp *Spoor) DebugSf(s string) {
	sp.Sugar().Debugf(s)
}

func (sp *Spoor) Error(s string) {
	sp.Logger.Error(s)
}

func (sp *Spoor) ErrorSf(s string) {
	sp.Sugar().Errorf(s)
}

// NewDefaultSpoor 构建默认的日志组件
func NewDefaultSpoor() (*Spoor, error) {
	return NewLogger(DefaultConfig())
}

// NewSpoor 创建日志组件
func NewSpoor(options ...Option) (*Spoor, error) {
	config := DefaultConfig()
	for _, option := range options {
		option(config)
	}
	return NewLogger(config)
}
