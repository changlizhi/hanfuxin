package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdxinxijuesedaos"
	"log"
	"testing"
)

func TestTianjiaduogeXinxijuese(t *testing.T) {
	xinxijuese2 := appmodels.Xinxijuese{
		Id:          1,
		Biaoji:      "BiaojiTianjiaduogeTest2",
		Xinxibianma: "XinxibianmaTianjiaduogeTest2",
		Juesebianma: "JuesebianmaTianjiaduogeTest2",
	}
	xinxijuese3 := appmodels.Xinxijuese{
		Biaoji:      "BiaojiTianjiaduogeTest3",
		Id:          1,
		Juesebianma: "JuesebianmaTianjiaduogeTest3",
		Xinxibianma: "XinxibianmaTianjiaduogeTest3",
	}
	xinxijueses := []appmodels.Xinxijuese{xinxijuese2, xinxijuese3}
	zdxinxijuesedaos.Tianjiaduoge(xinxijueses)
}
func TestTianjiayigeXinxijuese(t *testing.T) {
	xinxijuese := &appmodels.Xinxijuese{
		Biaoji:      "BiaojiTianjiayigeTest1",
		Id:          1,
		Juesebianma: "JuesebianmaTianjiayigeTest1",
		Xinxibianma: "XinxibianmaTianjiayigeTest1",
	}
	zdxinxijuesedaos.Tianjiayige(xinxijuese)
}
func TestXiugaiyigeXinxijuese(t *testing.T) {
	xinxijuese := &appmodels.Xinxijuese{
		Id:          1,
		Xinxibianma: "XinxibianmaXiugaiyigeTest1",
		Biaoji:      "BiaojiXiugaiyigeTest1",
		Juesebianma: "JuesebianmaXiugaiyigeTest1",
	}
	zdxinxijuesedaos.Xiugaiyige(xinxijuese)
}
func TestChaxunyigeXinxijuese(t *testing.T) {
	xinxijuese := zdxinxijuesedaos.Chaxunyige(1)
	log.Println(xinxijuese)
}
func TestShanchuyigeXinxijuese(t *testing.T) {
	zdxinxijuesedaos.Shanchuyige(1)
}
