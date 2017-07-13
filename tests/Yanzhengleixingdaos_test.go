package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdyanzhengleixingdaos"
	"log"
	"testing"
)

func TestChaxunyigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := zdyanzhengleixingdaos.Chaxunyige(1)
	log.Println(yanzhengleixing)
}
func TestShanchuyigeYanzhengleixing(t *testing.T) {
	zdyanzhengleixingdaos.Shanchuyige(1)
}
func TestTianjiayigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := &appmodels.Yanzhengleixing{
		Mingcheng: "MingchengTianjiayigeTest1",
		Biaoji:    "BiaojiTianjiayigeTest1",
		Id:        1,
		Bianma:    "BianmaTianjiayigeTest1",
	}
	zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)
}
func TestXiugaiyigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := &appmodels.Yanzhengleixing{
		Id:        1,
		Mingcheng: "MingchengXiugaiyigeTest1",
		Bianma:    "BianmaXiugaiyigeTest1",
		Biaoji:    "BiaojiXiugaiyigeTest1",
	}
	zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixing)
}
func TestTianjiaduogeYanzhengleixing(t *testing.T) {
	yanzhengleixing2 := appmodels.Yanzhengleixing{
		Bianma:    "BianmaTianjiaduogeTest2",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
	}
	yanzhengleixing3 := appmodels.Yanzhengleixing{
		Mingcheng: "MingchengTianjiaduogeTest3",
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
		Id:        1,
	}
	yanzhengleixings := []appmodels.Yanzhengleixing{yanzhengleixing2, yanzhengleixing3}
	zdyanzhengleixingdaos.Tianjiaduoge(yanzhengleixings)
}
