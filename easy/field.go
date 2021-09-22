package easy

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//CopyFields i1:我们需要赋值的，i2数据来源，根据同样的名字进行赋值
func CopyFields(i1, i2 interface{}) (err error) {
	if reflect.TypeOf(i1).Kind() != reflect.Ptr {
		return errors.New("i is not pointer")
	}
	if reflect.TypeOf(i2).Kind() != reflect.Ptr {
		return errors.New("i is not pointer")
	}
	val1 := reflect.ValueOf(i1)
	ele1 := val1.Elem()
	ty1 := ele1.Type()
	for k1 := 0; k1 < ele1.NumField(); k1++ {
		t1 := ty1.Field(k1)
		val2 := reflect.ValueOf(i2)
		ele2 := val2.Elem()
		ty2 := ele2.Type()
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
				}
				break
			}
		}
	}
	return
}

//Resolve 根据 struct 将特点的tag 解析成一个map数据返回
func Resolve(ptr interface{}, m map[string]interface{}, f string, filter string) (err error) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return errors.New("i is not pointer")
	}
	val := reflect.ValueOf(ptr)
	elem := val.Elem()
	ty := elem.Type()
	for k := 0; k < elem.NumField(); k++ {
		v := elem.Field(k)
		t := ty.Field(k)
		if t.Anonymous {
			structResolve(v, t, m, f)
		} else {
			key := t.Name
			if f != "" {
				r, ok := t.Tag.Lookup(f)
				if !ok {
					continue
				}
				key = r
			}
			m[key] = v.Interface() //一般是第一层优先原则
		}
	}
	if filter != "" {
		fArr := strings.Split(filter, ",")
		for _, v := range fArr {
			delete(m, v)
		}
	}
	return err
}
func structResolve(val reflect.Value, field reflect.StructField, m map[string]interface{}, f string) {
	ty := field.Type
	num := val.NumField()
	for i := 0; i < num; i++ {
		v := val.Field(i)
		t := ty.Field(i)
		if t.Anonymous {
			structResolve(v, t, m, f)
		} else {
			key := t.Name
			if f != "" {
				r, ok := t.Tag.Lookup(f)
				if !ok {
					continue
				}
				key = r
			}
			_, ok := m[key]
			if !ok {
				m[key] = v.Interface()
			}
		}
	}
}

//ResolveTag 根据struct 将他的 比如common 和 json 返回一个m的对应关系 方便程序进行处理
func ResolveTag(ptr interface{}, m map[string]string, f1 string, f2 string) (err error) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return errors.New("i is not pointer")
	}
	val := reflect.ValueOf(ptr)
	elem := val.Elem()
	ty := elem.Type()
	for k := 0; k < elem.NumField(); k++ {
		v := elem.Field(k)
		t := ty.Field(k)
		if t.Anonymous {
			structResolveTag(v, t, m, f1, f2)
		} else {
			if f1 != "" && f2 != "" {
				r1, ok1 := t.Tag.Lookup(f1)
				if !ok1 {
					continue
				}
				r2, ok2 := t.Tag.Lookup(f2)
				if !ok2 {
					continue
				}
				if r1 != "" && r2 != "" {
					m[r1] = r2
				}
			}
		}
	}
	return err
}
func structResolveTag(val reflect.Value, field reflect.StructField, m map[string]string, f1 string, f2 string) {
	ty := field.Type
	num := val.NumField()
	for i := 0; i < num; i++ {
		v := val.Field(i)
		t := ty.Field(i)
		if t.Anonymous {
			structResolveTag(v, t, m, f1, f2)
		} else {
			if f1 != "" && f2 != "" {
				r1, ok1 := t.Tag.Lookup(f1)
				if !ok1 {
					continue
				}
				r2, ok2 := t.Tag.Lookup(f2)
				if !ok2 {
					continue
				}
				if r1 != "" && r2 != "" {
					_, ok := m[r1]
					if !ok {
						m[r1] = r2
					}
				}

			}

		}
	}
}