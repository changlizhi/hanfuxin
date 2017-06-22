package main

import (
	"hanfuxin/appinits"
	"hanfuxin/apputils"
	"hanfuxin/baserun"
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
	obj := apputils.Yingyongjson()
	log.Println(obj, "\n")
}
func testchangliang() {
	obj := apputils.Changliangjson()
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
}
func main() {
	testshengcheng2()
}
