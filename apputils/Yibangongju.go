package apputils

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
)

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
func Tianjiawenjian(path string, wenjianming string, neirong []byte) {
	os.MkdirAll(basemodels.Getapppath()+path, os.ModePerm)
	ioutil.WriteFile(basemodels.Getapppath()+path+wenjianming, neirong, os.ModePerm)
}
