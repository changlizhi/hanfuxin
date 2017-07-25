package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdyanzhengleixingdaos"
	"log"
	"testing"
)

func TestTianjiaduogeYanzhengleixing(t *testing.T) {
	yanzhengleixing2 := appmodels.Yanzhengleixing{
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
		Id:        1,
	}
	yanzhengleixing3 := appmodels.Yanzhengleixing{
		Bianma:    "BianmaTianjiaduogeTest3",
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest3",
	}
	yanzhengleixings := []appmodels.Yanzhengleixing{yanzhengleixing2, yanzhengleixing3}
	zdyanzhengleixingdaos.Tianjiaduoge(yanzhengleixings)
}
func TestTianjiayigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := &appmodels.Yanzhengleixing{
		Mingcheng: "MingchengTianjiayigeTest1",
		Bianma:    "BianmaTianjiayigeTest1",
		Biaoji:    "BiaojiTianjiayigeTest1",
		Id:        1,
	}
	zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)
}
func TestXiugaiyigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := &appmodels.Yanzhengleixing{
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Id:        1,
		Bianma:    "BianmaXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
	}
	zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixing)
}
func TestChaxunyigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := zdyanzhengleixingdaos.Chaxunyige(1)
	log.Println(yanzhengleixing)
}
func TestShanchuyigeYanzhengleixing(t *testing.T) {
	zdyanzhengleixingdaos.Shanchuyige(1)
}
