# spoor
_**野兽的足迹.**_ 

一个日志组件库.
<hr>

目前支持的日志组件有:
- zap

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
    logger, err = spoor.NewLogger(&spoor.Config{
        LogDir:             "",
        FileName:           "",
        LogLevel:           0,
        LogPrefix:          "",
        LogWriterFile:      false,
        LogWriterFromLevel: false,
        LogSplitTime:       0,
        MaxSaveTime:        0,
        MaxFileSize:        0,
        Style:              false,
        Color:              false,
        RootPath:           "",
        Plugins:            nil,
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
