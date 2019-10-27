package department

import (
	"fmt"
	"github.com/DestinyWang/go-permission/database"
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
)

var (
	DeptAlreadyExits = fmt.Errorf("department")
)

func AddDepartment(department *database.Department) (err error) {
	var (
		level string
	)
	if database.CheckExistDB(department.ParentId, department.Name, department.Id) {
		return fmt.Errorf("department name [%s] already exists", department.Name)
	}
	if level, err = getLevel(department.ParentId); err != nil {
		logrus.WithError(err).Errorf("get level fail: id=[%d]", department.ParentId)
	}
	department.Level = level
	return database.AddDepartmentDB(department)
}

// 根据 id 获取 level
func getLevel(parentId int64) (level string, err error) {
	var (
		parentDept *database.Department
	)
	if parentDept, err = database.GetByIdDB(parentId); err != nil {
		logrus.WithError(err).Errorf("get level fail: id=[%d]", parentId)
		return
	}
	if parentDept == nil {
		return "", nil
	}
	return util.CalLevel(parentDept.Level, parentId), nil
}
