package baseutils

import (
	"encoding/json"
	"hanfuxin/basemodels"
	"io/ioutil"
	"log"
)

func Changliangjson() *basemodels.Changliang {
	changliang := basemodels.Changliang{}
	obj := Jiexi(basemodels.Getchangliangpath(), &changliang)
	return obj.(*basemodels.Changliang)
}
func Yingyongjson() *basemodels.Yingyong {
	yingyong := basemodels.Yingyong{}
	obj := Jiexi(basemodels.Getyingyongpath(), &yingyong)
	return obj.(*basemodels.Yingyong)
}
func Jiexi(path string, model interface{}) interface{} {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取json失败", err)
	}
	json.Unmarshal(bytes, &model)
	return model
}
