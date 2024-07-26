/**
  @author: Zero
  @date: 2023/4/23 11:21:07
  @desc: 日志配置选项

**/

package spoor

import (
	"go.uber.org/zap"
	"time"
)

// Option 可选配置选项闭包函数
type Option func(*Config)

// WithWriteFile 开启日志写入文件
func WithWriteFile() Option {
	return func(config *Config) {
		config.WriteFile = true
	}
}

// WithFileSeparate 开启根据日志级别写入不同日志文件
// DEBUG、INFO、ERROR共三个级别
func WithFileSeparate() Option {
	return func(config *Config) {
		config.FileSeparate = true
	}
}

// WithLogFileInfo 配置日志文件限制信息
// 多次时间分割一次文件
// 文件保留多长时间
func WithLogFileInfo(splitTIme, maxSaveTime time.Duration, maxFileSize int) Option {
	return func(config *Config) {
	}
}

// WithTimeCutter 日志文件按照时间分割.
func WithTimeCutter(separateTime, maxAge time.Duration, maxFileSize int64) Option {
	return func(config *Config) {
		config.timeCutter = &timeCutter{
			SeparateTime: separateTime,
			MaxAge:       maxAge,
			MaxFileSize:  maxFileSize,
		}
	}
}

// WithFileSizeCutter 日志文件按照大小分割.
func WithFileSizeCutter(maxBackups, maxAge, maxFileSize int) Option {
	return func(config *Config) {
		config.fileSizeCutter = &fileSizeCutter{
			MaxBackups:  maxBackups,
			MaxAge:      maxAge,
			MaxFileSize: maxBackups,
		}
	}
}

// WithJsonStyle 日志格式使用Json风格
func WithJsonStyle() Option {
	return func(config *Config) {
		// Json风格日志只写入日志文件内
		if config.WriteFile {
			config.JsonStyle = true
		}
	}
}

// WithPlugins 注册Zap选项插件
func WithPlugins(options ...zap.Option) Option {
	return func(config *Config) {
		config.Plugins = append(config.Plugins, options...)
	}
}
