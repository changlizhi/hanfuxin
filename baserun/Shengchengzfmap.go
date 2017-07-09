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
	"strconv"
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
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	sjk := sjk.Sjk{}
	ffm := lieming + zf.Changduzhi(true)
	v := reflect.ValueOf(&sjk)
	retstr := v.MethodByName(ffm).Call(nil)[0].String()
	if retstr == kzf {
		log.Println("类型错误-----空")
		return zfzhi.Shuzifu1zhi()
	}
	ret, _ := strconv.Atoi(retstr)
	return ret
}
func Huoquyigeleixing(lieming string) string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	kzf := zfzhi.Kongzifuzhi()
	sjk := sjk.Sjk{}
	ffm := lieming + zf.Leixingzhi(true)
	v := reflect.ValueOf(&sjk)
	ret := v.MethodByName(ffm).Call(nil)[0].String()
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

func Huoqubiaolieguanlian(biaoming string, lieming string) map[string]string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	biaolies := Huoqubiaolies()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	mh := zfzhi.Maohaozhi()
	jh := zfzhi.Jiahaozhi()
	cfh := zfzhi.Chengfanghaozhi()
	kzf := zfzhi.Kongzifuzhi()

	// [[:word:]]+
	regwordstr := zkhz + zkhz + mh + zf.Word(true) + mh + zkhy + zkhy + jh
	regbiaostr := cfh + biaoming + regwordstr
	regliestr := cfh + lieming
	regtoustr := cfh + biaoming
	regbiao, _ := regexp.CompilePOSIX(regbiaostr)
	reglie, _ := regexp.CompilePOSIX(regliestr)
	regtou, _ := regexp.CompilePOSIX(regtoustr)

	ret := make(map[string]string)
	for biaolie, _ := range biaolies {
		if regbiao.MatchString(biaolie) {
			biaotou := regbiao.FindString(biaolie)
			mixinlie := regtou.ReplaceAllString(biaotou, kzf)
			if reglie.MatchString(mixinlie) {
				lieneed := reglie.FindString(mixinlie)
				ret[lieneed] = strings.ToLower(lieneed)
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
	buffer.WriteString(zf.Package(true))
	buffer.WriteString(kgf)
	buffer.WriteString(baoming)
	buffer.WriteString(hhf)

	// type Juese struct{\n
	buffer.WriteString(zf.Type(true))
	buffer.WriteString(kgf)
	buffer.WriteString(bianma)
	buffer.WriteString(kgf)
	buffer.WriteString(zf.Struct(true))
	buffer.WriteString(zfzhi.Dakuohaozuozhi())
	buffer.WriteString(hhf)
}
