package database

type Role struct {
	Id          int64
	Name        string // 角色名称
	Type        int    // 角色类型
	Status      int    // 角色状态: 0-停用, 1-生效
	Remark      string // 备注
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}
