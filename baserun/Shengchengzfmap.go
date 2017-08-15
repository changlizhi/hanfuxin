package baserun

import (
	"bytes"
	"hanfuxin/jiexi"
	"hanfuxin/sjk"
	"changliang/zf"
	"changliang/zfzhi"
	"log"
	"reflect"
	"regexp"
	"strings"
)

func huoqujiegou(wenjianming string) map[string]string {
	ret := make(map[string]string)
	ffs := jiexi.Pipeifangfa(zf.Zfs.Sjk(true), zf.Zfs.Sjk(false), zf.Zfs.Sjk(true), wenjianming)
	for _, ff := range ffs {
		ret[ff] = strings.ToLower(ff)
	}
	return ret
}
func Huoqubiaos() map[string]string {
	ret := huoqujiegou(zf.Zfs.Biao(false))
	return ret
}
func Huoqulies() map[string]string {
	ret := huoqujiegou(zf.Zfs.Lie(false))
	return ret
}
func Huoqubiaolies() map[string]string {
	ret := huoqujiegou(zf.Zfs.Biaolie(false))
	return ret
}
func Huoquliechangdus() map[string]string {
	ret := huoqujiegou(zf.Zfs.Liechangdu(false))
	return ret
}
func Huoqulieleixings() map[string]string {
	ret := huoqujiegou(zf.Zfs.Lieleixing(false))
	return ret
}
func Huoquyigechangdu(lieming string) int {
	sjk := sjk.Sjk{}
	ffm := lieming + zf.Zfs.Changduzhi(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[0].Interface().(int)
	return ret
}

func Huoquyigeleixing(lieming string) string {
	sjk := sjk.Sjk{}
	ffm := lieming + zf.Zfs.Leixingzhi(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[0].Interface().(string)
	if ret == zfzhi.Zhi.Kzf() {
		log.Println("类型错误-----空")
		return zfzhi.Zhi.Kzf()
	}
	return ret
}

func Huoquyigebiaojiegou(biaoming string) map[string]string {
	regwordstr := zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Mh() + zf.Zfs.Word(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Jia()
	regbiaostr := zfzhi.Zhi.Cfh() + biaoming + regwordstr
	regbiao, _ := regexp.CompilePOSIX(regbiaostr)
	regtoustr := zfzhi.Zhi.Cfh() + biaoming
	regtou, _ := regexp.CompilePOSIX(regtoustr)

	regdaxiestr := zfzhi.Zhi.Cfh() + zfzhi.Zhi.Zkhz() + zf.Zfs.A(false) + zfzhi.Zhi.Jian() + zf.Zfs.Z(false) + zfzhi.Zhi.Zkhy()
	regdaxie, _ := regexp.CompilePOSIX(regdaxiestr)
	ret := make(map[string]string)
	for biaolie, _ := range Huoqubiaolies() {
		if regbiao.MatchString(biaolie) {
			biaotou := regbiao.FindString(biaolie)
			mixinlie := regtou.ReplaceAllString(biaotou, zfzhi.Zhi.Kzf())
			if regdaxie.MatchString(mixinlie) {
				ret[mixinlie] = strings.ToLower(mixinlie)
			}
		}
	}
	return ret
}

func Buffertoumodel(buffer *bytes.Buffer, bianma string, baoming string) {
	// package xxx \n
	pac := zf.Zfs.Package(true) + zfzhi.Zhi.Kgf() + baoming + zfzhi.Zhi.Hhf()
	buffer.WriteString(pac)

	// type Juese struct{\n
	typestr := zf.Zfs.Type(true) + zfzhi.Zhi.Kgf() + bianma + zfzhi.Zhi.Kgf() + zf.Zfs.Struct(true) + zfzhi.Zhi.Dkhz() + zfzhi.Zhi.Hhf()
	buffer.WriteString(typestr)
}
