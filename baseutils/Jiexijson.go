package baseutils

import (
	"encoding/json"
	"hanfuxin/basemodels"
	"io/ioutil"
	"log"
	"strings"
)

func Shezhijson() *basemodels.Shezhi {
	shezhi := basemodels.Shezhi{}
	obj := Jiexi(basemodels.Shezhipath(), &shezhi)
	return obj.(*basemodels.Shezhi)
}
func Guojihuajson(yuyan string) *basemodels.Guojihua {
	guojihua := basemodels.Guojihua{}
	obj := Jiexi(basemodels.Guojihuapath(strings.ToLower(yuyan)), &guojihua)
	return obj.(*basemodels.Guojihua)
}
func Jiexi(path string, model interface{}) interface{} {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取json失败", err)
	}
	json.Unmarshal(bytes, &model)
	return model
}
