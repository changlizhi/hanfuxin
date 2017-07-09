package basemodels

import (
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
	"path/filepath"
	"runtime"
)

// 获取项目所在目录，这个方法无论在个系统都可以准确获取到项目目录
func Getapppath() string {
	zfzhi := zfzhi.Zfzhi{}
	dh := zfzhi.Dianhaozhi()
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(
		filepath.Dir(
			filepath.Join(file, dh+dh+zfzhi.Xiexianzhi())))
	return apppath
}

// 获取文件目录，直接返回文件目录结构，不论文件是否存在
func Getwenjianmulu(mulu string, wenjian string, leixing string) string {
	zfzhi := zfzhi.Zfzhi{}
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
	zf := zf.Zf{}
	path := Getwenjianmulu(zf.Conf(false), bianma, zf.Json(true))
	return path
}
func Getchangliangpath() string {
	zf := zf.Zf{}
	return getjsonpath(zf.Changliang(true)) //changliang
}
func Getyingyongpath() string {
	zf := zf.Zf{}
	return getjsonpath(zf.Yingyong(true)) //yingyong
}
