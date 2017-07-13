package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdjuesequanxianservices"
	"log"
	"testing"
)

func TestJuesequanxianservicesTianjia(t *testing.T) {
	juesequanxian := appmodels.Juesequanxian{
		Quanxianbianma: "JuesequanxianTianjia",
		Id:             1,
		Juesebianma:    "JuesequanxianTianjia",
		Biaoji:         "JuesequanxianTianjia",
	}
	zdjuesequanxianservices.Tianjiajuesequanxian(&juesequanxian)
}

func TestJuesequanxianservicesXiugai(t *testing.T) {
	juesequanxian := appmodels.Juesequanxian{
		Juesebianma:    "JuesequanxianXiugai",
		Quanxianbianma: "JuesequanxianXiugai",
		Id:             1,
		Biaoji:         "JuesequanxianXiugai",
	}
	zdjuesequanxianservices.Xiugaijuesequanxian(&juesequanxian)
}
func TestJuesequanxianservicesChaxun(t *testing.T) {
	juesequanxian := zdjuesequanxianservices.Chaxunjuesequanxian(1)
	log.Println(juesequanxian)

}
func TestJuesequanxianservicesShanchu(t *testing.T) {
	zdjuesequanxianservices.Shanchujuesequanxian(1)
}
