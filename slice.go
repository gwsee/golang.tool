package common_tools

//mainly for slice or array

import (
	"errors"
	"reflect"
)

// InArr is used for d where is in arr or slice i,which i is a pointer data,but d is not;
func InArr(d interface{}, i interface{}) bool {
	if d == nil {
		panic("data is nil")
	}
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		panic("arr is not pointer ")
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		panic("arr is not a slice")
	}
	t := reflect.ValueOf(d)
	v := reflect.ValueOf(i)
	e := v.Elem()
	for k := 0; k < e.Len(); k++ {
		if t.Interface() == e.Index(k).Interface() {
			return true
		}
	}
	return false
}

// InArrE is used for d where is in arr or slice i,which i is a pointer data,but d is not,which will return err;
func InArrE(d interface{}, i interface{}) (b bool, err error) {
	if d == nil {
		err = errors.New("data is nil")
		return
	}
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		err = errors.New("arr is not pointer ")
		return
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		err = errors.New("arr is not a slice")
		return
	}
	t := reflect.ValueOf(d)
	v := reflect.ValueOf(i)
	e := v.Elem()
	for k := 0; k < e.Len(); k++ {
		if t.Interface() == e.Index(k).Interface() {
			b = true
			return
		}
	}
	return
}

/*
	i := 22
	arr := &[]int{1, 2, 3, 4}
	AppendArr(i, arr)
	for _, v := range *arr {
		log.Println(v)
	}
1
2
3
4
22

if i.type is not same as arr will log.Print err and not panic
*/
func AppendArr(d interface{}, i interface{}) {
	if d == nil {
		panic("data is nil")
	}
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		panic("arr is not pointer ")
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		panic("arr is not a slice")
	}
	t := reflect.ValueOf(d)
	v := reflect.ValueOf(i)
	e := v.Elem()
	for k := 0; k < e.Len(); k++ {
		if t.Interface() == e.Index(k).Interface() {
			return
		}
	}
	e.Set(reflect.Append(e, reflect.ValueOf(d)))
	return
}

/*
	i := 22
	arr := &[]int{1, 2, 3, 4}
	err := AppendArrE(i, arr)
	if err != nil {
		return err
	}
	for _, v := range *arr {
		log.Println(v)
	}
1
2
3
4
22

if i.type is not same as arr will log.Print err and not panic
*/
func AppendArrE(d interface{}, i interface{}) (err error) {
	if d == nil {
		err = errors.New("data is nil")
		return
	}
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		err = errors.New("arr is not pointer ")
		return
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		err = errors.New("arr is not a slice")
		return
	}
	t := reflect.ValueOf(d)
	v := reflect.ValueOf(i)
	e := v.Elem()
	for k := 0; k < e.Len(); k++ {
		if t.Interface() == e.Index(k).Interface() {
			return
		}
	}
	e.Set(reflect.Append(e, reflect.ValueOf(d)))
	return
}

//SliceArr make slice to slice slice
func SliceArr(i interface{}, l int, b interface{}) {
	if l < 1 {
		panic("num is lower than 1")
	}
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		panic("arr is not pointer ")
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		panic("arr is not a slice")
	}
	v := reflect.ValueOf(i)
	e := v.Elem()
	p := e.Len() / l
	n := e.Len() % l
	if n > 0 {
		p = p + 1
	}
	z := reflect.MakeSlice(reflect.ValueOf(b).Elem().Type(), 0, 0)
	for k := 0; k < p; k++ {
		x := reflect.New(e.Type())
		y := x.Elem()
		s := k * l
		d := (k + 1) * l
		if d > e.Len() {
			d = e.Len()
		}
		for k1 := s; k1 < d; k1++ {
			y = reflect.Append(y, e.Index(k1))
		}
		z = reflect.Append(z, y)
	}
	reflect.ValueOf(b).Elem().Set(z)
	return
}

//SliceArrM make slice to slice&slice (this not return anything)
func SliceArrM(i interface{}, l int) (r reflect.Value, err error) {
	if l < 1 {
		err = errors.New("num is lower than 1")
		return
	}
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		err = errors.New("arr is not pointer ")
		return
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		err = errors.New("arr is not a slice")
		return
	}
	v := reflect.ValueOf(i)
	e := v.Elem()
	p := e.Len() / l
	n := e.Len() % l
	if n > 0 {
		p = p + 1
	}
	t := reflect.SliceOf(e.Type())
	r = reflect.MakeSlice(t, 0, 0)
	for k := 0; k < p; k++ {
		x := reflect.New(e.Type())
		y := x.Elem()
		s := k * l
		d := (k + 1) * l
		if d > e.Len() {
			d = e.Len()
		}
		for k1 := s; k1 < d; k1++ {
			y = reflect.Append(y, e.Index(k1))
		}
		r = reflect.Append(r, y)
	}
	return
}
