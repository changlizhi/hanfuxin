package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdxinxidaos"
	"log"
	"testing"
)

func TestTianjiaduogeXinxi(t *testing.T) {
	xinxi2 := appmodels.Xinxi{
		Youxiang:  "YouxiangTianjiaduogeTest2",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest2",
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
	}
	xinxi3 := appmodels.Xinxi{
		Mingcheng: "MingchengTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
		Youxiang:  "YouxiangTianjiaduogeTest3",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest3",
	}
	xinxis := []appmodels.Xinxi{xinxi2, xinxi3}
	zdxinxidaos.Tianjiaduoge(xinxis)
}
func TestTianjiayigeXinxi(t *testing.T) {
	xinxi := &appmodels.Xinxi{
		Mingcheng: "MingchengTianjiayigeTest1",
		Id:        1,
		Bianma:    "BianmaTianjiayigeTest1",
		Biaoji:    "BiaojiTianjiayigeTest1",
		Youxiang:  "YouxiangTianjiayigeTest1",
	}
	zdxinxidaos.Tianjiayige(xinxi)
}
func TestXiugaiyigeXinxi(t *testing.T) {
	xinxi := &appmodels.Xinxi{
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Id:        1,
		Bianma:    "BianmaXiugaiyigeTest1",
		Youxiang:  "YouxiangXiugaiyigeTest1",
	}
	zdxinxidaos.Xiugaiyige(xinxi)
}
func TestChaxunyigeXinxi(t *testing.T) {
	xinxi := zdxinxidaos.Chaxunyige(1)
	log.Println(xinxi)
}
func TestShanchuyigeXinxi(t *testing.T) {
	zdxinxidaos.Shanchuyige(1)
}
