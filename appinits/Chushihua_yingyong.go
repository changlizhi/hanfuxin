package appinits

import (
	"encoding/json"
	"hanfuxin/appmodels"
	"hanfuxin/baseinits"
	"log"
)

var Yingyongzi appmodels.Yingyongzi

func init() {
	logone()
}
func logone() {
	jsonobj, err := json.Marshal(baseinits.Zifus)
	if err != nil {
		log.Println(err)
		log.Println("解析Zifus[]到json格式出错")
	}
	json.Unmarshal(jsonobj, &Yingyongzi)
}
