/**
  @author: Zero
  @date: 2023/4/2 13:36:19
  @desc: 创建日志组件实例

**/

package spoor

import "go.uber.org/zap"

// NewDefaultSpoor 默认的Spoor日志组件
func NewDefaultSpoor() (*zap.Logger, error) {
	return NewLogger(DefaultConfig())
}

// NewSpoor 创建日志组件
func NewSpoor(options ...Option) (*zap.Logger,error) {
	config := DefaultConfig()
	for _, option := range options {
		option(config)
	}
	return NewLogger(config)
}

