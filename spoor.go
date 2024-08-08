/**
  @author: Zero
  @date: 2023/4/2 13:36:19
  @desc: 创建日志组件实例

**/

package spoor

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logLevel zapcore.Level

const (
	DEBUG = zapcore.DebugLevel
	INFO  = zapcore.InfoLevel
	WARN  = zapcore.WarnLevel
	ERROR = zapcore.ErrorLevel
	PANIC = zapcore.PanicLevel
	FATAL = zapcore.FatalLevel
)

type Spoor struct {
	*zap.Logger
}

func (sp *Spoor) Log(level zapcore.Level, s string, fields ...zap.Field) {
	sp.Logger.Log(level, s, fields...)
}

func (sp *Spoor) LogSf(level zapcore.Level, s string, v ...interface{}) {
	sp.Logger.Sugar().Logf(level, s, v...)
}

func (sp *Spoor) Info(s string, fields ...zap.Field) {
	sp.Logger.Info(s, fields...)
}

func (sp *Spoor) InfoSf(s string, v ...interface{}) {
	sp.Sugar().Infof(s, v...)
}

func (sp *Spoor) Debug(s string, fields ...zap.Field) {
	sp.Logger.Debug(s, fields...)
}

func (sp *Spoor) DebugSf(s string, v ...interface{}) {
	sp.Sugar().Debugf(s, v...)
}

func (sp *Spoor) Error(s string, fields ...zap.Field) {
	sp.Logger.Error(s, fields...)
}

func (sp *Spoor) ErrorSf(s string, v ...interface{}) {
	sp.Sugar().Errorf(s, v...)
}

func (sp *Spoor) Panic(s string, fields ...zap.Field) {
	sp.Logger.Panic(s, fields...)
}

func (sp *Spoor) DPanic(s string, fields ...zap.Field) {
	sp.Logger.DPanic(s, fields...)
}

func (sp *Spoor) PanicSf(s string, v ...interface{}) {
	sp.Sugar().Panicf(s, v...)
}

func (sp *Spoor) Fatal(s string, fields ...zap.Field) {
	sp.Logger.Fatal(s, fields...)
}

func (sp *Spoor) FatalSf(s string, v ...interface{}) {
	sp.Sugar().Fatalf(s, v...)
}

// NewDefaultSpoor 构建默认的日志组件
func NewDefaultSpoor() (*Spoor, error) {
	return newLogger(DefaultConfig())
}

// NewSpoor 根据配置构建日志组件
func NewSpoor(config *Config, options ...Option) (*Spoor, error) {
	for _, opt := range options {
		opt(config)
	}
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
