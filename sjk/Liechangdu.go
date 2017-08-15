package sjk

import (
	"changliang/zfzhi"
)

func (sjk *Sjk) Idchangdu() int {
	return zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Biaojichangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Bianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Mingchengchangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Juesebianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Quanxianbianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Lujingchangdu() int {
	return zfzhi.Zhi.Shuzi3() * zfzhi.Zhi.Shuzi10() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Youxiangchangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Leixingchangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Zhichangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
func (sjk *Sjk) Xinxibianmachangdu() int {
	return zfzhi.Zhi.Shuzi5() * zfzhi.Zhi.Shuzi10()
}
