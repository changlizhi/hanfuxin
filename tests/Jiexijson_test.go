package tests

import (
	"hanfuxin/appinits"
	"hanfuxin/baseutils"
	"log"
	"testing"
)

func TestChangliangjson(t *testing.T) {
	changliang := baseutils.Changliangjson()
	log.Println("changliang---", changliang)
}
func TestYingyongjson(t *testing.T) {
	yingyong := baseutils.Yingyongjson()
	log.Println("yingyong---", yingyong)
}
func TestJiexi(t *testing.T) {
	path := appinits.Getyingyongpath()
	yingyong := appinits.Yingyong{}
	baseutils.Jiexi(path, &yingyong)
	log.Println(yingyong)
}
