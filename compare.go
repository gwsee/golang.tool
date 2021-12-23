package common_tools

import (
	"errors"
	"reflect"
)

/*
IsSames
	o is same as n;
	o is a pointer:old struct;
	n is a pointer:new struct;

	tag is from struct field Tag

	m = 0
	compare o with n and ,return n‘s all tag->val map;
	m = 1
	compare o with n and ,return n‘s  tag->val map which not return zero val;
	m = 2
	compare o with n and ,return n‘s  tag->val map which val is changed;
	m = 3
	compare o with n and ,return n‘s  tag->val map which val is changed and ignore zero val;

	m = 4
	compare o with n and ignore o zero val,return n‘s  all tag->val map;
	m = 5
	compare o with n and ignore o zero val,return n‘s  all tag->val map and ignore zero val;
	m = 6
	compare o with n and ignore o zero val,return n‘s  all tag->val map which val is changed;
	m = 7
	compare o with n and ignore o zero val,return n‘s  all tag->val map which val is changed and ignore zero val;

	m = 8
	compare o with n and ignore n zero val,return n‘s  all tag->val map;
	m = 9
	compare o with n and ignore n zero val,return n‘s  all tag->val map and ignore zero val;
	m = 10
	compare o with n and ignore n zero val,return n‘s  all tag->val map which val is changed;
	m = 11
	compare o with n and ignore n zero val,return n‘s  all tag->val map which val is changed and ignore zero val;

	m = 12
	compare o with n and ignore each zero val,return n‘s  all tag->val map;
	m = 13
	compare o with n and ignore each zero val,return n‘s  all tag->val map and ignore zero val;
	m = 14
	compare o with n and ignore each zero val,return n‘s  all tag->val map which val is changed;
	m = 15
	compare o with n and ignore each zero val,return n‘s  all tag->val map which val is changed and ignore zero val;

	b: o==n
	mp: tag->val
	err: error

*/
func IsSames(o, n interface{}, tag string, m int) (b bool, mp map[string]interface{}, err error) {
	if reflect.TypeOf(o).Kind() != reflect.Ptr {
		err = errors.New("o is not pointer")
		return
	}
	if reflect.TypeOf(n).Kind() != reflect.Ptr {
		err = errors.New("n is not pointer")
		return
	}
	mo := make(map[string]interface{})
	mn := make(map[string]interface{})
	ResolveVal(o, mo, tag, "", false)
	ResolveVal(n, mn, tag, "", false)
	switch m {
	case 0:
		return same0(mo, mn)
	case 1:
		return same1(mo, mn)
	case 2:
		return same2(mo, mn)
	case 3:
		return same3(mo, mn)
	case 4:
		return same4(mo, mn)
	case 5:
		return same5(mo, mn)
	case 6:
		return same6(mo, mn)
	case 7:
		return same7(mo, mn)
	case 8:
		return same8(mo, mn)
	case 9:
		return same9(mo, mn)
	case 10:
		return same10(mo, mn)
	case 11:
		return same11(mo, mn)
	case 12:
		return same12(mo, mn)
	case 13:
		return same13(mo, mn)
	case 14:
		return same14(mo, mn)
	case 15:
		return same15(mo, mn)
	}
	err = errors.New("m is not in schedule")
	return
}

//m = 0
//compare o with n and ,return n‘s all tag->val map;
func same0(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	m = n
	b = true
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					return
				}
				break
			}
		}
	}
	return
}

//m = 1
//compare o with n and ,return n‘s  tag->val map which not return zero val;
func same1(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	b = true
l:
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					break l
				}
				break
			}
		}
	}
	k := make(map[string]interface{})
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		k[kn] = vn
	}
	m = k
	return
}

//m = 2
//compare o with n and ,return n‘s  tag->val map which val is changed;
func same2(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					k[kn] = vn
				}
				break
			}
		}
	}
	m = k
	return
}

//m = 3
//compare o with n and ,return n‘s  tag->val map which val is changed and ignore zero val;
func same3(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					if !isBlank(reflect.ValueOf(vn)) {
						k[kn] = vn
					}
				}
				break
			}
		}
	}
	m = k
	return
}

//m = 4
//compare o with n and ignore o zero val,return n‘s  all tag->val map;
func same4(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	m = n
	b = true
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					return
				}
				break
			}
		}
	}
	return
}

//m = 5
//compare o with n and ignore o zero val,return n‘s  all tag->val map and ignore zero val;
func same5(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	b = true
l:
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					break l
				}
				break
			}
		}
	}
	k := make(map[string]interface{})
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		k[kn] = vn
	}
	m = k
	return
}

//m = 6
//compare o with n and ignore o zero val,return n‘s  all tag->val map which val is changed;
func same6(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					k[kn] = vn
				}
				break
			}
		}
	}
	m = k
	return
}

//m = 7
//compare o with n and ignore o zero val,return n‘s  all tag->val map which val is changed and ignore zero val;
func same7(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					if !isBlank(reflect.ValueOf(vn)) {
						k[kn] = vn
					}
				}
				break
			}
		}
	}
	m = k
	return
}

//
//m = 8
//compare o with n and ignore n zero val,return n‘s  all tag->val map;
func same8(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	m = n
	b = true
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					return
				}
				break
			}
		}
	}
	return
}

//m = 9
//compare o with n and ignore n zero val,return n‘s  all tag->val map and ignore zero val;
func same9(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	b = true
l:
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					break l
				}
				break
			}
		}
	}
	k := make(map[string]interface{})
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		k[kn] = vn
	}
	m = k
	return
}

//m = 10
//compare o with n and ignore n zero val,return n‘s  all tag->val map which val is changed;
func same10(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					k[kn] = vn
				}
				break
			}
		}
	}
	m = k
	return
}

//m = 11
//compare o with n and ignore n zero val,return n‘s  all tag->val map which val is changed and ignore zero val;
func same11(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if vn != vo {
					b = false
					if !isBlank(reflect.ValueOf(vn)) {
						k[kn] = vn
					}
				}
				break
			}
		}
	}
	m = k
	return
}

//
//m = 12
//compare o with n and ignore each zero val,return n‘s  all tag->val map;
func same12(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	m = n
	b = true
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					return
				}
				break
			}
		}
	}
	return
}

//m = 13
//compare o with n and ignore each zero val,return n‘s  all tag->val map and ignore zero val;
func same13(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	b = true
l:
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					break l
				}
				break
			}
		}
	}
	k := make(map[string]interface{})
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		k[kn] = vn
	}
	m = k
	return
}

//m = 14
//compare o with n and ignore each zero val,return n‘s  all tag->val map which val is changed;
func same14(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					k[kn] = vn
				}
				break
			}
		}
	}
	m = k
	return
}

//m = 15
//compare o with n and ignore each zero val,return n‘s  all tag->val map which val is changed and ignore zero val;
func same15(o, n map[string]interface{}) (b bool, m map[string]interface{}, err error) {
	k := make(map[string]interface{})
	b = true
	for kn, vn := range n {
		if isBlank(reflect.ValueOf(vn)) {
			continue
		}
		for ko, vo := range o {
			if kn == ko {
				if isBlank(reflect.ValueOf(vo)) {
					break
				}
				if vn != vo {
					b = false
					if !isBlank(reflect.ValueOf(vn)) {
						k[kn] = vn
					}
				}
				break
			}
		}
	}
	m = k
	return
}
