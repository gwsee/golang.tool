package tool

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//copy struct

/*
CopyFields
	f1:data which we need assignment
	f2:data source

	easy copy need promote
*/
func CopyFields(f1, f2 interface{}) (err error) {
	if reflect.TypeOf(f1).Kind() != reflect.Ptr {
		return errors.New("i is not pointer")
	}
	if reflect.TypeOf(f2).Kind() != reflect.Ptr {
		return errors.New("i is not pointer")
	}
	val1 := reflect.ValueOf(f1)
	ele1 := val1.Elem()
	ty1 := ele1.Type()
	val2 := reflect.ValueOf(f2)
	ele2 := val2.Elem()
	ty2 := ele2.Type()
	for k1 := 0; k1 < ele1.NumField(); k1++ {
		t1 := ty1.Field(k1)
		for k2 := 0; k2 < ele2.NumField(); k2++ {
			t2 := ty2.Field(k2)
			if t2.Name == t1.Name {
				v2 := ele2.Field(k2)
				if t2.Type == t1.Type {
					ele1.FieldByName(t1.Name).Set(reflect.ValueOf(v2.Interface()))
				} else if fmt.Sprintf("%v", t1.Type) == "string" &&
					fmt.Sprintf("%v", reflect.TypeOf(v2.Interface()).Kind()) == "struct" {
					vs, _ := json.Marshal(v2.Interface())
					ele1.FieldByName(t1.Name).Set(reflect.ValueOf(string(vs)))
				} else if fmt.Sprintf("%v", t1.Type) == "int" &&
					fmt.Sprintf("%v", t2.Type) != "int" {
					val := fmt.Sprintf("%v", v2.Interface())
					n, _ := strconv.ParseFloat(val, 64)
					ele1.FieldByName(t1.Name).Set(reflect.ValueOf(int(n)))
				} else if fmt.Sprintf("%v", t1.Type) == "float64" &&
					fmt.Sprintf("%v", t2.Type) != "float64" {
					val := fmt.Sprintf("%v", v2.Interface())
					n, _ := strconv.ParseFloat(val, 64)
					ele1.FieldByName(t1.Name).Set(reflect.ValueOf(n))
				}
				break
			}
		}
	}
	return
}
