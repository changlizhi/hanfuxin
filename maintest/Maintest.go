package main

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"hanfuxin/baserun"
	"hanfuxin/baseutils"
	"hanfuxin/juesedaos"
	"hanfuxin/jueseservices"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
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
func testjuesedaotianjia() {
	juese := appmodels.Juese{Id: 1, Bianma: "ROLE_ADMIN", Mingcheng: "管理员", Biaoji: "Youxiao"}
	juesedaos.Tianjia_yige_juese(&juese)
}

func testjuesedaochaxun() {
	juese := juesedaos.Chaxun_yige(4)
	log.Println(juese)
}
func testjuesedaoshanchu() {
	juesedaos.Shanchu_yige(4)
}

func testjueses() {
	juese1 := appmodels.Juese{Id: 2, Bianma: "ROLE_USER", Mingcheng: "用户", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 3, Bianma: "ROLE_CANGUAN", Mingcheng: "参观", Biaoji: "Youxiao"}
	jueses := []appmodels.Juese{juese1, juese2}
	log.Println(jueses)
	juesedaos.Tianjia_duoge_juese(jueses)
}
func testjueseservice() {
	juese1 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLIeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Mingcheng: "经理", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLI", Mingcheng: "经理eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Biaoji: "Youxiao"}
	jueseservices.Tianjiajuese(&juese1)
	jueseservices.Tianjiajuese(&juese2)
}
func testjiami() {
	data := "clz1aaaassssdddd"
	key := "abcd1234"
	jm, err := apputils.Encrypt([]byte(data), []byte(key))

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("jm", convert(jm))
	log.Println("typeof------", reflect.TypeOf(jm))
}
func convert(b []byte) string {
	s := make([]string, len(b))
	for i, val := range b {
		s[i] = string(val)
	}
	return strings.Join(s, ",")
}
func testbasemodel() {
	log.Println(basemodels.Getapppath())
}
func testioutil() {
	os.MkdirAll(basemodels.Getapppath()+"/views/app/", 0777)
	ioutil.WriteFile(basemodels.Getapppath()+"/views/app/sss.go", []byte("ssss"), 0777)
}
func testbaserun() {
	baserun.Shengchengdaos()
}
func main() {
	testbaserun()
}
