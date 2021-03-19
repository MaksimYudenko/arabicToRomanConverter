package utils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n",
			filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

func Equals(tb testing.TB, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texpected: %#v\n\n\tgot: %#v\033[39m\n\n",
			filepath.Base(file), line, expected, actual)
		tb.FailNow()
	}
}

func LogValues(tb testing.TB, expected, actual interface{}) {
	tb.Logf("\nexpected value: %v\nactual value:\t%v", expected, actual)
}
