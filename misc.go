package searcher

import (
	"reflect"
	"regexp"
)

func Addr(p interface{}) uintptr {
	return reflect.Indirect(reflect.ValueOf(p)).UnsafeAddr()
}

var word = regexp.MustCompile(`[^@#\$%\^&\*\\\/\(\)\[\]\{\}\s]+`)
