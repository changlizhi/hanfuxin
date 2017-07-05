package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdjuesedaos"
	"log"
	"testing"
)

func TestChaxunyigeJuese(t *testing.T) {
	juese := zdjuesedaos.Chaxunyige(1)
	log.Println(juese)
}
func TestShanchuyigeJuese(t *testing.T) {
	zdjuesedaos.Shanchuyige(1)
}
func TestTianjiayigeJuese(t *testing.T) {
	juese := &appmodels.Juese{
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
	}
	zdjuesedaos.Tianjiayige(juese)
}
func TestXiugaiyigeJuese(t *testing.T) {
	juese := &appmodels.Juese{
		Bianma:    "BianmaXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
	}
	zdjuesedaos.Xiugaiyige(juese)
}
func TestTianjiaduoge(t *testing.T) {
	juese2 := appmodels.Juese{
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
	}
	juese3 := appmodels.Juese{
		Bianma:    "BianmaTianjiaduogeTest3",
		Mingcheng: "MingchengTianjiaduogeTest3",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest3",
	}
	jueses := []appmodels.Juese{juese2, juese3}
	zdjuesedaos.Tianjiaduoge(jueses)
}
