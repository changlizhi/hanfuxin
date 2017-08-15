package appinits

import (
	"jichu/gongju"
	"clzbase/zf"
)

func Getapppath() string {
	return gongju.Getpath(zf.Zfs.Hanfuxin(true))
}