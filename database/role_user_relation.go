package database

type RoleUser struct {
	Id          int64  //
	UserId      int64  // 用户 id
	RoleId      int64  // 角色 id
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}
