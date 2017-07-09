package sjk

import (
	"hanfuxin/zfzhi"
)

func (sjk *Sjk) Idchangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Biaojichangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Bianmachangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Mingchengchangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Juesebianmachangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Quanxianbianmachangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Lujingchangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi3zhi() * zfzhi.Shuzi10zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Youxiangchangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Leixingchangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Zhichangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
func (sjk *Sjk) Xinxibianmachangduzhi() int {
	zfzhi := zfzhi.Zfzhi{}
	return zfzhi.Shuzi5zhi() * zfzhi.Shuzi10zhi()
}
