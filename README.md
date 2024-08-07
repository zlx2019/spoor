# Spoor
_**野兽的足迹.**_ 

基于[Zap](https://github.com/uber-go/zap)日志框架简单封装的日志工具库.
<hr>


## 🚀 Install
```shell
go get github.com/zlx2019/spoor@latest
```
<hr>

## 💡Example
### use default logger
```go
func main() {
    logger, err := spoor.NewDefaultSpoor()
    if err != nil {
        panic(err)
    }
    defer logger.Sync() 
    // 获取SugarLogger
    sugaredLogger := logger.Sugar()
    // 输出日志
    logger.Info("INFO")
    logger.Debug("DEBUG")
    logger.Error("ERROR")
    sugaredLogger.Infof("username %s", "admin")
}
```
### use options logger
```go
func main() {
    // 开启日志写入文件,并且按照日志等级区分
    logger, err := spoor.NewSpoor(spoor.WithWriterFile(), spoor.WithWriterFileFromLevel())
    if err != nil {
        panic(err)
    }
    defer logger.Sync() 
    // 获取SugarLogger
    sugaredLogger := logger.Sugar()
    // 输出日志
    logger.Info("INFO")
    logger.Debug("DEBUG")
    logger.Error("ERROR")
    sugaredLogger.Infof("username %s", "admin")
}
```
### custom logger

```go
func main() {
    // LogDir  日志文件存放目录,默认为 [./logs]
    // FileName 日志文件名。默认为 `app`
    // Level 日志级别。默认为INFO
    // LogPrefix 日志前缀(暂无作用)。
    // WriteFile 是否将日志写入文件。
    // FileSeparate 是否按照日志级别写入不同的日志文件。默认为false
    // LogSplitTime 日志分割时间。 默认为24小时
    // MaxSaveTime 日志文件最大保留时间。 默认为7天
    // MaxFileSize 日志文件最大限制,超过后生成新的日志文件。 默认100mb
    // JsonStyle 写入文件内的日志格式是否以Json格式。默认为false
    // Plugins ZapOptions插件选项
    // WrapSkip 要省略的调用栈层
    logger, err = spoor.NewLogger(&spoor.Config{
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
	})
    if err != nil {
        panic(err)
    }
    defer logger.Sync()
    // 获取SugarLogger
    sugaredLogger := logger.Sugar()
    // 输出日志
    logger.Info("INFO")
    logger.Debug("DEBUG")
    logger.Error("ERROR")
    sugaredLogger.Infof("username %s", "admin")
}
```
