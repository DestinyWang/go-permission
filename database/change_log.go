package database

type ChangeLog struct {
	Id          int64  //
	Type        int    // 其他所有表的类型
	TargetId    int    // 所在表的主键
	OldValue    string // 旧值
	NewValue    string // 新值
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}
