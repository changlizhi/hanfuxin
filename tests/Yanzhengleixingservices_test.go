package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdyanzhengleixingservices"
	"log"
	"testing"
)

func TestYanzhengleixingservicesTianjia(t *testing.T) {
	yanzhengleixing := appmodels.Yanzhengleixing{
		Mingcheng: "YanzhengleixingTianjia",
		Bianma:    "YanzhengleixingTianjia",
		Id:        1,
		Biaoji:    "YanzhengleixingTianjia",
	}
	zdyanzhengleixingservices.Tianjiayanzhengleixing(&yanzhengleixing)
}

func TestYanzhengleixingservicesXiugai(t *testing.T) {
	yanzhengleixing := appmodels.Yanzhengleixing{
		Id:        1,
		Mingcheng: "YanzhengleixingXiugai",
		Bianma:    "YanzhengleixingXiugai",
		Biaoji:    "YanzhengleixingXiugai",
	}
	zdyanzhengleixingservices.Xiugaiyanzhengleixing(&yanzhengleixing)
}
func TestYanzhengleixingservicesChaxun(t *testing.T) {
	yanzhengleixing := zdyanzhengleixingservices.Chaxunyanzhengleixing(1)
	log.Println(yanzhengleixing)

}
func TestYanzhengleixingservicesShanchu(t *testing.T) {
	zdyanzhengleixingservices.Shanchuyanzhengleixing(1)
}
