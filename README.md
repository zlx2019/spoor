# spoor
_**野兽的足迹.**_ 

 基于zap日志库实现.
<hr>

## 🚀 Install
```shell
go get github.com/zlx2019/spoor
```
<hr>

## 💡Example
```go
func main() {
    logger, sugaredLogger, err := spoor.DefaultSpoor()
    if err != nil {
        panic(err)
    }
    defer logger.Sync()
    logger.Info("INFO")
    logger.Debug("DEBUG")
    logger.Error("ERROR")
    sugaredLogger.Infof("username %s","admin")
}
```
or
```go
func main() {
    logger, sugaredLogger, err := spoor.NewSpoor(spoor.Options{
    LogDir:         "",
    FileName:       "",
    LogLevel:       0,
    LogPrefix:      "",
    EnableFileSave: false,
    LevelRecording: false,
    LogSplitTime:   0,
    MaxSaveTime:    0,
    MaxFileSize:    0,
    Style:          false,
    Color:          false,
    })
}
```
