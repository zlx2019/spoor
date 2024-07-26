/**
  @author: Zero
  @date: 2023/4/23 11:22:02
  @desc: 日志组件配置

**/

package spoor

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Config 日志配置
type Config struct {
	LogDir         string          // 日志文件存放目录,默认为 [./logs]
	FileName       string          // 日志级别。默认为INFO
	Level          zapcore.Level   // 日志前缀(暂无作用)。
	LogPrefix      string          // 日志前缀
	WriteFile      bool            // 日志是否写入文件。
	FileSeparate   bool            // 日志文件按级别分离
	JsonStyle      bool            // 写入文件内的日志格式是否以Json格式。默认为false
	Plugins        []zap.Option    // zap 选项
	WrapSkip       int             // 要省略的调用栈层
	timeCutter     *timeCutter     // 日志文件分割规则
	fileSizeCutter *fileSizeCutter // 日志文件分割规则
}

// TimeCutter 按照时间分割日志配置
type timeCutter struct {
	SeparateTime time.Duration // 日志多久分割一次，产生新的日志文件
	MaxAge       time.Duration // 日志文件能保留的最大时间
	MaxFileSize  int64         // 日志文件最大可写入的日志量(byte)
}

// FileSizeCutter 按照文件大小分割日志配置
type fileSizeCutter struct {
	MaxBackups  int // 最多保留的日志文件数量
	MaxAge      int // 日志文件最大能保留的天数
	MaxFileSize int // 日志文件最大可写入的日志量(mb)
}

// DefaultConfig 获取默认配置项
func DefaultConfig() *Config {
	return &Config{
		LogDir:       "./logs",
		FileName:     "app",
		Level:        zapcore.DebugLevel,
		WriteFile:    false,
		FileSeparate: false,
		JsonStyle:    false,
		Plugins:      []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)},
		WrapSkip:     0,
		fileSizeCutter: &fileSizeCutter{
			MaxBackups:  10,
			MaxAge:      30,
			MaxFileSize: 100,
		},
	}
}
