package main

import (
	"go/parser"
	"go/token"
	"hanfuxin/appmodels"
	"hanfuxin/apputils"
	"hanfuxin/basemodels"
	"hanfuxin/zdjueseservices"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

func testjueseservice() {
	juese1 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLIeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Mingcheng: "经理", Biaoji: "Youxiao"}
	juese2 := appmodels.Juese{Id: 4, Bianma: "ROLE_JINGLI", Mingcheng: "经理eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", Biaoji: "Youxiao"}
	zdjueseservices.Tianjiajuese(&juese1)
	zdjueseservices.Tianjiajuese(&juese2)
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

func Testf() {
	testfuncname()
}
func Testf2() {
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
	zfss := zf.Zf{}
	log.Println(zfss.Appinits(true))
}
func testparserfile() {
	zfss := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	fset := token.NewFileSet()
	dh := zfzhi.Dianhaozhi()
	xx := zfzhi.Xiexianzhi()
	biaopath := basemodels.Getapppath() + xx + zfss.Zfzhi(true) + xx + zfss.Fuhao(false) + dh + zfss.Go(true)
	log.Println(biaopath)
	f, err := parser.ParseFile(fset, biaopath, nil, parser.AllErrors)

	if err != nil {
		log.Println(err)
		return
	}
	for k, _ := range f.Scope.Objects {
		log.Println(k)
	}
}
func testtypefile() {
	zfss := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	fset := token.NewFileSet()
	dh := zfzhi.Dianhaozhi()
	xx := zfzhi.Xiexianzhi()
	biaopath := basemodels.Getapppath() + xx + zfss.Zfz(true) + xx + zfss.Biao(false) + dh + zfss.Go(true)
	log.Println(biaopath)
	log.Println(zfss.Fangfaming(true))
	b, _ := ioutil.ReadFile(biaopath)
	f, err := parser.ParseFile(fset, "", b, parser.AllErrors)
	if err != nil {
		log.Println("err-----", err)
		return
	}
	log.Println("f------", f)
	for k, _ := range f.Scope.Objects {
		log.Println(k)
	}
}
func testreflect() {
	zf := zf.Zf{}
	t := reflect.TypeOf(&zf)
	log.Println(t.Method(1))
	m, _ := t.MethodByName("Testf")
	log.Println(m.Name)
	v := reflect.ValueOf(&zf)
	log.Println(v.MethodByName("Testf2").Call(nil))
}
func testreg() {
	zfzhi := zfzhi.Zfzhi{}
	zf := zf.Zf{}
	xx := zfzhi.Xiexianzhi()
	dh := zfzhi.Dianhaozhi()

	path := basemodels.Getapppath() + xx + zf.Zfz(true) + xx + zf.Biao(false) + dh + zf.Go(true)
	b, _ := ioutil.ReadFile(path)
	reg, err := regexp.CompilePOSIX(`^func \(zf \*Zf\) [[:word:]]+`)
	if err != nil {
		log.Println(err)
		return
	}
	bfind := reg.FindAllString(string(b), -1)
	log.Println(bfind)
}

func testreg2() {
	regstr := "^Id"
	str := "Idsssss"
	reg, err := regexp.CompilePOSIX(regstr)
	if err != nil {
		log.Println("regerr-------", err)
	}
	log.Println(reg.MatchString(str))
}
func testreg3() {
	regstr := "^[A-Z]"

	str := "sssIdsssss"
	reg, err := regexp.CompilePOSIX(regstr)
	if err != nil {
		log.Println("regerr-------", err)
	}
	log.Println(reg.MatchString(str))
}

func main() {
	testreg3()
}
