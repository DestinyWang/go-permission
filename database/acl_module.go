package database

type AclModule struct {
	Id          int64  // 主键
	Name        string // 权限模块名称
	ParentId    int64  // 父模块名称
	Level       int    // 模块级别
	Status      int    // 状态
	Seq         int    // 同级别所有权限模块的顺序
	Remark      string // 备注
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}
