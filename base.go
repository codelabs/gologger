package logger

// Logger interface provides necessary structure
type Logger interface {
	ShowEpochTimeStamp(state bool)
	ShowCallerInfo(state bool)
	Info(v ...interface{}) string
	Infof(format string, v ...interface{}) string
	Infoln(v ...interface{}) string
}
