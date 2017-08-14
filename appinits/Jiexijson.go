package appinits

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

func Shezhijson() *Shezhi {
	shezhi := Shezhi{}
	obj := Jiexi(Shezhipath(), &shezhi)
	return obj.(*Shezhi)
}
func Guojihuajson(yuyan string) *Guojihua {
	guojihua := Guojihua{}
	obj := Jiexi(Guojihuapath(strings.ToLower(yuyan)), &guojihua)
	return obj.(*Guojihua)
}
func Jiexi(path string, model interface{}) interface{} {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取json失败", err)
	}
	json.Unmarshal(bytes, &model)
	return model
}
