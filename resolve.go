package common_tools

import (
	"reflect"
	"strings"
)

//SrVal is used for struct to map
func SrVal(val reflect.Value, field reflect.StructField, m map[string]interface{}, f string, ig bool) {
	ty := field.Type
	num := val.NumField()
	for i := 0; i < num; i++ {
		v := val.Field(i)
		t := ty.Field(i)
		if t.Anonymous {
			SrVal(v, t, m, f, ig)
		} else {
			if ig && isBlank(v) {
				continue
			}
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

// ResolveVal
// is used for a struct return a map which  key is tag,value is struct value ;
// ptr = struct data, m = return data,f = tag name,filter = ignore tag, ig = true = ignore zero val；false back all
func ResolveVal(ptr interface{}, m map[string]interface{}, f string, filter string, ig bool) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic("ptr is not pointer")
	}
	if reflect.TypeOf(ptr).Elem().Kind() != reflect.Struct {
		panic("ptr is not a struct")
	}
	val := reflect.ValueOf(ptr)
	elem := val.Elem()
	ty := elem.Type()
	for k := 0; k < elem.NumField(); k++ {
		v := elem.Field(k)
		t := ty.Field(k)
		if t.Anonymous {
			SrVal(v, t, m, f, ig)
		} else {
			if ig && isBlank(v) {
				continue
			}
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
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func SrTag(value reflect.Value, field reflect.StructField, m map[string]string, key string, val string) {
	ty := field.Type
	num := value.NumField()
	for i := 0; i < num; i++ {
		v := value.Field(i)
		t := ty.Field(i)
		if t.Anonymous {
			SrTag(v, t, m, key, val)
		} else {
			r1, ok1 := t.Tag.Lookup(key)
			if !ok1 {
				continue
			}
			r1 = strings.Split(r1, ",")[0]
			r2, ok2 := t.Tag.Lookup(val)
			if !ok2 {
				continue
			}
			r2 = strings.Split(r2, ",")[0]
			if r1 != "" && r2 != "" {
				_, ok := m[r1]
				if !ok {
					m[r1] = r2
				}
			}
		}
	}
}

//ResolveTag is used for a struct return a map which  key is tag1,value is tag2;
func ResolveTag(ptr interface{}, m map[string]string, key string, val string) {
	if key == "" || val == "" {
		panic("tg1 or tg2 is empty")
	}
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic("ptr is not pointer")
	}
	if reflect.TypeOf(ptr).Elem().Kind() != reflect.Struct {
		panic("ptr is not a struct")
	}
	value := reflect.ValueOf(ptr)
	elem := value.Elem()
	ty := elem.Type()
	for k := 0; k < elem.NumField(); k++ {
		v := elem.Field(k)
		t := ty.Field(k)
		if t.Anonymous {
			SrTag(v, t, m, key, val)
		} else {
			r1, ok1 := t.Tag.Lookup(key)
			if !ok1 {
				continue
			}
			r1 = strings.Split(r1, ",")[0]
			r2, ok2 := t.Tag.Lookup(val)
			if !ok2 {
				continue
			}
			r2 = strings.Split(r2, ",")[0]
			if r1 != "" && r2 != "" {
				m[r1] = r2
			}
		}
	}
}
