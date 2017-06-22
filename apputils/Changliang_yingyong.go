package apputils

import (
	"errors"
	"hanfuxin/basemodels"
	"hanfuxin/baseutils"
	"reflect"
)

func Changliangjson() *basemodels.Changliang {
	path := "conf/changliang.json"
	changliang := basemodels.Changliang{}
	obj := baseutils.Jiexi(path, &changliang)
	return obj.(*basemodels.Changliang)
}
func Yingyongjson() *basemodels.Yingyong {
	path := "conf/Yingyong.json"
	yingyong := basemodels.Yingyong{}
	obj := baseutils.Jiexi(path, &yingyong)
	return obj.(*basemodels.Yingyong)
}
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
