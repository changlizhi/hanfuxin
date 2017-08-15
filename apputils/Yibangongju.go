package apputils

import (
	"errors"
	"reflect"
)

// 在三个类型中匹配字符串，匹配成功返回true，string,map,map-key
func Pipei3lei(zi interface{}, fu interface{}) (bool, error) {
	fustr := reflect.ValueOf(fu)
	switch reflect.TypeOf(fu).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < fustr.Len(); i++ {
			if fustr.Index(i).Interface() == zi {
				return true, nil
			}
		}
	case reflect.Map:
		if fustr.MapIndex(reflect.ValueOf(zi)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array!")
}
