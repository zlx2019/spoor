package spoor

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// 日志级别颜色
const (
	infoLevelColor  = "\033[92m[INFO]\033[0m"
	debugLevelColor = "\033[93m[DEBUG]\033[0m"
	warnLevelColor  = "\033[93m[WARN]\033[0m"
	errorLevelColor = "\033[91m[ERROR]\033[0m"
	panicLevelColor = "\033[91m[PANIC]\033[0m"
	fatalLevelColor = "\033[91m[FATAL]\033[0m"
)

// Zap 日志格式化编码器

// 输出在终端上风格的日志编码器
func consoleLoggerEncoder() zapcore.EncoderConfig {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = timeFormater             // 日志时间自定义格式化
	encoder.EncodeLevel = terminalLogLevelEncoder // 终端日志级别自定义处理
	encoder.EncodeCaller = callerEncoder          // 日志输出位置处理
	encoder.ConsoleSeparator = "  "               // 日志行分隔符
	return encoder
}

// 输出在日志文件中的风格编码器
func fileLoggerEncoder() zapcore.EncoderConfig {
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = timeFormater         // 日志时间自定义格式化
	encoder.EncodeLevel = fileLogLevelEncoder // 终端日志级别自定义处理
	encoder.EncodeCaller = callerEncoder      // 日志输出位置处理
	encoder.ConsoleSeparator = "  "           // 日志行分隔符
	encoder.TimeKey = "time"                  //日志时间字段
	encoder.CallerKey = "line"                //代码行号字段
	return encoder
}

// 日志级别格式化
// 不同级别使用不同颜色，易于区分
func terminalLogLevelEncoder(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		encoder.AppendString(infoLevelColor)
	case zapcore.DebugLevel:
		encoder.AppendString(debugLevelColor)
	case zapcore.WarnLevel:
		encoder.AppendString(warnLevelColor)
	case zapcore.ErrorLevel:
		encoder.AppendString(errorLevelColor)
	case zapcore.PanicLevel:
		encoder.AppendString(panicLevelColor)
	case zapcore.FatalLevel:
		encoder.AppendString(fatalLevelColor)
	default:
		encoder.AppendString(level.String())
	}
}

// 输出在文件中的日志级别格式化
func fileLogLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		enc.AppendString("[DEBUG]")
	case zapcore.InfoLevel:
		enc.AppendString("[INFO]")
	case zapcore.WarnLevel:
		enc.AppendString("[WARN]")
	case zapcore.ErrorLevel:
		enc.AppendString("[ERROR]")
	case zapcore.PanicLevel:
		enc.AppendString("[PANIC]")
	case zapcore.FatalLevel:
		enc.AppendString("[FATAL]")
	default:
		enc.AppendString(fmt.Sprintf("[%d]", level))
	}
}

// 日志时间格式化
func timeFormater(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Local().Format("2006-01-02 15:04:05.000"))
}

// 日志输出位置格式化
func callerEncoder(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(caller.TrimmedPath())
}
