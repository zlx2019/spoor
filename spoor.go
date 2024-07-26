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

func (sp *Spoor) InfoSf(s string, v ...interface{}) {
	sp.Sugar().Infof(s, v...)
}

func (sp *Spoor) Debug(s string) {
	sp.Logger.Debug(s)
}

func (sp *Spoor) DebugSf(s string, v ...interface{}) {
	sp.Sugar().Debugf(s, v...)
}

func (sp *Spoor) Error(s string) {
	sp.Logger.Error(s)
}

func (sp *Spoor) ErrorSf(s string, v ...interface{}) {
	sp.Sugar().Errorf(s, v...)
}

func (sp *Spoor) Panic(s string) {
	sp.Logger.Panic(s)
}

func (sp *Spoor) PanicSf(s string, v ...interface{}) {
	sp.Sugar().Panicf(s, v...)
}

func (sp *Spoor) Fatal(s string) {
	sp.Logger.Fatal(s)
}

func (sp *Spoor) FatalSf(s string, v ...interface{}) {
	sp.Sugar().Fatalf(s, v...)
}

// NewDefaultSpoor 构建默认的日志组件
func NewDefaultSpoor() (*Spoor, error) {
	return newLogger(DefaultConfig())
}

// NewSpoor 根据配置构建日志组件
func NewSpoor(config *Config) (*Spoor, error) {
	return newLogger(config)
}

// NewSpoorWithOptions 创建日志组件
func NewSpoorWithOptions(options ...Option) (*Spoor, error) {
	config := DefaultConfig()
	for _, option := range options {
		option(config)
	}
	return newLogger(config)
}
