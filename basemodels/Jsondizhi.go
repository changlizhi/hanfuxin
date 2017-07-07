package basemodels

import (
	"go/token"
	"hanfuxin/zf"
	"hanfuxin/zfmulu"
	"hanfuxin/zfwenjian"
	"hanfuxin/zfz"
	"hanfuxin/zfzhi"
	"path/filepath"
	"runtime"
)

// 获取项目所在目录，这个方法无论在个系统都可以准确获取到项目目录
func Getapppath() string {
	dh := zfzhi.Dianhaozhi()
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(
		filepath.Dir(
			filepath.Join(file, dh+dh+zfzhi.Xiexianzhi())))
	return apppath
}

// 获取文件目录，直接返回文件目录结构，不论文件是否存在
func Getwenjianmulu(mulu string, wenjian string, leixing string) string {
	path := Getapppath() + // apppath
		zfzhi.Xiexianzhi() + // /
		mulu +
		zfzhi.Xiexianzhi() + // /
		wenjian +
		zfzhi.Dianhaozhi() +
		leixing
	return path
}

// 获取两个json的目录
func getjsonpath(bianma string) string {
	zf := zfz.Zf{}
	path := Getwenjianmulu(zf.Conf(false), bianma, zf.Json(true))
	return path
}
func Getchangliangpath() string {
	zf := zfz.Zf{}
	return getjsonpath(zf.Changliang(true)) //changliang
}
func Getyingyongpath() string {
	zf := zfz.Zf{}
	return getjsonpath(zf.Yingyong(true)) //yingyong
}
