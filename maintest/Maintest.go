package main

import (
	"go/parser"
	"go/token"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/basemodels"
	"hanfuxin/zdjueseservices"
	"hanfuxin/zfmulu"
	"log"
	"reflect"
	"runtime"
	"strings"
)

func testjueseservice() {
	juese1 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLIeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Mingcheng: "经理", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLI", Mingcheng: "经理eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Biaoji: "Youxiao"}
	zdjueseservices.Tianjiajuese(&juese1)
	zdjueseservices.Tianjiajuese(&juese2)
}

type zf struct {
}

func testfuncname() {
	pc, file, _, _ := runtime.Caller(1)
	log.Println("pc----", pc)
	log.Println("file-------", file)
	f := runtime.FuncForPC(pc)
	log.Println("funcname------", strings.Split(f.Name(), token.PERIOD.String())[2])
	log.Println("PERIOD--------", token.PERIOD.String())
	log.Println("QUO--------", token.QUO.String())
	log.Println("BREAK--------", token.BREAK.String())
	log.Println("illegal-----", token.ILLEGAL.String())
	log.Println("EOF-----", token.EOF.String())
	log.Println("IDENT-----", token.IDENT.String())
	log.Println("COLON-----", token.COLON.String())
}

func (zf *zf) Testf() {
	testfuncname()
}
func (zf *zf) Testf2() {
	testfuncname()
}

func testpipei3lei() {
	bu, err := apputils.Pipei3lei("abc", "abcd")
	if err != nil {
		log.Println(err)
	}
	log.Println(bu)
}
func testzfmulu() {
	log.Println(zfmulu.Appinits(true))
}
func testparserfile() {
	fset := token.NewFileSet()
	zimu := basemodels.Getapppath() + "/zfzimu/Zimu.go"
	f, err := parser.ParseFile(fset, zimu, nil, parser.AllErrors)

	if err != nil {
		log.Println(err)
		return
	}
	for k, _ := range f.Scope.Objects {
		log.Println(k)
	}
}
func testreflect() {
	zf := zf{}
	t := reflect.TypeOf(&zf)
	log.Println(t.Method(1))
	m, _ := t.MethodByName("Testf")
	log.Println(m.Name)
	v := reflect.ValueOf(&zf)
	log.Println(v.MethodByName("Testf2").Call(nil))
}
func main() {
	testreflect()
}
