package logger

import (
	"reflect"
	"testing"
)

func TestNewConsoleLogger(t *testing.T) {
	var c = NewConsoleLogger()
	if reflect.TypeOf(c).String() != "*logger.ConsoleLogger" {
		t.Error("Invalid type returned")
	}
}

func TestGetSkipPosition(t *testing.T) {
	var c = NewConsoleLogger()
	assertEquality(t, SkipPosition(constSkipPosition), c.GetSkipPosition())
}

func TestSetSkipPosition(t *testing.T) {
	var c = NewConsoleLogger()
	c.SetSkipPosition(4)
	assertEquality(t, SkipPosition(4), c.GetSkipPosition())
}

func TestConsoleLoggerInfo(t *testing.T) {

	var c = NewConsoleLogger()
	c.Info("woot", "hello", 10, "10", 12.45)

	c.ShowEpochTimeStamp(true)
	c.Info("foo", "bar")

	c.ShowCallerInfo(false)
	c.Info("biz", "baz")
}

func TestConsoleLoggerInfoln(t *testing.T) {
	var c = NewConsoleLogger()
	c.Infoln("")
}

func TestConsoleLoggerInfof(t *testing.T) {
	var c = NewConsoleLogger()
	c.Infof("%s", "")
}
