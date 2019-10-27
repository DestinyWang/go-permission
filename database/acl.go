package database

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
)

type Acl struct {
	Id          int64  `json:"id",gorm:"id",validate:"required"`                     // 主键
	Code        string `json:"code",gorm:"code",validate:"required"`                 //
	Name        string `json:"name",gorm:"name",validate:"required"`                 // 名称
	AclModuleId int64  `json:"aclModuleId",gorm:"acl_module_id",validate:"required"` // 所属权限模块 id
	Url         string `json:"url",gorm:"url",validate:"required"`                   // 权限 URL, type 为菜单时必填, 按钮和其他允许为空
	Type        int    `json:"type",gorm:"type",validate:"required"`                 // 菜单, 按钮, 其他
	Status      int    `json:"status",gorm:"status",validate:"min=0,max=2"`          // 0-冻结, 1-正常
	Seq         int    `json:"seq",gorm:"seq",validate:"required"`                   // 当前模块下的排序
	Remark      string `json:"remark",gorm:"remark",validate:"required"`             // 备注
	Operator    string `json:"operator",gorm:"operator",validate:"required"`         // 操作者
	OperateTime int64  `json:"operateTime",grom:"operate_time",validate:"required"`  // 操作时间
	OperateIp   string `json:"operateIp",gorm:"operate_ip",validate:"required"`      // 操作者 ip
}

func (acl *Acl) TableName() string {
	return "sys_acl"
}

func AddAcl(acl *Acl) (err error) {
	if err = util.Db.Create(acl).Error; err != nil {
		logrus.WithError(err).Errorf("acl [%+v] insert fail", acl)
	}
	return
}

func FindById(id int64) (acl *Acl, err error) {
	acl = new(Acl)
	if err = util.Db.First(acl, id).Error; err != nil {
		logrus.WithError(err).Error("find by id fail")
	}
	return acl, err
}
