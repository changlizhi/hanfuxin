package appinits

import (
	"changliang/zf"
	"gongju"
)

func Shezhijson() *Shezhi {
	shezhi := Shezhi{}
	ret := gongju.Jiexi(
		Shezhipath(),
		&shezhi,
	)
	return ret.(*Shezhi)
}
func Guojihuajson() *Guojihua {
	guojihua := Guojihua{}
	ret := gongju.Jiexi(
		Guojihuapath(),
		&guojihua,
	)
	return ret.(*Guojihua)
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
