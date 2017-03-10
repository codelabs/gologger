package logger

import (
	"path"
	"runtime"
)

type caller struct {
	file     string
	function string
	line     int
}

func newCaller(skip int) *caller {
	var pc, file, line, ok = runtime.Caller(skip)
	if !ok {
		return nil
	}

	var funcPtr = runtime.FuncForPC(pc)
	var funcName = ""

	if funcPtr != nil {
		funcName = funcPtr.Name()
	}

	var c = &caller{
		file:     path.Base(file),
		function: funcName,
		line:     line,
	}

	return c
}

// getFile returns the name of the file who invoked logger
func (c *caller) getFile() string {
	return c.file
}

// getFunction returns the name of the function who invoked logger
func (c *caller) getFunction() string {
	return c.function
}

// getLine returns the line number where the logger is called
func (c *caller) getLine() int {
	return c.line
}
