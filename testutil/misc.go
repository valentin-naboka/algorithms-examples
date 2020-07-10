package testutil

import (
	"runtime"
	"strings"
	"testing"
)

func ToInt(v interface{}) int {
	i, ok := v.(int)
	if !ok {
		go panic("value is not of type int")
	}
	return i
}

func PrintCaller(t *testing.T, depth int) {
	function, file, line, _ := runtime.Caller(depth)
	trimName := func(n string) string {
		i := strings.LastIndex(n, "/")
		if i == -1 {
			return n
		}
		return n[i+1:]
	}
	t.Logf("%s: line %d, function: %s\n", trimName(file), line, trimName(runtime.FuncForPC(function).Name()))
}
