/**
  @author: Zero
  @date: 2023/4/2 13:36:19
  @desc: 日志组件配置项

**/

package spoor

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Options Zap日志组件实例化配置
//
// LogDir  日志文件存放目录,默认为 [当前项目根目录/logs/]
// FileName 日志文件名。默认为 `app`
// LogLevel 日志级别。默认为INFO
// LogPrefix 日志前缀。
// EnableFileSave 是否启用日志文件持久化。默认为False
// LevelRecording 是否按照日志级别写入不同的日志文件。默认为false
// MaxSaveTime 日志分割时间。 默认为24小时
// MaxSaveTime 日志文件最大保留时间。 默认为7天
// MaxFileSize 日志文件最大限制,超过后生成新的日志文件。 默认100mb
// Style 写入文件内的日志格式是否以Json格式。默认为false
// Color 终端日志级别是否高亮显示。默认为True
// RootPath 当前项目根目录
type Options struct {
	LogDir         string
	FileName       string
	LogLevel       zapcore.Level
	LogPrefix      string
	EnableFileSave bool
	LevelRecording bool
	LogSplitTime   time.Duration
	MaxSaveTime    time.Duration
	MaxFileSize    int64
	Style          bool
	Color          bool
	RootPath       string
}

// GetFileName 获取日志目录+文件名
// logs/xxx
func (opt Options) GetFileName() string {
	if !strings.HasSuffix(opt.LogDir, string(filepath.Separator)) {
		return fmt.Sprintf("%s%s%s", opt.LogDir, string(filepath.Separator), opt.FileName)
	} else {
		return fmt.Sprintf("%s%s", opt.LogDir, opt.FileName)
	}
}

// GetFileNameLevel 根据日志级别,获取日志目录+文件名
// logs/info/xxx
func (opt Options) GetFileNameLevel(level string) string {
	if !strings.HasSuffix(opt.LogDir, string(filepath.Separator)) {
		return fmt.Sprintf("%s%s%s%s%s", opt.LogDir, string(filepath.Separator), level, string(filepath.Separator), opt.FileName)
	} else {
		return fmt.Sprintf("%s%s%s%s", opt.LogDir, level, string(filepath.Separator), opt.FileName)
	}
}

// DefaultOption 获取默认配置项
func DefaultOption() *Options {
	// 获取当前工作主目录
	rootPath, _ := os.Getwd()
	rootPath = rootPath + string(filepath.Separator)
	return &Options{
		LogDir:         rootPath + "logs",
		FileName:       "app",
		LogLevel:       zapcore.DebugLevel,
		LogPrefix:      "[Spoor]",
		EnableFileSave: false,
		LevelRecording: false,
		LogSplitTime:   time.Hour * 24,
		MaxSaveTime:    time.Hour * 24 * 7,
		MaxFileSize:    1024 * 1024 * 100,
		Style:          false,
		Color:          true,
		RootPath:       rootPath,
	}
}
