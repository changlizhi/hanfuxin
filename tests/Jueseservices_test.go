package tests

import (
	"hanfuxin/appmodels"
	"hanfuxin/zdjueseservices"
	"log"
	"testing"
)

func TestJueseservicesTianjia(t *testing.T) {
	juese := appmodels.Juese{
		Mingcheng: "JueseTianjia",
		Bianma:    "JueseTianjia",
		Id:        1,
		Biaoji:    "JueseTianjia",
	}
	zdjueseservices.Tianjiajuese(&juese)
}

func TestJueseservicesXiugai(t *testing.T) {
	juese := appmodels.Juese{
		Mingcheng: "JueseXiugai",
		Id:        1,
		Bianma:    "JueseXiugai",
		Biaoji:    "JueseXiugai",
	}
	zdjueseservices.Xiugaijuese(&juese)
}
func TestJueseservicesChaxun(t *testing.T) {
	juese := zdjueseservices.Chaxunjuese(1)
	log.Println(juese)

}
func TestJueseservicesShanchu(t *testing.T) {
	zdjueseservices.Shanchujuese(1)
}
