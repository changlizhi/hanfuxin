package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdxinxidaos"
	"log"
	"testing"
)

func TestChaxunyigeXinxi(t *testing.T) {
	xinxi := zdxinxidaos.Chaxunyige(1)
	log.Println(xinxi)
}
func TestShanchuyigeXinxi(t *testing.T) {
	zdxinxidaos.Shanchuyige(1)
}
func TestTianjiayigeXinxi(t *testing.T) {
	xinxi := &appmodels.Xinxi{
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
		Youxiang:  "YouxiangTianjiayigeTest1",
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
	}
	zdxinxidaos.Tianjiayige(xinxi)
}
func TestXiugaiyigeXinxi(t *testing.T) {
	xinxi := &appmodels.Xinxi{
		Bianma:    "BianmaXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Youxiang:  "YouxiangXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
	}
	zdxinxidaos.Xiugaiyige(xinxi)
}
func TestTianjiaduoge(t *testing.T) {
	xinxi2 := appmodels.Xinxi{
		Bianma:    "BianmaTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
		Youxiang:  "YouxiangTianjiaduogeTest2",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
	}
	xinxi3 := appmodels.Xinxi{
		Bianma:    "BianmaTianjiaduogeTest3",
		Mingcheng: "MingchengTianjiaduogeTest3",
		Youxiang:  "YouxiangTianjiaduogeTest3",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest3",
	}
	xinxis := []appmodels.Xinxi{xinxi2, xinxi3}
	zdxinxidaos.Tianjiaduoge(xinxis)
}
