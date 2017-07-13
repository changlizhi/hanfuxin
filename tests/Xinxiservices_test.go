package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdxinxiservices"
	"log"
	"testing"
)

func TestXinxiservicesTianjia(t *testing.T) {
	xinxi := appmodels.Xinxi{
		Bianma:    "XinxiTianjia",
		Mingcheng: "XinxiTianjia",
		Id:        1,
		Biaoji:    "XinxiTianjia",
		Youxiang:  "XinxiTianjia",
	}
	zdxinxiservices.Tianjiaxinxi(&xinxi)
}

func TestXinxiservicesXiugai(t *testing.T) {
	xinxi := appmodels.Xinxi{
		Youxiang:  "XinxiXiugai",
		Id:        1,
		Biaoji:    "XinxiXiugai",
		Bianma:    "XinxiXiugai",
		Mingcheng: "XinxiXiugai",
	}
	zdxinxiservices.Xiugaixinxi(&xinxi)
}
func TestXinxiservicesChaxun(t *testing.T) {
	xinxi := zdxinxiservices.Chaxunxinxi(1)
	log.Println(xinxi)

}
func TestXinxiservicesShanchu(t *testing.T) {
	zdxinxiservices.Shanchuxinxi(1)
}
