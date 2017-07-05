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
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
	}
	zdyanzhengleixingdaos.Tianjiayige(yanzhengleixing)
}
func TestXiugaiyigeYanzhengleixing(t *testing.T) {
	yanzhengleixing := &appmodels.Yanzhengleixing{
		Bianma:    "BianmaXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
	}
	zdyanzhengleixingdaos.Xiugaiyige(yanzhengleixing)
}
func TestTianjiaduoge(t *testing.T) {
	yanzhengleixing2 := appmodels.Yanzhengleixing{
		Bianma:    "BianmaTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
	}
	yanzhengleixing3 := appmodels.Yanzhengleixing{
		Bianma:    "BianmaTianjiaduogeTest3",
		Mingcheng: "MingchengTianjiaduogeTest3",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest3",
	}
	yanzhengleixings := []appmodels.Yanzhengleixing{yanzhengleixing2, yanzhengleixing3}
	zdyanzhengleixingdaos.Tianjiaduoge(yanzhengleixings)
}
