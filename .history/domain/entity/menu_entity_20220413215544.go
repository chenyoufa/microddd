package entity

type Menu struct {
	Name       string // 菜单名称
	Icon       string // 菜单图标
	Router     string // 访问路由
	ParentID   uint64 // 父级内码
	ParentPath string // 父级路径 ***
	IsShow     int    // 是否显示(1:显示 2:隐藏)
	Status     int    // 状态(1:启用 2:禁用)
	Sequence   int    // 排序值
	Memo       string // 备注
	Creator    uint64 // 创建人
}

type MenuAction struct {
	MenuID uint64 // 菜单ID
	Code   string // 动作编号
	Name   string // 动作名称
}

// type MenuActionResource struct {
// 	ActionID uint64 // 菜单动作ID
// 	Method   string // 资源请求方式(支持正则)
// 	Path     string // 资源请求路径（支持/:id匹配）
// }
