package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdyanzhengdaos"
	"log"
	"testing"
)

func TestTianjiaduogeYanzheng(t *testing.T) {
	yanzheng2 := appmodels.Yanzheng{
		Leixing:   "LeixingTianjiaduogeTest2",
		Id:        1,
		Bianma:    "BianmaTianjiaduogeTest2",
		Zhi:       "ZhiTianjiaduogeTest2",
		Biaoji:    "BiaojiTianjiaduogeTest2",
		Mingcheng: "MingchengTianjiaduogeTest2",
	}
	yanzheng3 := appmodels.Yanzheng{
		Mingcheng: "MingchengTianjiaduogeTest3",
		Bianma:    "BianmaTianjiaduogeTest3",
		Leixing:   "LeixingTianjiaduogeTest3",
		Biaoji:    "BiaojiTianjiaduogeTest3",
		Zhi:       "ZhiTianjiaduogeTest3",
		Id:        1,
	}
	yanzhengs := []appmodels.Yanzheng{yanzheng2, yanzheng3}
	zdyanzhengdaos.Tianjiaduoge(yanzhengs)
}
func TestTianjiayigeYanzheng(t *testing.T) {
	yanzheng := &appmodels.Yanzheng{
		Id:        1,
		Biaoji:    "BiaojiTianjiayigeTest1",
		Zhi:       "ZhiTianjiayigeTest1",
		Leixing:   "LeixingTianjiayigeTest1",
		Bianma:    "BianmaTianjiayigeTest1",
		Mingcheng: "MingchengTianjiayigeTest1",
	}
	zdyanzhengdaos.Tianjiayige(yanzheng)
}
func TestXiugaiyigeYanzheng(t *testing.T) {
	yanzheng := &appmodels.Yanzheng{
		Mingcheng: "MingchengXiugaiyigeTest1",
		Leixing:   "LeixingXiugaiyigeTest1",
		Id:        1,
		Biaoji:    "BiaojiXiugaiyigeTest1",
		Zhi:       "ZhiXiugaiyigeTest1",
		Bianma:    "BianmaXiugaiyigeTest1",
	}
	zdyanzhengdaos.Xiugaiyige(yanzheng)
}
func TestChaxunyigeYanzheng(t *testing.T) {
	yanzheng := zdyanzhengdaos.Chaxunyige(1)
	log.Println(yanzheng)
}
func TestShanchuyigeYanzheng(t *testing.T) {
	zdyanzhengdaos.Shanchuyige(1)
}
