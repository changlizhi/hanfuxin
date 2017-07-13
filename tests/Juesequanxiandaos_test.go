package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdjuesequanxiandaos"
	"log"
	"testing"
)

func TestChaxunyigeJuesequanxian(t *testing.T) {
	juesequanxian := zdjuesequanxiandaos.Chaxunyige(1)
	log.Println(juesequanxian)
}
func TestShanchuyigeJuesequanxian(t *testing.T) {
	zdjuesequanxiandaos.Shanchuyige(1)
}
func TestTianjiayigeJuesequanxian(t *testing.T) {
	juesequanxian := &appmodels.Juesequanxian{
		Biaoji:         "BiaojiTianjiayigeTest1",
		Id:             1,
		Quanxianbianma: "QuanxianbianmaTianjiayigeTest1",
		Juesebianma:    "JuesebianmaTianjiayigeTest1",
	}
	zdjuesequanxiandaos.Tianjiayige(juesequanxian)
}
func TestXiugaiyigeJuesequanxian(t *testing.T) {
	juesequanxian := &appmodels.Juesequanxian{
		Id:             1,
		Quanxianbianma: "QuanxianbianmaXiugaiyigeTest1",
		Juesebianma:    "JuesebianmaXiugaiyigeTest1",
		Biaoji:         "BiaojiXiugaiyigeTest1",
	}
	zdjuesequanxiandaos.Xiugaiyige(juesequanxian)
}
func TestTianjiaduogeJuesequanxian(t *testing.T) {
	juesequanxian2 := appmodels.Juesequanxian{
		Id:             1,
		Quanxianbianma: "QuanxianbianmaTianjiaduogeTest2",
		Juesebianma:    "JuesebianmaTianjiaduogeTest2",
		Biaoji:         "BiaojiTianjiaduogeTest2",
	}
	juesequanxian3 := appmodels.Juesequanxian{
		Juesebianma:    "JuesebianmaTianjiaduogeTest3",
		Biaoji:         "BiaojiTianjiaduogeTest3",
		Quanxianbianma: "QuanxianbianmaTianjiaduogeTest3",
		Id:             1,
	}
	juesequanxians := []appmodels.Juesequanxian{juesequanxian2, juesequanxian3}
	zdjuesequanxiandaos.Tianjiaduoge(juesequanxians)
}
