package easy

import (
	"fmt"
	"reflect"
	"strings"
)

//IsIn val ptr，arr ptr; 判断val是不是在arr里面
func IsIn(val interface{},arr interface{}) (in bool) {
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
