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
	"os"
	"strings"
	"time"
)



// NewLogger 根据配置选项,创建一个Zap日志组件
func NewLogger(opt *Config) (*zap.Logger, error) {
	// 创建Zap Encoder
	var encoder zapcore.Encoder
	// 根据配置不同的日志风格,获取不同的日志编码器
	if opt.Style {
		encoder = newJsonEncoder(opt)
	} else {
		encoder = newTextEncoder(false, opt)
	}
	// 创建Zap Core
	var err error
	var core zapcore.Core
	if !opt.LogWriterFile {
		// 日志不写入文件
		core = zapcore.NewCore(newTextEncoder(opt.Color, opt), os.Stdout, opt.LogLevel)
	} else {
		// 日志文件持久化,创建多个Core
		core, err = newCores(opt, encoder)
	}
	if err != nil {
		return nil, err
	}
	// Err级别日志,打印堆栈错误信息
	//addStacktrace := zap.AddStacktrace(zap.ErrorLevel)
	// 开启文件及行号显示
	//addCallerOpt := zap.AddCaller()
	// 创建Zap实例,并且注册插件
	logger := zap.New(core,opt.Plugins...)
	// zap全局实例也使用该实例
	zap.ReplaceGlobals(logger)
	return logger, nil
}

// newCores 构建Zap Core
func newCores(opt *Config, encoder zapcore.Encoder) (zapcore.Core, error) {
	// 判断需要生成一或多个日志文件
	if opt.LogWriterFromLevel {
		// 详细记录 按照不同的日志级别,写入不到不同的日志文件中.只划分三个等级 info、debug、error。error文件中存储所有大于info级别的日志
		// 创建info级别Core
		infoCore, err := newLevelCore(zapcore.InfoLevel, opt, encoder)
		if err != nil {
			return nil, err
		}
		// 创建debug级别Core
		debugCore, err := newLevelCore(zapcore.DebugLevel, opt, encoder)
		if err != nil {
			return nil, err
		}
		// 创建Warn、Error、Panic、Fatal级别Core
		errorCore, err := newLevelCore(zapcore.ErrorLevel, opt, encoder)
		if err != nil {
			return nil, err
		}
		// 创建终端日志Core
		var stdoutCore zapcore.Core
		// 是否启用终端日志级别高亮
		if opt.Color {
			stdoutCore = zapcore.NewCore(newTextEncoder(true, opt), os.Stdout, opt.LogLevel)
		} else {
			stdoutCore = zapcore.NewCore(encoder, os.Stdout, opt.LogLevel)
		}
		cores := zapcore.NewTee(infoCore, debugCore, errorCore, stdoutCore)
		return cores, nil
	} else {
		// 所有日志记录到一个日志文件中
		fileWriter, err := newLoggerWriter(opt.GetFileName(), opt)
		if err != nil {
			return nil, err
		}
		// 合并输入流,将日志同时写到终端和文件中
		// 创建文件流Core,这里不使用Color高亮,会产生乱码
		fileCore := zapcore.NewCore(encoder, fileWriter, opt.LogLevel)
		// 创建终端文件流
		var stdoutCore zapcore.Core
		// 是否启用终端日志级别高亮
		if opt.Color {
			stdoutCore = zapcore.NewCore(newTextEncoder(true, opt), os.Stdout, opt.LogLevel)
		} else {
			stdoutCore = zapcore.NewCore(encoder, os.Stdout, opt.LogLevel)
		}
		// 合并为一个core
		core := zapcore.NewTee(fileCore, stdoutCore)
		return core, nil
	}
}

// newLevelCore 根据不同日志级别创建Zap Core
func newLevelCore(level zapcore.Level, opt *Config, encoder zapcore.Encoder) (zapcore.Core, error) {
	// 创建日志文件流
	writer, err := newLoggerWriter(opt.GetFileNameLevel(level.String()), opt)
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
	core := zapcore.NewCore(encoder, writer, condition)
	return core, nil
}

// newTextEncoder 构建一个Text风格日志编码器。color表示是否启用高亮颜色
func newTextEncoder(color bool, opt *Config) zapcore.Encoder {
	// 创建一个生产级别专用的EncoderConfig
	config := zap.NewProductionEncoderConfig()
	// 日志级别默认为小写`info`,转为大写`INFO`
	if color {
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	// 日志输出位置信息过滤,默认显示全部路径
	// 只输出相对路径即可,通过项目根目录去除多余的层级
	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ReplaceAll(caller.String(), opt.RootPath, ""))
	}
	// 日志时间格式化
	config.EncodeTime = datetimeFormat
	return zapcore.NewConsoleEncoder(config)
}

// newJsonEncoder 创建一个Json风格日志编码器
func newJsonEncoder(opt *Config) zapcore.Encoder {
	// 创建一个生产级别专用的EncoderConfig
	config := zap.NewProductionEncoderConfig()
	// 默认的Json key并不友好,定义为自己喜欢的标识
	config.TimeKey = "time"   //日志时间字段
	config.CallerKey = "line" //代码行号字段
	// 日志级别默认为小写`info`,转为大写`INFO`
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	// 日志输出位置信息过滤,默认显示全部路径
	// 只输出相对路径即可,通过项目根目录去除多余的层级
	config.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(strings.ReplaceAll(caller.String(), opt.RootPath, ""))
	}
	// 日志时间格式化
	config.EncodeTime = datetimeFormat
	return zapcore.NewJSONEncoder(config)
}

// 日志时间`time`字段格式化处理函数 2006-01-02 15:04:05
func datetimeFormat(t time.Time, e zapcore.PrimitiveArrayEncoder) {
	datetime := t.Local().Format("2006-01-02 15:04:05.000")
	e.AppendString(fmt.Sprintf("[%s]", datetime))
}

// newLoggerWriter 创建日志文件输入流
func newLoggerWriter(fileName string, opt *Config) (zapcore.WriteSyncer, error) {
	// 日志文件名,加上根据日期时间后缀
	logFileName := fmt.Sprintf("%s.%s", fileName, "%Y-%m-%d.log")
	// 日志临时当前文件名
	logTempFileName := fmt.Sprintf("%s.log", fileName)
	// 创建日志文件
	write, err := rotatelogs.New(logFileName,
		rotatelogs.WithLinkName(logTempFileName),      //生成正在写入的日志文件软链接,方便查看
		rotatelogs.WithRotationTime(opt.LogSplitTime), //日志切割时间间隔
		rotatelogs.WithMaxAge(opt.MaxSaveTime),        //日志最长保留时间
		rotatelogs.WithRotationSize(opt.MaxFileSize),  //日志文件最大限制
	)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(write), nil
}
