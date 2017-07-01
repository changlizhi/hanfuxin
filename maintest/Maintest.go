package main

import (
	"hanfuxin/appinits"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/baseinits"
	"hanfuxin/basemodels"
	"hanfuxin/baserun"
	"hanfuxin/baseutils"
	"hanfuxin/juesedaos"
	"hanfuxin/jueseservices"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func testshengcheng1() {
	baserun.Shengcheng_yingyongzi_model()
}
func testshengcheng2() {
	baserun.Shengcheng_yingyong_model()
}
func testjuesedaotianjia() {
	juese := appmodels.Juese{Id: 1, Bianma: "ROLE_ADMIN", Mingcheng: "管理员", Biaoji: "Youxiao"}
	juesedaos.Tianjiayige(&juese)
}

func testjuesedaochaxun() {
	juese := juesedaos.Chaxunyige(4)
	log.Println(juese)
}
func testjuesedaoshanchu() {
	juesedaos.Shanchuyige(4)
}

func testjueses() {
	juese1 := appmodels.Juese{Id: 2, Bianma: "ROLE_USER", Mingcheng: "用户", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 3, Bianma: "ROLE_CANGUAN", Mingcheng: "参观", Biaoji: "Youxiao"}
	jueses := []appmodels.Juese{juese1, juese2}
	log.Println(jueses)
	juesedaos.Tianjiaduoge(jueses)
}
func testjueseservice() {
	juese1 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLIeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Mingcheng: "经理", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLI", Mingcheng: "经理eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Biaoji: "Youxiao"}
	jueseservices.Tianjiajuese(&juese1)
	jueseservices.Tianjiajuese(&juese2)
}
func testbaserun() {
	baserun.Shengchengdaos()
}
func main() {
	testjuesedaotianjia()
}
