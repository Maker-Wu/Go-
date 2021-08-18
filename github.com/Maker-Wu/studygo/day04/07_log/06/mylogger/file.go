package mylogger

// FileLogger 文件日志结构体
type FileLogger struct {
	Level LogLevel
}

// 判断是否需要记录该日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

// Debug ...
func (f *FileLogger) Debug(format string, a ...interface{}) {
	//f.log(DEBUG, format, a...)
}
