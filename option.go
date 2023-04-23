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

// WithWriterFile 开启日志写入文件
func WithWriterFile() Option {
	return func(config *Config) {
		config.LogWriterFile = true
	}
}


// WithWriterFileFromLevel 开启根据日志级别写入不同日志文件
// DEBUG、INFO、ERROR共三个级别
func WithWriterFileFromLevel() Option {
	return func(config *Config) {
		config.LogWriterFromLevel = true
	}
}

// WithLogFileInfo 配置日志文件限制信息
// 多次时间分割一次文件
// 文件保留多长时间
func WithLogFileInfo(splitTIme,maxSaveTime time.Duration) Option{
	return func(config *Config) {
		config.LogSplitTime = splitTIme
		config.MaxSaveTime = maxSaveTime
	}
}

// WithJsonStyle 日志格式使用Json风格
func WithJsonStyle() Option {
	return func(config *Config) {
		// Json风格日志只写入日志文件内
		if config.LogWriterFile {
			config.Style = true
		}
	}
}

// WithPlugins 注册Zap选项插件
func WithPlugins(options ...zap.Option) Option{
	return func(config *Config) {
		config.Plugins = append(config.Plugins,options...)
	}
}