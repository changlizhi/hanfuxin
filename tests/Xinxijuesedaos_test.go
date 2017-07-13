package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdxinxijuesedaos"
	"log"
	"testing"
)

func TestChaxunyigeXinxijuese(t *testing.T) {
	xinxijuese := zdxinxijuesedaos.Chaxunyige(1)
	log.Println(xinxijuese)
}
func TestShanchuyigeXinxijuese(t *testing.T) {
	zdxinxijuesedaos.Shanchuyige(1)
}
func TestTianjiayigeXinxijuese(t *testing.T) {
	xinxijuese := &appmodels.Xinxijuese{
		Xinxibianma: "XinxibianmaTianjiayigeTest1",
		Juesebianma: "JuesebianmaTianjiayigeTest1",
		Id:          1,
		Biaoji:      "BiaojiTianjiayigeTest1",
	}
	zdxinxijuesedaos.Tianjiayige(xinxijuese)
}
func TestXiugaiyigeXinxijuese(t *testing.T) {
	xinxijuese := &appmodels.Xinxijuese{
		Xinxibianma: "XinxibianmaXiugaiyigeTest1",
		Juesebianma: "JuesebianmaXiugaiyigeTest1",
		Id:          1,
		Biaoji:      "BiaojiXiugaiyigeTest1",
	}
	zdxinxijuesedaos.Xiugaiyige(xinxijuese)
}
func TestTianjiaduogeXinxijuese(t *testing.T) {
	xinxijuese2 := appmodels.Xinxijuese{
		Juesebianma: "JuesebianmaTianjiaduogeTest2",
		Xinxibianma: "XinxibianmaTianjiaduogeTest2",
		Biaoji:      "BiaojiTianjiaduogeTest2",
		Id:          1,
	}
	xinxijuese3 := appmodels.Xinxijuese{
		Xinxibianma: "XinxibianmaTianjiaduogeTest3",
		Id:          1,
		Biaoji:      "BiaojiTianjiaduogeTest3",
		Juesebianma: "JuesebianmaTianjiaduogeTest3",
	}
	xinxijueses := []appmodels.Xinxijuese{xinxijuese2, xinxijuese3}
	zdxinxijuesedaos.Tianjiaduoge(xinxijueses)
}
