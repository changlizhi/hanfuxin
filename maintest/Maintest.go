package main

import (
	"hanfuxin/appmodels"
	"hanfuxin/jueseservices"
	"hanfuxin/zdjuesedaos"
	"log"
)

func testjuesedaotianjia() {
	juese := appmodels.Juese{Id: 1, Bianma: "ROLE_ADMIN", Mingcheng: "管理员", Biaoji: "Youxiao"}
	zdjuesedaos.Tianjiayige(&juese)
}
func testxiugaijuese() {
	juese := appmodels.Juese{Id: 1, Bianma: "ROLE_ADMIN", Mingcheng: "管理员", Biaoji: "Youxiao"}
	zdjuesedaos.Xiugaiyige(&juese)
}

func testjuesedaochaxun() {
	juese := zdjuesedaos.Chaxunyige(4)
	log.Println(juese)
}
func testjuesedaoshanchu() {
	zdjuesedaos.Shanchuyige(4)
}

func testjueses() {
	juese1 := appmodels.Juese{Id: 2, Bianma: "ROLE_USER", Mingcheng: "用户", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 3, Bianma: "ROLE_CANGUAN", Mingcheng: "参观", Biaoji: "Youxiao"}
	jueses := []appmodels.Juese{juese1, juese2}

	log.Println(jueses)
	zdjuesedaos.Tianjiaduoge(jueses)
}
func testjueseservice() {
	juese1 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLIeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Mingcheng: "经理", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLI", Mingcheng: "经理eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Biaoji: "Youxiao"}
	jueseservices.Tianjiajuese(&juese1)
	jueseservices.Tianjiajuese(&juese2)
}
func main() {
	testxiugaijuese()
}
