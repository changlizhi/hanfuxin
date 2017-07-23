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
	zfzhi := zfzhi.Zfzhi{}
	pc, _, _, _ := runtime.Caller(1)
	ff := runtime.FuncForPC(pc)
	f := strings.Split(ff.Name(), zfzhi.Dianhaozhi())[2]
	if xiaoxie {
		return strings.ToLower(f)
	}
	return f
}
