package database

type User struct {
	Id           int64  `json:"id",gorm:"column:id"`                       // 主键
	Username     string `json:"username",gorm:"column:iusername"`          // 用户名
	Telephone    string `json:"telephone",gorm:"column:itelephone"`        // 电话号码
	Mail         string `json:"mail",gorm:"column:imail"`                  // 邮箱
	Password     string `json:"password",gorm:"column:ipassword"`          // 密码
	Remark       string `json:"remark",gorm:"column:iremark"`              // 备注
	DepartmentId int64  `json:"departmentId",gorm:"column:idepartment_id"` // 所属部门
	Status       int    `json:"status",gorm:"column:istatus"`              // 状态
	Operator     string `json:"operator",gorm:"column:ioperator"`          // 操作者
	OperateTime  int64  `json:"operateTime",gorm:"column:ioperate_time"`   // 操作时间
	OperateIp    string `json:"operateIp",gorm:"column:ioperate_ip"`       // 操作者 ip
}


