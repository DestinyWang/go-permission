package database

import (
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
)

type Department struct {
	Id          int64  `json:"id",gorm:"id"`                    // 主键
	Name        string `json:"name",gorm:"name"`                // 名称
	Level       string `json:"level",gorm:"level"`              // 层级
	Seq         int    `json:"seq",gorm:"seq"`                  // 顺序
	Remark      string `json:"remark",gorm:"remark"`            // 备注
	ParentId    int64  `json:"parentId",gorm:"parent_id"`       // 上级部门 id
	Operator    string `json:"operator",gorm:"operator"`        // 操作者
	OperateTime int64  `json:"operateTime",gorm:"operate_time"` // 操作时间
	OperateIp   string `json:"operateIp",gorm:"operate_ip"`     // 操作者 ip
}

func (dept *Department) TableName() string {
	return "sys_dept"
}

// 判断同级下部门是否重复
func CheckExistDB(parentId int64, name string, deptId int64) bool {
	return false
}

func GetByIdDB(id int64) (department *Department, err error) {
	department = new(Department)
	if err := util.Db.First(department, id).Error; err != nil {
		logrus.WithError(err).Errorf("get level fail: id=[%d]", id)
	}
	return department, err
}

func AddDepartmentDB(department *Department) (err error) {
	if err = util.Db.Create(department).Error; err != nil {
		logrus.WithError(err).Errorf("add department fail: department=[%+v]", department)
	}
	return err
}

func GetAllDept() (depts []*Department, err error) {
	if err = util.Db.Find(&depts).Error; err != nil {
		logrus.WithError(err).Error("get all departments fail")
	}
	return depts, err
}
