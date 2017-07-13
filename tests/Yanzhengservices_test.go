package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdyanzhengservices"
	"log"
	"testing"
)

func TestYanzhengservicesTianjia(t *testing.T) {
	yanzheng := appmodels.Yanzheng{
		Biaoji:    "YanzhengTianjia",
		Bianma:    "YanzhengTianjia",
		Leixing:   "YanzhengTianjia",
		Zhi:       "YanzhengTianjia",
		Mingcheng: "YanzhengTianjia",
		Id:        1,
	}
	zdyanzhengservices.Tianjiayanzheng(&yanzheng)
}

func TestYanzhengservicesXiugai(t *testing.T) {
	yanzheng := appmodels.Yanzheng{
		Biaoji:    "YanzhengXiugai",
		Mingcheng: "YanzhengXiugai",
		Bianma:    "YanzhengXiugai",
		Leixing:   "YanzhengXiugai",
		Zhi:       "YanzhengXiugai",
		Id:        1,
	}
	zdyanzhengservices.Xiugaiyanzheng(&yanzheng)
}
func TestYanzhengservicesChaxun(t *testing.T) {
	yanzheng := zdyanzhengservices.Chaxunyanzheng(1)
	log.Println(yanzheng)

}
func TestYanzhengservicesShanchu(t *testing.T) {
	zdyanzhengservices.Shanchuyanzheng(1)
}
