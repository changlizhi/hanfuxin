package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdjueseservices"
	"testing"
	"log"
)

func TestJueseserviceTianjia(t *testing.T) {
	juese1 := appmodels.Juese{Id: 1, Bianma: "ROLE_JINGLI", Mingcheng: "经理", Biaoji: "Youxiao"}
	zdjueseservices.Tianjiajuese(&juese1)
}
func TestJueseserviceXiugai(t *testing.T) {
	juese := appmodels.Juese{Id: 1, Bianma: "", Mingcheng: "经理", Biaoji: "Youxiao"}
	zdjueseservices.Xiugaijuese(&juese)
}
func TestJueseserviceChaxun(t *testing.T){
	js := zdjueseservices.Chaxunjuese(1)
	log.Println(js)
}
func TestJueseserviceShanchu(t *testing.T){
	zdjueseservices.Shanchujuese(1)
}
