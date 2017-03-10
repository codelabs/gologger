package logger

import (
	"reflect"
	"testing"
)

const constBadSkipPostion = 4
const constGoodSkipPostion = 1

func TestNewCaller(t *testing.T) {

	var caller = newCaller(constBadSkipPostion)
	if caller != nil {
		t.Errorf("runtime stack empty at position %d. Unable to trap", constBadSkipPostion)
	}

	caller = newCaller(constGoodSkipPostion)
	assertEquality(t, "*logger.caller", reflect.TypeOf(caller).String())
}

func TestGetFile(t *testing.T) {
	var caller = newCaller(constGoodSkipPostion)
	assertEquality(t, "caller_test.go", caller.getFile())
}

func TestGetFunction(t *testing.T) {
	var caller = newCaller(constGoodSkipPostion)
	assertEquality(t, "github.com/codelabs/gologger.TestGetFunction", caller.getFunction())
}

func TestGetLine(t *testing.T) {
	var caller = newCaller(constGoodSkipPostion)
	assertEquality(t, 33, caller.getLine())
}
