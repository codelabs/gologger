package logger

// SkipPosition defines which postion in caller stack that has caller information
type SkipPosition uint

// Logger interface provides necessary structure
type Logger interface {
	GetSkipPosition() SkipPosition
	SetSkipPosition(p SkipPosition)

	ShowEpochTimeStamp(state bool)
	ShowCallerInfo(state bool)

	Info(v ...interface{}) string
	Infof(format string, v ...interface{}) string
	Infoln(v ...interface{}) string
}
