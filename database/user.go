package database

type User struct {
	Id           int64  // 主键
	Username     string // 用户名
	Telephone    string // 电话号码
	Mail         string // 邮箱
	Password     string // 密码
	Remark       string // 备注
	DepartmentId int64  // 所属部门
	Status       int    // 状态
	Operator     string // 操作者
	OperateTime  int64  // 操作时间
	OperateIp    string // 操作者 ip
}
