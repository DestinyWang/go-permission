package database

type Department struct {
	Id          int64  // 主键
	Name        string // 名称
	Level       string // 层级
	Seq         int    // 顺序
	Remark      string // 备注
	ParentId    int    // 上级部门 id
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}