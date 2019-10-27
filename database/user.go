package database

type User struct {
	Id           int64  `json:"id",gorm:"id"`                      // 主键
	Username     string `json:"username",gorm:"username"`          // 用户名
	Telephone    string `json:"telephone",gorm:"telephone"`        // 电话号码
	Mail         string `json:"mail",gorm:"mail"`                  // 邮箱
	Password     string `json:"password",gorm:"password"`          // 密码
	Remark       string `json:"remark",gorm:"remark"`              // 备注
	DepartmentId int64  `json:"departmentId",gorm:"department_id"` // 所属部门
	Status       int    `json:"status",gorm:"status"`              // 状态
	Operator     string `json:"operator",gorm:"operator"`          // 操作者
	OperateTime  int64  `json:"operateTime",gorm:"operate_time"`   // 操作时间
	OperateIp    string `json:"operateIp",gorm:"operate_ip"`       // 操作者 ip
}
