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
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
		Lujing:    "LujingTianjiayigeTest1",
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
	}
	zdquanxiandaos.Tianjiayige(quanxian)
}
func TestXiugaiyigeQuanxian(t *testing.T) {
	quanxian := &appmodels.Quanxian{
		Bianma:    "BianmaXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Lujing:    "LujingXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
	}
	zdquanxiandaos.Xiugaiyige(quanxian)
}
func TestTianjiaduoge(t *testing.T) {
	quanxian2 := appmodels.Quanxian{
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest2",
		Lujing:    "LujingTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
	}
	quanxian3 := appmodels.Quanxian{
		Bianma:    "BianmaTianjiaduogeTest3",
		Mingcheng: "MingchengTianjiaduogeTest3",
		Lujing:    "LujingTianjiaduogeTest3",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest3",
	}
	quanxians := []appmodels.Quanxian{quanxian2, quanxian3}
	zdquanxiandaos.Tianjiaduoge(quanxians)
}
