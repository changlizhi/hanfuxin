package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdquanxiandaos"
	"log"
	"testing"
)

func TestChaxunyigeQuanxian(t *testing.T) {
	quanxian := zdquanxiandaos.Chaxunyige(1)
	log.Println(quanxian)
}
func TestShanchuyigeQuanxian(t *testing.T) {
	zdquanxiandaos.Shanchuyige(1)
}
func TestTianjiayigeQuanxian(t *testing.T) {
	quanxian := &appmodels.Quanxian{
		Id:        1,
		Bianma:    "BianmaTianjiayigeTest1",
		Lujing:    "LujingTianjiayigeTest1",
		Biaoji:    "BiaojiTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
	}
	zdquanxiandaos.Tianjiayige(quanxian)
}
func TestXiugaiyigeQuanxian(t *testing.T) {
	quanxian := &appmodels.Quanxian{
		Bianma:    "BianmaXiugaiyigeTest1",
		Lujing:    "LujingXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Id:        1,
	}
	zdquanxiandaos.Xiugaiyige(quanxian)
}
func TestTianjiaduogeQuanxian(t *testing.T) {
	quanxian2 := appmodels.Quanxian{
		Bianma:    "BianmaTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Id:        1,
		Lujing:    "LujingTianjiaduogeTest2",
	}
	quanxian3 := appmodels.Quanxian{
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
		Lujing:    "LujingTianjiaduogeTest3",
		Mingcheng: "MingchengTianjiaduogeTest3",
		Id:        1,
	}
	quanxians := []appmodels.Quanxian{quanxian2, quanxian3}
	zdquanxiandaos.Tianjiaduoge(quanxians)
}
