package helper

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFunctionName(i interface{}) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	split := strings.Split(funcName, "/")
	result := split[len(split)-1]
	result = strings.Replace(result, "-fm", "", -1)
	split = strings.Split(result, ".")
	return split[len(split)-1]
}
