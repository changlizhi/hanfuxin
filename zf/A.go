package zf

import (
	"hanfuxin/zfzhi"
	"runtime"
	"strings"
)

type Zf struct {
}

var Zfs = &Zf{}

func Fangfaming(xiaoxie bool) string {
	pc, _, _, _ := runtime.Caller(zfzhi.Zhi.Shuzi1zhi())
	ff := runtime.FuncForPC(pc)
	f := strings.Split(ff.Name(), zfzhi.Zhi.Dh())[zfzhi.Zhi.Shuzi2zhi()]
	if xiaoxie {
		return strings.ToLower(f)
	}
	return f
}
