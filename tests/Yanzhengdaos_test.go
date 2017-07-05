package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdyanzhengdaos"
	"log"
	"testing"
)

func TestChaxunyigeYanzheng(t *testing.T) {
	yanzheng := zdyanzhengdaos.Chaxunyige(1)
	log.Println(yanzheng)
}
func TestShanchuyigeYanzheng(t *testing.T) {
	zdyanzhengdaos.Shanchuyige(1)
}
func TestTianjiayigeYanzheng(t *testing.T) {
	yanzheng := &appmodels.Yanzheng{
		Zhi:       "ZhiTianjiayigeTest1",
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
		Leixing:   "LeixingTianjiayigeTest1",
	}
	zdyanzhengdaos.Tianjiayige(yanzheng)
}
func TestXiugaiyigeYanzheng(t *testing.T) {
	yanzheng := &appmodels.Yanzheng{
		Bianma:    "BianmaXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
		Leixing:   "LeixingXiugaiyigeTest1",
		Zhi:       "ZhiXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
	}
	zdyanzhengdaos.Xiugaiyige(yanzheng)
}
func TestTianjiaduoge(t *testing.T) {
	yanzheng2 := appmodels.Yanzheng{
		Bianma:    "BianmaTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
		Leixing:   "LeixingTianjiaduogeTest2",
		Zhi:       "ZhiTianjiaduogeTest2",
		Id:        1,
		Biaoji:    "BiaojiTianjiaduogeTest2",
	}
	yanzheng3 := appmodels.Yanzheng{
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Zhi:       "ZhiTianjiaduogeTest3",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest3",
		Leixing:   "LeixingTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
	}
	yanzhengs := []appmodels.Yanzheng{yanzheng2, yanzheng3}
	zdyanzhengdaos.Tianjiaduoge(yanzhengs)
}
