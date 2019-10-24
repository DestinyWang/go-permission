package database

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
)

type Acl struct {
	Id          int64  `json:"id",gorm:"id"`                     // 主键
	Code        string `json:"code",gorm:"code"`                 //
	Name        string `json:"name",gorm:"name"`                 // 名称
	AclModuleId int64  `json:"aclModuleId",gorm:"acl_module_id"` // 所属权限模块 id
	Url         string `json:"url",gorm:"url"`                   // 权限 URL, type 为菜单时必填, 按钮和其他允许为空
	Type        int    `json:"type",gorm:"type"`                 // 菜单, 按钮, 其他
	Status      int    `json:"status",gorm:"status"`             // 0-冻结, 1-正常
	Seq         int    `json:"seq",gorm:"seq"`                   // 当前模块下的排序
	Remark      string `json:"remark",gorm:"remark"`             // 备注
	Operator    string `json:"operator",gorm:"operator"`         // 操作者
	OperateTime int64  `json:"operateTime",xml:"operate_time"`   // 操作时间
	OperateIp   string `json:"operateIp",gorm:"operate_ip"`      // 操作者 ip
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
