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
		Id:        1,
		Bianma:    "BianmaTianjiayigeTest1",
		Biaoji:    "BiaojiTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
	}
	zdjuesedaos.Tianjiayige(juese)
}
func TestXiugaiyigeJuese(t *testing.T) {
	juese := &appmodels.Juese{
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Id:        1,
		Mingcheng: "MingchengXiugaiyigeTest1",
		Bianma:    "BianmaXiugaiyigeTest1",
	}
	zdjuesedaos.Xiugaiyige(juese)
}
func TestTianjiaduogeJuese(t *testing.T) {
	juese2 := appmodels.Juese{
		Mingcheng: "MingchengTianjiaduogeTest2",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
	}
	juese3 := appmodels.Juese{
		Mingcheng: "MingchengTianjiaduogeTest3",
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
		Id:        1,
	}
	jueses := []appmodels.Juese{juese2, juese3}
	zdjuesedaos.Tianjiaduoge(jueses)
}
