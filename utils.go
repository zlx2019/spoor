// @Title utils.go
// @Description 工具函数
// @Author Zero - 2024/7/27 00:14:14

package spoor

import (
	"fmt"
	"path/filepath"
	"strings"
)

// GetFileName 获取日志文件路径
// ./logs/xxx
func (opt *Config) GetFileName() string {
	if !strings.HasSuffix(opt.LogDir, string(filepath.Separator)) {
		return fmt.Sprintf("%s%s%s", opt.LogDir, string(filepath.Separator), opt.FileName)
	} else {
		return fmt.Sprintf("%s%s", opt.LogDir, opt.FileName)
	}
}

// GetFileNameLevel 根据日志级别,获取日志文件路径
// ./logs/info/xxx
func (opt *Config) GetFileNameLevel(level string) string {
	if !strings.HasSuffix(opt.LogDir, string(filepath.Separator)) {
		return fmt.Sprintf("%s%s%s%s%s", opt.LogDir, string(filepath.Separator), level, string(filepath.Separator), opt.FileName)
	} else {
		return fmt.Sprintf("%s%s%s%s", opt.LogDir, level, string(filepath.Separator), opt.FileName)
	}
}
