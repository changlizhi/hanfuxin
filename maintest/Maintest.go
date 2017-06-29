package main

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"hanfuxin/baserun"
	"hanfuxin/baseutils"
	"hanfuxin/daos"
	"hanfuxin/services"
	"log"
)

func teststringarr() {
	str := "abc"
	strarr := []string{"abcd", "def", "pfv"}
	log.Println(apputils.Pipei3lei(str, strarr))
	mpall := map[string]string{"abc": "123", "456": "def"}
	log.Println(apputils.Pipei3lei(str, mpall))
}
func testyingyong() {
	obj := baseutils.Yingyongjson()
	log.Println(obj, "\n")
}
func testchangliang() {
	obj := baseutils.Changliangjson()
	log.Println(obj, "\n")
}
func testshengcheng1() {
	baserun.Shengcheng_yingyongzi_model()
}
func testshengcheng2() {
	baserun.Shengcheng_yingyong_model()
}
func testappinits() {
	log.Println(appinits.Yingyongzi.Fuhao)
	log.Println(baseinits.Zhongwens)
	log.Println(baseinits.Yingwens)
}
func testjuesedao() {
	juese := appmodels.Juese{Id: 1, Bianma: "ROLE_ADMIN", Mingcheng: "管理员", Biaoji: "Youxiao"}
	daos.Tianjia_yige_juese(&juese)
}
func testjueses() {
	juese1 := appmodels.Juese{Id: 2, Bianma: "ROLE_USER", Mingcheng: "用户", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 3, Bianma: "ROLE_CANGUAN", Mingcheng: "参观", Biaoji: "Youxiao"}
	jueses := []appmodels.Juese{juese1, juese2}
	log.Println(jueses)
	daos.Tianjia_duoge_juese(jueses)
}
func testjueseservice() {
	juese1 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLIeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Mingcheng: "经理", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLI", Mingcheng: "经理eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Biaoji: "Youxiao"}
	services.Tianjiajuese(&juese1)
	services.Tianjiajuese(&juese2)
}
func main() {
	log.Println(basemodels.Getapppath())
	testshengcheng1()
}
