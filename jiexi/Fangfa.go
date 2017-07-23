package jiexi

import (
	"hanfuxin/basemodels"
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"io/ioutil"
	"regexp"
	"strings"
)

func Pipeifangfa(canshu string, canshuleixing string, muluming string, wenjianming string) []string {
	fu1 := zfzhi.Shuzifu1zhi()

	path := basemodels.Getapppath() + zfzhi.Zhi.Xx() + muluming + zfzhi.Zhi.Xx() + wenjianming + zfzhi.Zhi.Dh() + zf.Go(true)
	b, _ := ioutil.ReadFile(path)
	//从匹配的方法名中去除前面对于的正则表达式
	// func \(zf \*Zf\)
	// func \(sjk \*Sjk\)
	// func \(zfzhi \*Zfzhi\)
	repstr := zf.Func(true) + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Fxx() + zfzhi.Zhi.Xkhz() + canshu + zfzhi.Zhi.Kgf() + zfzhi.Zhi.Fxx() + zfzhi.Zhi.Xh() + canshuleixing + zfzhi.Zhi.Fxx() + zfzhi.Zhi.Xkhy() + zfzhi.Zhi.Kgf()
	//拼接从文件中匹配方法名的正则表达式
	// func \(zf \*Zf\) [[:word:]]+
	// func \(sjk \*Sjk\) [[:word:]]+
	// func \(zfzhi \*Zfzhi\) [[:word:]]+
	regstr := repstr + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Zkhz() + zfzhi.Zhi.Mh() + zf.Word(true) + zfzhi.Zhi.Mh() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Zkhy() + zfzhi.Zhi.Jia()
	//生成正则表达式
	reg, _ := regexp.CompilePOSIX(regstr)
	repreg, _ := regexp.CompilePOSIX(repstr)
	//从文件中匹配方法名
	sfind := reg.FindAllString(string(b), fu1)
	// 匹配到的方法名合在一起
	joinstr := strings.Join(sfind, zfzhi.Zhi.Kgf())
	//去除前缀func (zf *Zf)
	srep := repreg.ReplaceAllString(joinstr, zfzhi.Zhi.Kzf())
	// 再恢复为数组返回调用
	ret := strings.Split(srep, zfzhi.Zhi.Kgf())

	return ret
}
