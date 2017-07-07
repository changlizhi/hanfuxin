package zfz

import (
	"hanfuxin/zfzhi"
	"runtime"
	"strings"
)

type Zf struct {
}

func Fangfaming(xiaoxie bool) string {
	pc, _, _, _ := runtime.Caller(1)
	ff := runtime.FuncForPC(pc)
	f := strings.Split(ff.Name(), zfzhi.Dianhaozhi())[2]
	if xiaoxie {
		return strings.ToLower(f)
	}
	return f
}
func (zf *Zf) Hanfuxin(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Moren(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Github(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Changdu(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Shuzu(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Test(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Now(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Null(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Zd(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Genggai(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Shi(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Fou(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Youxiao(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Wuxiao(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Error01(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Error02(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
func (zf *Zf) Error03(xiaoxie bool) string {
	return Fangfaming(xiaoxie)
}
