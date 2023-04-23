/**
  @author: Zero
  @date: 2023/4/23 11:22:02
  @desc: 日志组件配置

**/

package spoor

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Config Zap日志组件配置体
//
// LogDir  日志文件存放目录,默认为 [当前项目根目录/logs/]
// FileName 日志文件名。默认为 `app`
// LogLevel 日志级别。默认为INFO
// LogPrefix 日志前缀。
// LogWriterFile 是否启用日志文件持久化。默认为False
// LogWriterFromLevel 是否按照日志级别写入不同的日志文件。默认为false
// LogSplitTime 日志分割时间。 默认为24小时
// MaxSaveTime 日志文件最大保留时间。 默认为7天
// MaxFileSize 日志文件最大限制,超过后生成新的日志文件。 默认100mb
// Style 写入文件内的日志格式是否以Json格式。默认为false
// Color 终端日志级别是否高亮显示。默认为True
// RootPath 当前项目根目录
// Plugins ZapOptions插件选项
type Config struct {
	LogDir             string
	FileName           string
	LogLevel           zapcore.Level
	LogPrefix          string
	LogWriterFile      bool
	LogWriterFromLevel bool
	LogSplitTime       time.Duration
	MaxSaveTime        time.Duration
	MaxFileSize        int64
	Style              bool
	Color              bool
	RootPath           string
	Plugins            []zap.Option
}

// GetFileName 获取日志目录+文件名
// logs/xxx
func (opt Config) GetFileName() string {
	if !strings.HasSuffix(opt.LogDir, string(filepath.Separator)) {
		return fmt.Sprintf("%s%s%s", opt.LogDir, string(filepath.Separator), opt.FileName)
	} else {
		return fmt.Sprintf("%s%s", opt.LogDir, opt.FileName)
	}
}

// GetFileNameLevel 根据日志级别,获取日志目录+文件名
// logs/info/xxx
func (opt Config) GetFileNameLevel(level string) string {
	if !strings.HasSuffix(opt.LogDir, string(filepath.Separator)) {
		return fmt.Sprintf("%s%s%s%s%s", opt.LogDir, string(filepath.Separator), level, string(filepath.Separator), opt.FileName)
	} else {
		return fmt.Sprintf("%s%s%s%s", opt.LogDir, level, string(filepath.Separator), opt.FileName)
	}
}

// DefaultConfig 获取默认配置项
func DefaultConfig() *Config {
	// 获取当前工作主目录
	rootPath, _ := os.Getwd()
	rootPath = rootPath + string(filepath.Separator)
	return &Config{
		LogDir:             rootPath + "logs",
		FileName:           "app",
		LogLevel:           zapcore.DebugLevel,
		LogPrefix:          "[Spoor]",
		LogWriterFile:      false,
		LogWriterFromLevel: false,
		LogSplitTime:       time.Hour * 24,
		MaxSaveTime:        time.Hour * 24 * 7,
		MaxFileSize:        1024 * 1024 * 100,
		Style:              false,
		Color:              true,
		RootPath:           rootPath,
		Plugins:            []zap.Option{zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)},
	}
}
