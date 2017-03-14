package logger

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const constSkipPosition = 3

// ConsoleLogger logs to console
type ConsoleLogger struct {
	// Flag var to toggle caller info - enabled by default
	callerInfo bool
	// Flag var to toggle timestamp as epoch or stringified - disabled by default
	epochTimestamp bool
	skipPosition   SkipPosition
	logger         *log.Logger
}

// GetSkipPosition ...
func (c *ConsoleLogger) GetSkipPosition() SkipPosition {
	return c.skipPosition
}

// SetSkipPosition ...
func (c *ConsoleLogger) SetSkipPosition(p SkipPosition) {
	c.skipPosition = p
}

func (c *ConsoleLogger) generateTimeStamp() string {

	var now = time.Now()
	if c.epochTimestamp {
		var nowNano = now.UnixNano()
		var nowMicro = int(nowNano) / 1000
		return strconv.Itoa(nowMicro)
	}

	var timestamp = fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d.%06d",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
		now.Nanosecond()/1000)
	return timestamp
}

func (c *ConsoleLogger) getCallerData() string {

	var callerData = ""
	if c.callerInfo {
		var clr = newCaller(constSkipPosition)
		callerData = fmt.Sprintf("%s %s %d", clr.getFile(), clr.getFunction(), clr.getLine())
	}
	return callerData
}

func (c *ConsoleLogger) generateLogContent(format string, v []interface{}) string {
	return fmt.Sprintf(format, v)
}

// NewConsoleLogger intializes new console logger which logs to STDOUT
// Enables caller info and has stringified timestamp
func NewConsoleLogger() Logger {

	// Initializing new logger to STDERR and we dont want to
	// have any prefix or flags set ot this, as we want to set
	// them later
	var c = &ConsoleLogger{
		callerInfo:     true,
		epochTimestamp: false,
		skipPosition:   constSkipPosition,
		logger:         log.New(os.Stderr, "", 0),
	}
	return c
}

// ShowEpochTimeStamp toggles between epoch timestamp vs stringified timestamp
func (c *ConsoleLogger) ShowEpochTimeStamp(state bool) {
	c.epochTimestamp = state
}

// ShowCallerInfo toggles caller info true/false
func (c *ConsoleLogger) ShowCallerInfo(state bool) {
	c.callerInfo = state
}

// Info logs with INFO level
func (c *ConsoleLogger) Info(v ...interface{}) string {

	var content = fmt.Sprintf(" INFO %s %s", c.getCallerData(), c.generateLogContent("%v", v))
	c.logger.SetPrefix(c.generateTimeStamp())
	c.logger.Print(content)
	return content
}

// Infoln logs with INFO level and a new line
func (c *ConsoleLogger) Infoln(v ...interface{}) string {

	var content = fmt.Sprintf(" INFO %s %s", c.getCallerData(), c.generateLogContent("%v", v))
	c.logger.SetPrefix(c.generateTimeStamp())
	c.logger.Println(content)
	return content
}

// Infof ...
func (c *ConsoleLogger) Infof(format string, v ...interface{}) string {

	var content = fmt.Sprintf(" INFO %s %s", c.getCallerData(), c.generateLogContent("%v", v))
	c.logger.SetPrefix(c.generateTimeStamp())
	c.logger.Printf(format, content)
	return content
}
