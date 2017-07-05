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
		Id:             1,
		Biaoji:         "BiaojiTianjiayigeTest1",
		Juesebianma:    "JuesebianmaTianjiayigeTest1",
		Quanxianbianma: "QuanxianbianmaTianjiayigeTest1",
	}
	zdjuesequanxiandaos.Tianjiayige(juesequanxian)
}
func TestXiugaiyigeJuesequanxian(t *testing.T) {
	juesequanxian := &appmodels.Juesequanxian{
		Id:             1,
		Biaoji:         "BiaojiXiugaiyigeTest1",
		Juesebianma:    "JuesebianmaXiugaiyigeTest1",
		Quanxianbianma: "QuanxianbianmaXiugaiyigeTest1",
	}
	zdjuesequanxiandaos.Xiugaiyige(juesequanxian)
}
func TestTianjiaduoge(t *testing.T) {
	juesequanxian2 := appmodels.Juesequanxian{
		Id:             1,
		Biaoji:         "BiaojiTianjiaduogeTest2",
		Juesebianma:    "JuesebianmaTianjiaduogeTest2",
		Quanxianbianma: "QuanxianbianmaTianjiaduogeTest2",
	}
	juesequanxian3 := appmodels.Juesequanxian{
		Id:             1,
		Biaoji:         "BiaojiTianjiaduogeTest3",
		Juesebianma:    "JuesebianmaTianjiaduogeTest3",
		Quanxianbianma: "QuanxianbianmaTianjiaduogeTest3",
	}
	juesequanxians := []appmodels.Juesequanxian{juesequanxian2, juesequanxian3}
	zdjuesequanxiandaos.Tianjiaduoge(juesequanxians)
}
