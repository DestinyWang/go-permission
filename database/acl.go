package database

type Acl struct {
	Id          int64  // 主键
	Code        string //
	Name        string // 名称
	AclModuleId int64  // 所属权限模块 id
	Url         string // 权限 URL, type 为菜单时必填, 按钮和其他允许为空
	Type        int    // 菜单, 按钮, 其他
	Status      int    // 0-冻结, 1-正常
	Seq         int    // 当前模块下的排序
	Remark      string // 备注
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}
