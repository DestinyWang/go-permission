package database

type RoleAcl struct {
	Id          int64  //
	UserId      int64  // 用户 id
	AclId      int64  // 角色 id
	Operator    string // 操作者
	OperateTime int64  // 操作时间
	OperateIp   string // 操作者 ip
}
