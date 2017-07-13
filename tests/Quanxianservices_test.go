package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdquanxianservices"
	"log"
	"testing"
)

func TestQuanxianservicesTianjia(t *testing.T) {
	quanxian := appmodels.Quanxian{
		Lujing:    "QuanxianTianjia",
		Mingcheng: "QuanxianTianjia",
		Biaoji:    "QuanxianTianjia",
		Bianma:    "QuanxianTianjia",
		Id:        1,
	}
	zdquanxianservices.Tianjiaquanxian(&quanxian)
}

func TestQuanxianservicesXiugai(t *testing.T) {
	quanxian := appmodels.Quanxian{
		Mingcheng: "QuanxianXiugai",
		Lujing:    "QuanxianXiugai",
		Id:        1,
		Biaoji:    "QuanxianXiugai",
		Bianma:    "QuanxianXiugai",
	}
	zdquanxianservices.Xiugaiquanxian(&quanxian)
}
func TestQuanxianservicesChaxun(t *testing.T) {
	quanxian := zdquanxianservices.Chaxunquanxian(1)
	log.Println(quanxian)

}
func TestQuanxianservicesShanchu(t *testing.T) {
	zdquanxianservices.Shanchuquanxian(1)
}
