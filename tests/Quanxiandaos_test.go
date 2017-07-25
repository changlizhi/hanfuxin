package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdquanxiandaos"
	"log"
	"testing"
)

func TestTianjiaduogeQuanxian(t *testing.T) {
	quanxian2 := appmodels.Quanxian{
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
		Lujing:    "LujingTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
	}
	quanxian3 := appmodels.Quanxian{
		Bianma:    "BianmaTianjiaduogeTest3",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest3",
		Lujing:    "LujingTianjiaduogeTest3",
		Biaoji:    "BiaojiTianjiaduogeTest3",
	}
	quanxians := []appmodels.Quanxian{quanxian2, quanxian3}
	zdquanxiandaos.Tianjiaduoge(quanxians)
}
func TestTianjiayigeQuanxian(t *testing.T) {
	quanxian := &appmodels.Quanxian{
		Mingcheng: "MingchengTianjiayigeTest1",
		Id:        1,
		Lujing:    "LujingTianjiayigeTest1",
		Biaoji:    "BiaojiTianjiayigeTest1",
		Bianma:    "BianmaTianjiayigeTest1",
	}
	zdquanxiandaos.Tianjiayige(quanxian)
}
func TestXiugaiyigeQuanxian(t *testing.T) {
	quanxian := &appmodels.Quanxian{
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Bianma:    "BianmaXiugaiyigeTest1",
		Lujing:    "LujingXiugaiyigeTest1",
	}
	zdquanxiandaos.Xiugaiyige(quanxian)
}
func TestChaxunyigeQuanxian(t *testing.T) {
	quanxian := zdquanxiandaos.Chaxunyige(1)
	log.Println(quanxian)
}
func TestShanchuyigeQuanxian(t *testing.T) {
	zdquanxiandaos.Shanchuyige(1)
}
