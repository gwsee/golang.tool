package common_tools

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestThreadData_Run(t *testing.T) {
	var td ThreadData
	var arr []int
	for k := 1; k <= 100000; k++ {
		arr = append(arr, k)
	}
	td.Line = 500
	td.Data = &arr
	td.RunTime = 10
	td.Interval = time.Second
	td.Func = d
	td.Run()
}
func d(i interface{}) (err error) {
	if i0, ok := i.(int); ok {
		if fmt.Sprintf("%v", i0) == time.Now().Format("05") {
			err = errors.New("dddd")
			return err
		}
	}
	return nil
}
