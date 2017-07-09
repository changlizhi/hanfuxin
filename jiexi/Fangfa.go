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
	//字符串获取的对象
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	// 各种符号值的定义
	xx := zfzhi.Xiexianzhi()
	dh := zfzhi.Dianhaozhi()
	fxx := zfzhi.Fanxiexianzhi()
	xh := zfzhi.Xinghaozhi()
	mh := zfzhi.Maohaozhi()
	zkhz := zfzhi.Zhongkuohaozuozhi()
	zkhy := zfzhi.Zhongkuohaoyouzhi()
	xkhz := zfzhi.Xiaokuohaozuozhi()
	xkhy := zfzhi.Xiaokuohaoyouzhi()
	jh := zfzhi.Jiahaozhi()
	kgf := zfzhi.Konggefuzhi()
	kzf := zfzhi.Kongzifuzhi()
	fu1 := zfzhi.Shuzifu1zhi()

	path := basemodels.Getapppath() + xx + muluming + xx + wenjianming + dh + zf.Go(true)
	b, _ := ioutil.ReadFile(path)
	//从匹配的方法名中去除前面对于的正则表达式
	// func \(zf \*Zf\)
	// func \(sjk \*Sjk\)
	// func \(zfzhi \*Zfzhi\)
	repstr := zf.Func(true) + kgf + fxx + xkhz + canshu + kgf + fxx + xh + canshuleixing + fxx + xkhy + kgf
	//拼接从文件中匹配方法名的正则表达式
	// func \(zf \*Zf\) [[:word:]]+
	// func \(sjk \*Sjk\) [[:word:]]+
	// func \(zfzhi \*Zfzhi\) [[:word:]]+
	regstr := repstr + zkhz + zkhz + mh + zf.Word(true) + mh + zkhy + zkhy + jh
	//生成正则表达式
	reg, _ := regexp.CompilePOSIX(regstr)
	repreg, _ := regexp.CompilePOSIX(repstr)
	//从文件中匹配方法名
	sfind := reg.FindAllString(string(b), fu1)
	// 匹配到的方法名合在一起
	joinstr := strings.Join(sfind, kgf)
	//去除前缀func (zf *Zf)
	srep := repreg.ReplaceAllString(joinstr, kzf)
	// 再恢复为数组返回调用
	ret := strings.Split(srep, kgf)

	return ret
}
