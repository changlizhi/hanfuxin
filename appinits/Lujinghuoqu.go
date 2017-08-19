package appinits

import (
	"gongju"
	"changliang/zfzhi"
)

func Getapppath() string {
	return gongju.Getpath(zfzhi.Zhi.Kzf())
}