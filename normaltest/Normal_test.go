package normaltest

import (
	"go/parser"
	"go/token"
	"hanfuxin/apputils"
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"log"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func TestFuncname(t *testing.T) {
	pc, file, _, _ := runtime.Caller(1)
	log.Println("pc----", pc)
	log.Println("file-------", file)
	f := runtime.FuncForPC(pc)
	log.Println("f-----------", f)
	log.Println("f.Name()-----------", f.Name())
	log.Println("PERIOD--------", token.PERIOD.String())

	log.Println("funcname------", strings.Split(f.Name(), token.PERIOD.String())[1])
	log.Println("QUO--------", token.QUO.String())
	log.Println("BREAK--------", token.BREAK.String())
	log.Println("illegal-----", token.ILLEGAL.String())
	log.Println("EOF-----", token.EOF.String())
	log.Println("IDENT-----", token.IDENT.String())
	log.Println("COLON-----", token.COLON.String())
}

func TestF(t *testing.T) {
	TestFuncname(t)
}
func TestF2(t *testing.T) {
	TestFuncname(t)
}

func TestPipei3lei(t *testing.T) {
	bu, err := apputils.Pipei3lei("abc", "abcd")
	if err != nil {
		log.Println(err)
	}
	log.Println(bu)
}
func TestZfmulu(t *testing.T) {
	log.Println(zf.Zfs.Appinits(true))
}
func TestParserfile(t *testing.T) {
	fset := token.NewFileSet()
	biaopath := basemodels.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Zfzhi(true) + zfzhi.Zhi.Xx() + zf.Zfs.Fuhao(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
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
func TestTypefile(t *testing.T) {
	fset := token.NewFileSet()
	biaopath := basemodels.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Zfz(true) + zfzhi.Zhi.Xx() + zf.Zfs.Biao(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	log.Println(biaopath)
	log.Println(zf.Zfs.Fangfaming(true))
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
func TestReflect(t *testing.T) {
	zf := zf.Zf{}
	v := reflect.ValueOf(&zf)
	vparam := reflect.ValueOf(&t)
	log.Println("param-------", vparam)
	param := make([]reflect.Value, 1)
	param[0] = vparam
	log.Println(v.MethodByName("Testf2"))
}
func TestReg(t *testing.T) {
	path := basemodels.Getapppath() + zfzhi.Zhi.Xx() + zf.Zfs.Zfz(true) + zfzhi.Zhi.Xx() + zf.Zfs.Biao(false) + zfzhi.Zhi.Dh() + zf.Zfs.Go(true)
	b, _ := ioutil.ReadFile(path)
	reg, err := regexp.CompilePOSIX(`^func \(zf \*Zf\) [[:word:]]+`)
	if err != nil {
		log.Println(err)
		return
	}
	bfind := reg.FindAllString(string(b), -1)
	log.Println(bfind)
}

func TestReg2(t *testing.T) {
	regstr := "^Id"
	str := "Idsssss"
	reg, err := regexp.CompilePOSIX(regstr)
	if err != nil {
		log.Println("regerr-------", err)
	}
	log.Println(reg.MatchString(str))
}
func TestReg3(t *testing.T) {
	regstr := "^[A-Z]"

	str := "sssIdsssss"
	reg, err := regexp.CompilePOSIX(regstr)
	if err != nil {
		log.Println("regerr-------", err)
	}
	log.Println(reg.MatchString(str))
}

type Age int

func TestType(t *testing.T) {
	a := new(Age)
	log.Println("a------", a)
}
