/**
  @author: Zero
  @date: 2023/3/30 13:02:16
  @desc: 日志组件实例化

**/

package spoor

import (
	"errors"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// 根据配置选项,创建一个Zap日志组件
func newLogger(opt *Config) (*Spoor, error) {
	// 创建Zap Encoder
	var encoder zapcore.Encoder
	if opt.Style {
		encoder = zapcore.NewJSONEncoder(consoleLoggerEncoder())
	} else {
		encoder = zapcore.NewConsoleEncoder(consoleLoggerEncoder())
	}
	// 创建Zap Core
	var err error
	var core zapcore.Core
	if !opt.LogWriterFile {
		// 日志不写入文件
		core = zapcore.NewCore(encoder, os.Stdout, opt.LogLevel)
	} else {
		// 日志文件持久化,创建多个Core
		core, err = newCores(opt, encoder)
	}
	if err != nil {
		return nil, err
	}
	if opt.WrapSkip > 0 {
		opt.Plugins = append(opt.Plugins, zap.AddCallerSkip(opt.WrapSkip))
	}
	// 创建Zap实例,并且注册插件
	logger := zap.New(core, opt.Plugins...)
	// zap全局实例也使用该实例
	zap.ReplaceGlobals(logger)
	return &Spoor{logger}, nil
}

// newCores 构建Zap Core
func newCores(opt *Config, encoder zapcore.Encoder) (zapcore.Core, error) {
	// 判断需要生成一或多个日志文件
	if opt.LogWriterFromLevel {
		// 详细记录 按照不同的日志级别,写入不到不同的日志文件中.只划分三个等级 info、debug、error。error文件中存储所有大于info级别的日志
		// 创建info级别Core
		infoCore, err := newLevelCore(zapcore.InfoLevel, opt)
		if err != nil {
			return nil, err
		}
		// 创建debug级别Core
		debugCore, err := newLevelCore(zapcore.DebugLevel, opt)
		if err != nil {
			return nil, err
		}
		// 创建Warn、Error、Panic、Fatal级别Core
		errorCore, err := newLevelCore(zapcore.ErrorLevel, opt)
		if err != nil {
			return nil, err
		}
		// 创建终端日志Core
		var stdoutCore zapcore.Core
		stdoutCore = zapcore.NewCore(encoder, os.Stdout, opt.LogLevel)
		cores := zapcore.NewTee(infoCore, debugCore, errorCore, stdoutCore)
		return cores, nil
	} else {
		// 所有日志记录到一个日志文件中
		// 创建日志文件流
		fileWriter, err := writerLoggerByFileSize(opt.GetFileName(), opt)
		if err != nil {
			return nil, err
		}
		// 合并输入流,将日志同时写到终端和文件中
		var fileCore zapcore.Core
		if opt.Style {
			fileCore = zapcore.NewCore(zapcore.NewJSONEncoder(fileLoggerEncoder()), fileWriter, opt.LogLevel)
		} else {
			fileCore = zapcore.NewCore(zapcore.NewConsoleEncoder(fileLoggerEncoder()), fileWriter, opt.LogLevel)
		}
		// 创建终端日志流
		stdoutCore := zapcore.NewCore(encoder, os.Stdout, opt.LogLevel)
		// 是否启用终端日志级别高亮
		// 合并为一个core
		core := zapcore.NewTee(fileCore, stdoutCore)
		return core, nil
	}
}

// newLevelCore 根据不同日志级别创建Zap Core
func newLevelCore(level zapcore.Level, opt *Config) (zapcore.Core, error) {
	// 创建日志文件流
	writer, err := writerLoggerByFileSize(opt.GetFileNameLevel(level.String()), opt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("create %s level file error \n", level.String()))
	}
	// 创建该文件对应的日志级别写入条件
	var condition zap.LevelEnablerFunc
	switch level {
	case zapcore.InfoLevel:
		condition = func(l zapcore.Level) bool {
			return l == zapcore.InfoLevel
		}
	case zapcore.DebugLevel:
		condition = func(l zapcore.Level) bool {
			return l == zapcore.DebugLevel
		}
	default:
		condition = func(l zapcore.Level) bool {
			return l >= zapcore.WarnLevel
		}
	}
	// 创建Core
	var core zapcore.Core
	if opt.Style {
		core = zapcore.NewCore(zapcore.NewJSONEncoder(fileLoggerEncoder()), writer, condition)
	} else {
		core = zapcore.NewCore(zapcore.NewConsoleEncoder(fileLoggerEncoder()), writer, condition)
	}
	return core, nil
}

// 日志输入流，按照日志文件大小分割
func writerLoggerByFileSize(fileName string, opt *Config) (zapcore.WriteSyncer, error) {
	writer := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", fileName),
		MaxSize:    int(opt.MaxFileSize),      // 文件最大写入限制, 单位MB
		MaxBackups: 30,                        // 最大保留日志文件数量
		MaxAge:     int(opt.MaxSaveTime / 24), // 最大保留日志文件天数
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(&writer), nil
}

// 日志输入流，按照日期分割写入
func writerLoggerByTime(fileName string, opt *Config) (zapcore.WriteSyncer, error) {
	// 日志文件名,加上根据日期时间后缀
	logFileName := fmt.Sprintf("%s.%s", fileName, "%Y-%m-%d.log")
	// 日志临时当前文件名
	_ = fmt.Sprintf("%s.log", fileName)
	// 创建日志文件
	write, err := rotatelogs.New(logFileName,
		//rotatelogs.WithLinkName(logTempFileName),      //生成正在写入的日志文件软链接,方便查看
		rotatelogs.WithRotationTime(opt.LogSplitTime), //日志切割时间间隔
		rotatelogs.WithMaxAge(opt.MaxSaveTime),        //日志最长保留时间
		rotatelogs.WithRotationSize(opt.MaxFileSize),  //日志文件最大限制
	)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(write), nil
}
