package baserun

import (
	"bytes"
	"hanfuxin/jiexi"
	"hanfuxin/sjk"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"log"
	"reflect"
	"regexp"
	"strings"
)

func huoqujiegou(wenjianming string) map[string]string {
	zf := zf.Zf{}
	ret := make(map[string]string)
	ffs := jiexi.Pipeifangfa(zf.Sjk(true), zf.Sjk(false), zf.Sjk(true), wenjianming)
	for _, ff := range ffs {
		ret[ff] = strings.ToLower(ff)
	}
	return ret
}
func Huoqubiaos() map[string]string {
	zf := zf.Zf{}
	ret := huoqujiegou(zf.Biao(false))
	return ret
}
func Huoqulies() map[string]string {
	zf := zf.Zf{}
	ret := huoqujiegou(zf.Lie(false))
	return ret
}
func Huoqubiaolies() map[string]string {
	zf := zf.Zf{}
	ret := huoqujiegou(zf.Biaolie(false))
	return ret
}
func Huoquliechangdus() map[string]string {
	zf := zf.Zf{}
	ret := huoqujiegou(zf.Liechangdu(false))
	return ret
}
func Huoqulieleixings() map[string]string {
	zf := zf.Zf{}
	ret := huoqujiegou(zf.Lieleixing(false))
	return ret
}
func Huoquyigechangdu(lieming string) int {
	zf := zf.Zf{}
	sjk := sjk.Sjk{}
	ffm := lieming + zf.Changduzhi(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[0].Interface().(int)
	return ret
}

func Huoquyigeleixing(lieming string) string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	sjk := sjk.Sjk{}
	ffm := lieming + zf.Leixingzhi(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[0].Interface().(string)
	if ret == kzf {
		log.Println("类型错误-----空")
		return kzf
	}
	return ret
}

func Huoquyigebiaojiegou(biaoming string) map[string]string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	biaolies := Huoqubiaolies()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	jiahao := zfzhi.Jiahaozhi()
	jianhao := zfzhi.Jianhaozhi()
	cfh := zfzhi.Chengfanghaozhi()
	kzf := zfzhi.Kongzifuzhi()

	regwordstr := zkhz + zkhz + mh + zf.Word(true) + mh + zkhy + zkhy + jiahao
	regbiaostr := cfh + biaoming + regwordstr
	regbiao, _ := regexp.CompilePOSIX(regbiaostr)
	regtoustr := cfh + biaoming
	regtou, _ := regexp.CompilePOSIX(regtoustr)

	regdaxiestr := cfh + zkhz + zf.A(false) + jianhao + zf.Z(false) + zkhy
	regdaxie, _ := regexp.CompilePOSIX(regdaxiestr)
	ret := make(map[string]string)
	for biaolie, _ := range biaolies {
		if regbiao.MatchString(biaolie) {
			biaotou := regbiao.FindString(biaolie)
			mixinlie := regtou.ReplaceAllString(biaotou, kzf)
			if regdaxie.MatchString(mixinlie) {
				ret[mixinlie] = strings.ToLower(mixinlie)
			}
		}
	}
	return ret
}

func Buffertoumodel(buffer *bytes.Buffer, bianma string, baoming string) {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	// package xxx \n
	kgf := zfzhi.Konggefuzhi()
	hhf := zfzhi.Huanhangfuzhi()

	pac := zf.Package(true) + kgf + baoming + hhf
	buffer.WriteString(pac)

	// type Juese struct{\n
	typestr := zf.Type(true) + kgf + bianma + kgf + zf.Struct(true) + zfzhi.Dakuohaozuozhi() + hhf
	buffer.WriteString(typestr)
}
