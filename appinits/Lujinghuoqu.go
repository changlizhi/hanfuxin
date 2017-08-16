package appinits

import (
	"gongju"
	"changliang/zf"
)

func Getapppath() string {
	return gongju.Getpath(zf.Zfs.Hanfuxin(true))
}