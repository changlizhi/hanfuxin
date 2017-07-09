package sjk

import (
	"hanfuxin/zf"
	"hanfuxin/zfzhi"
)

func (sjk *Sjk) Idleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Id(false); z != zfzhi.Kongzifuzhi() {
		return zf.Int(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Biaojileixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Biaoji(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Bianmaleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Bianma(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Mingchengleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Mingcheng(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Juesebianmaleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Juesebianma(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Quanxianbianmaleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Quanxianbianma(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Lujingleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Lujing(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Youxiangleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Youxiang(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Leixingleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Leixing(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Zhileixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Zhi(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
func (sjk *Sjk) Xinxibianmaleixingzhi() string {
	zf := zf.Zf{}
	zfzhi := zfzhi.Zfzhi{}
	if z := sjk.Xinxibianma(false); z != zfzhi.Kongzifuzhi() {
		return zf.String(true)
	}
	return zfzhi.Kongzifuzhi()
}
