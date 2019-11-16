package database

import (
	"github.com/DestinyWang/go-permission/util"
)

type User struct {
	Id           int64  `json:"id" gorm:"column:id"`                                               // 主键
	Username     string `json:"username" gorm:"column:username" validate:"require,gte=0,lte=20"`   // 用户名
	Telephone    string `json:"telephone" gorm:"column:telephone" validate:"require,gte=0,lte=13"` // 电话号码
	Mail         string `json:"mail" gorm:"column:mail" validate:"require,email,gte=0,lte=50"`     // 邮箱
	Password     string `json:"password" gorm:"column:password"`                                   // 密码
	Remark       string `json:"remark" gorm:"column:remark" validate:"gte=0,lte=200"`              // 备注
	DepartmentId int64  `json:"departmentId" gorm:"column:dept_id" validate:"require"`             // 所属部门
	Status       int    `json:"status" gorm:"column:status" validate:"required,gte=0,lte=2"`       // 状态
	Operator     string `json:"operator" gorm:"column:operator" validate:"required"`               // 操作者
	OperateTime  int64  `json:"operateTime" gorm:"column:operate_time" validate:"required"`        // 操作时间
	OperateIp    string `json:"operateIp" gorm:"column:operate_ip" validate:"required"`            // 操作者 ip
}

func (u *User) TableName() string {
	return "sys_user"
}

func AddUser(user *User) (err error) {
	//if err  = validator.New().Struct(user); err != nil {
	//	if _, ok := err.(validator.ValidationErrors); ok {
	//		return err
	//	}
	//}
	if err = util.Db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
