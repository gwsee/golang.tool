package easy

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestIsIn(t *testing.T) {
	var a int
	var b [4]int64
	a = 1
	b[0]=1
	b[2]=2
	b[3]=4
	fmt.Println(IsInSlice(&a,&b))
}

//IsIn val ptr，arr ptr; 判断val是不是在arr里面 ---- 需要进一步处理 这种不通用
func IsInSlice(val interface{},arr interface{}) (in bool) {
	if reflect.TypeOf(val).Kind() != reflect.Ptr {
		return
	}
	if reflect.TypeOf(arr).Kind() != reflect.Ptr {
		return
	}
	if !strings.HasPrefix(fmt.Sprintf("%T",arr),"*["){
		return
	}
	inStr,ok:=arr.(*[]string)
	if ok {
		for _,v:=range *inStr{
			if v == val{
				in = true
				return
			}
		}
	}
	inInt,ok:=arr.(*[]int)
	if ok {
		for _,v:=range *inInt{
			if v == val{
				in = true
				return
			}
		}
	}
	inItf,ok:=arr.(*[]interface{})
	if ok {
		for _,v:=range *inItf{
			if v == val{
				in = true
				return
			}
		}
	}
	return
}
