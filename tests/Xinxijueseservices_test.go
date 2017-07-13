package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdxinxijueseservices"
	"log"
	"testing"
)

func TestXinxijueseservicesTianjia(t *testing.T) {
	xinxijuese := appmodels.Xinxijuese{
		Juesebianma: "XinxijueseTianjia",
		Xinxibianma: "XinxijueseTianjia",
		Id:          1,
		Biaoji:      "XinxijueseTianjia",
	}
	zdxinxijueseservices.Tianjiaxinxijuese(&xinxijuese)
}

func TestXinxijueseservicesXiugai(t *testing.T) {
	xinxijuese := appmodels.Xinxijuese{
		Id:          1,
		Xinxibianma: "XinxijueseXiugai",
		Biaoji:      "XinxijueseXiugai",
		Juesebianma: "XinxijueseXiugai",
	}
	zdxinxijueseservices.Xiugaixinxijuese(&xinxijuese)
}
func TestXinxijueseservicesChaxun(t *testing.T) {
	xinxijuese := zdxinxijueseservices.Chaxunxinxijuese(1)
	log.Println(xinxijuese)

}
func TestXinxijueseservicesShanchu(t *testing.T) {
	zdxinxijueseservices.Shanchuxinxijuese(1)
}
