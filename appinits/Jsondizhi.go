package appinits

import (
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"path/filepath"
	"runtime"
)

// 获取项目所在目录，这个方法无论在个系统都可以准确获取到项目目录
func Getapppath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(
		filepath.Dir(
			filepath.Join(file, zfzhi.Zhi.Dh()+zfzhi.Zhi.Dh()+zfzhi.Zhi.Xx())))
	return apppath
}

// 获取文件目录，直接返回文件目录结构，不论文件是否存在
func Getwenjianmulu(mulu string, wenjian string, leixing string) string {
	path := Getapppath() + // apppath
		zfzhi.Zhi.Xx() + // /
		mulu +
		zfzhi.Zhi.Xx() + // /
		wenjian +
		zfzhi.Zhi.Dh() +
		leixing
	return path
}

// 获取两个json的目录
func getjsonpath(bianma string) string {
	path := Getwenjianmulu(zf.Zfs.Conf(false), bianma, zf.Zfs.Json(true))
	return path
}
func Guojihuapath(yuyan string) string {
	return getjsonpath(yuyan) //
}
func Shezhipath() string {
	return getjsonpath(zf.Zfs.Shezhi(true)) //yingyong
}
