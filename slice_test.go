package tool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInArr(t *testing.T) {
	var i int
	i = 1
	fmt.Println(InArr(i, &i))
}
func TestAppendArr(t *testing.T) {
	i := 22
	arr := &[]int{1, 2, 3, 4}
	AppendArr(i, arr)
	for _, v := range *arr {
		fmt.Println(v)
	}
	fmt.Println(len(*arr))
}

//func TestAppendArr1(t *testing.T) {
//	var i int
//	i = 22
//	arr := &[]int{1, 2, 3, 4}
//	r := AppendArr1(i, arr).([]int)
//	fmt.Println(r)
//	for _, v := range r {
//		fmt.Println(v)
//	}
//}
func TestAppendArr2(t *testing.T) {
	var arr []int
	arr = append(arr, 1, 2, 3, 4)
	fmt.Println(arr)
	tool(&arr)
	fmt.Println(arr)
}
func tool(arr *[]int) {
	*arr = append(*arr, 33, 44)
	return
}

func TestSliceArr(t *testing.T) {
	var arr []int
	for k := 0; k < 15; k++ {
		arr = append(arr, k)
	}
	var dd [][]int
	dd = append(dd, arr)
	SliceArr(&arr, 4, &dd)
	for _, v := range dd {
		fmt.Println("---->", v)
	}
}

func TestSliceArrM(t *testing.T) {
	var arr []int
	for k := 0; k < 15; k++ {
		arr = append(arr, k)
	}
	r, _ := SliceArrM(&arr, 4)
	fmt.Println(r)
	ty := reflect.ValueOf(r)
	fmt.Printf("结果  r: %T\n", ty)
	rs, ok := r.Interface().([][]int)
	fmt.Println(ok)
	for _, v := range rs {
		fmt.Println("---->", v)
	}
}
func TestSliceArrM2(t *testing.T) {
	// 反射创建map slice channel
	intSlice := make([]int, 0)
	mapStringInt := make(map[string]int)
	sliceType := reflect.TypeOf(intSlice)
	mapType := reflect.TypeOf(mapStringInt)

	// 创建新值
	intSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	mapReflect := reflect.MakeMap(mapType)

	// 使用新创建的变量
	v := 10
	rv := reflect.ValueOf(v)
	intSliceReflect = reflect.Append(intSliceReflect, rv)
	intSlice2 := intSliceReflect.Interface().([]int)
	fmt.Println(intSlice2)

	k := "hello"
	rk := reflect.ValueOf(k)
	mapReflect.SetMapIndex(rk, rv)
	mapStringInt2 := mapReflect.Interface().(map[string]int)
	fmt.Println(mapStringInt2)
}
