package basemodels

import (
	"path/filepath"
	"runtime"
)

type jsondizhi struct {
	Xiexian    string
	Dianhao    string
	Conf       string
	Changliang string
	Yingyong   string
	Json       string
}

var jsondizhis = jsondizhi{
	Xiexian:    "/",
	Dianhao:    ".",
	Conf:       "conf",
	Changliang: "changliang",
	Yingyong:   "yingyong",
	Json:       "json",
}

func Getapppath() string {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(
		filepath.Dir(
			filepath.Join(
				file,
				jsondizhis.Dianhao+
					jsondizhis.Dianhao+
					jsondizhis.Xiexian)))
	return apppath
}
func Getwenjianmulu(mulu string, wenjian string, leixing string) string {
	path := Getapppath() +
		jsondizhis.Xiexian +
		mulu +
		jsondizhis.Xiexian +
		wenjian +
		jsondizhis.Dianhao +
		leixing
	return path
}
func getjsonpath(bianma string) string {
	path := Getwenjianmulu(jsondizhis.Conf, bianma, jsondizhis.Json)
	return path
}
func Getchangliangpath() string {
	return getjsonpath(jsondizhis.Changliang)
}
func Getyingyongpath() string {
	return getjsonpath(jsondizhis.Yingyong)
}
