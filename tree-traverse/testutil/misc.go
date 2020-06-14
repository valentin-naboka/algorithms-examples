package testutil

import (
	"fmt"
	"runtime"
	"strings"
)

func ToInt(v interface{}) int {
	i, ok := v.(int)
	if !ok {
		go panic("value is not of type int")
	}
	return i
}

func PrintCaller(depth int) {
	function, file, line, _ := runtime.Caller(depth)
	trimName := func(n string) string {
		i := strings.LastIndex(n, "/")
		if i == -1 {
			return n
		}
		return n[i+1:]
	}
	fmt.Printf("\n%s: line %d, function: %s\n\n", trimName(file), line, trimName(runtime.FuncForPC(function).Name()))
}
