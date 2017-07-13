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
		Mingcheng: "MingchengTianjiayigeTest1",
		Bianma:    "BianmaTianjiayigeTest1",
		Leixing:   "LeixingTianjiayigeTest1",
	}
	zdyanzhengdaos.Tianjiayige(yanzheng)
}
func TestXiugaiyigeYanzheng(t *testing.T) {
	yanzheng := &appmodels.Yanzheng{
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Bianma:    "BianmaXiugaiyigeTest1",
		Leixing:   "LeixingXiugaiyigeTest1",
		Id:        1,
		Zhi:       "ZhiXiugaiyigeTest1",
		Mingcheng: "MingchengXiugaiyigeTest1",
	}
	zdyanzhengdaos.Xiugaiyige(yanzheng)
}
func TestTianjiaduogeYanzheng(t *testing.T) {
	yanzheng2 := appmodels.Yanzheng{
		Leixing:   "LeixingTianjiaduogeTest2",
		Id:        1,
		Mingcheng: "MingchengTianjiaduogeTest2",
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Zhi:       "ZhiTianjiaduogeTest2",
		Bianma:    "BianmaTianjiaduogeTest2",
	}
	yanzheng3 := appmodels.Yanzheng{
		Mingcheng: "MingchengTianjiaduogeTest3",
		Id:        1,
		Bianma:    "BianmaTianjiaduogeTest3",
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Leixing:   "LeixingTianjiaduogeTest3",
		Zhi:       "ZhiTianjiaduogeTest3",
	}
	yanzhengs := []appmodels.Yanzheng{yanzheng2, yanzheng3}
	zdyanzhengdaos.Tianjiaduoge(yanzhengs)
}
