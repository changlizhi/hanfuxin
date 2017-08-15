package appinits

import (
	"changliang/zf"
	"gongju"
)

func Shezhijson() *Shezhi {
	shezhi := Shezhi{}
	obj := gongju.Jiexi(Shezhipath(), &shezhi)
	return obj.(*Shezhi)
}
func Guojihuajson(yuyan string) *Guojihua {
	guojihua := Guojihua{}
	obj := gongju.Jiexi(Guojihuapath(yuyan), &guojihua)
	return obj.(*Guojihua)
}

func Guojihuapath(yuyan string) string {
	path := gongju.Getwenjianmulu(zf.Zfs.Hanfuxin(true), zf.Zfs.Conf(false), yuyan, zf.Zfs.Json(true))
	return path
}
func Shezhipath() string {
	path := gongju.Getwenjianmulu(zf.Zfs.Hanfuxin(true), zf.Zfs.Conf(false), zf.Zfs.Shezhi(false), zf.Zfs.Json(true))
	return path
}
