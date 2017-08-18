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
func Guojihuajson() *Guojihua {
	guojihua := Guojihua{}
	obj := gongju.Jiexi(Guojihuapath(), &guojihua)
	return obj.(*Guojihua)
}

func Guojihuapath() string {
	return gongju.Getwenjianmulu(
		zf.Zfs.Hanfuxin(true),
		zf.Zfs.Conf(false),
		Chushihuas[zf.Zfs.Yuyan(false)].Zhi,
		zf.Zfs.Json(true),
	)
}
func Shezhipath() string {
	return gongju.Getwenjianmulu(
		zf.Zfs.Hanfuxin(true),
		zf.Zfs.Conf(false),
		zf.Zfs.Shezhi(false),
		zf.Zfs.Json(true),
	)
}
