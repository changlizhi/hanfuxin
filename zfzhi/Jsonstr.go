package zfzhi

func (zfzhi *Zfzhi) Postjuesezhi() string {
	ret := `{"Id":1,"Bianma":"ROLE_ADMIN","Mingcheng":"管理员角色","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postjuesequanxianzhi() string {
	ret := `{"Id":1,"Juesebianma":"ROLE_ADMIN","Quanxianbianma":"SRC_VIEW","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postquanxianzhi() string {
	ret := `{"Id":1,"Bianma":"SRC_VIEW","Mingcheng":"供浏览的资源","Lujing":"/juese","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postxinxizhi() string {
	ret := `{"Id":1,"Bianma":"USER_CLZ","Mingcheng":"用户clz","Youxiang":"clz@aa.com","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postxinxijuesezhi() string {
	ret := `{"Id":1,"Xinxibianma":"USER_CLZ","Juesebianma":"ROLE_ADMIN","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postyanzhengzhi() string {
	ret := `{"Id":1,"Bianma":"USER_CLZ","Mingcheng":"验证1","Leixing":"Mima","Zhi":"123456","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Postyanzhengleixingzhi() string {
	ret := `{"Id":1,"Bianma":"Mima","Mingcheng":"密码验证","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchjuesezhi() string {
	ret := `{"Id":1,"Bianma":"ROLE_ADMINxiugai","Mingcheng":"管理员角色xiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchjuesequanxianzhi() string {
	ret := `{"Id":1,"Juesebianma":"ROLE_ADMINxiugai","Quanxianbianma":"SRC_VIEWxiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchquanxianzhi() string {
	ret := `{"Id":1,"Bianma":"SRC_VIEWxiugai","Mingcheng":"供浏览的资源xiugai","Lujing":"/juesexiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchxinxizhi() string {
	ret := `{"Id":1,"Bianma":"USER_CLZxiugai","Mingcheng":"用户clzxiugai","Youxiang":"clz@aa.com","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchxinxijuesezhi() string {
	ret := `{"Id":1,"Xinxibianma":"USER_CLZxiugai","Juesebianma":"ROLE_ADMINxiugai","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchyanzhengzhi() string {
	ret := `{"Id":1,"Bianma":"USER_CLZxiugai","Mingcheng":"验证1xiugai","Leixing":"Mimaxiugai","Zhi":"123456","Biaoji":"Youxiao"}`
	return ret
}
func (zfzhi *Zfzhi) Patchyanzhengleixingzhi() string {
	ret := `{"Id":1,"Bianma":"Mimaxiugai","Mingcheng":"密码验证xiugaxiugaii","Biaoji":"Youxiao"}`
	return ret
}
