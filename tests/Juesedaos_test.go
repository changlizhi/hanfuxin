package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdjuesedaos"
	"log"
	"testing"
)

func TestTianjiaduogeJuese(t *testing.T) {
	juese2 := appmodels.Juese{
		Bianma:    "BianmaTianjiaduogeTest2",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest2",
		Biaoji:    "BiaojiTianjiaduogeTest2",
	}
	juese3 := appmodels.Juese{
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
		Mingcheng: "MingchengTianjiaduogeTest3",
	}
	jueses := []appmodels.Juese{juese2, juese3}
	zdjuesedaos.Tianjiaduoge(jueses)
}
func TestTianjiayigeJuese(t *testing.T) {
	juese := &appmodels.Juese{
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
	}
	zdjuesedaos.Tianjiayige(juese)
}
func TestXiugaiyigeJuese(t *testing.T) {
	juese := &appmodels.Juese{
		Bianma:    "BianmaXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
	}
	zdjuesedaos.Xiugaiyige(juese)
}
func TestChaxunyigeJuese(t *testing.T) {
	juese := zdjuesedaos.Chaxunyige(1)
	log.Println(juese)
}
func TestShanchuyigeJuese(t *testing.T) {
	zdjuesedaos.Shanchuyige(1)
}
