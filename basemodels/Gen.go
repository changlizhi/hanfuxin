package basemodels

// 主要为生成第一个model而定制,golang关键字和第一个所需的字符都在这里定义好
// 应尽量保证此结构体的字段足够少,这是唯一需要手动编辑明细字段的一个文件，足够的简单就不会出错
type Gen struct {
	Package     string
	Type        string
	Struct      string
	String      string
	Kongge      string
	Huanhang    string
	Dakuohaozuo string
	Dakuohaoyou string
	Models      string
	Import      string
	Yingyongzi  string
	Xiexian     string
	Dianhao     string
	Go          string
	Appmodels   string
}
