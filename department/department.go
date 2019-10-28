package department

import (
	"fmt"
	"github.com/DestinyWang/go-permission/contoller"
	"github.com/DestinyWang/go-permission/database"
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
	"sort"
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

func Tree() (deptLevelDTOs []*contoller.DeptLevelDTO, err error) {
	var (
		deptList []*database.Department
	)
	if deptList, err = database.GetAllDept(); err != nil {
		logrus.WithError(err).Error("get all department fail")
		return deptLevelDTOs, err
	}
	for _, d := range deptList {
		deptLevelDTOs = append(deptLevelDTOs, &contoller.DeptLevelDTO{
			Department:     d,
			SubDepartments: nil,
		})
	}
	return deptListToTree(deptLevelDTOs), nil
}

// 将线性结构转换成树形结构
func deptListToTree(deptLevelDTOs []*contoller.DeptLevelDTO) (deptDTOTree []*contoller.DeptLevelDTO) {
	var (
		levelMap = make(map[string][]*contoller.DeptLevelDTO)
		rootDept []*contoller.DeptLevelDTO
	)
	if len(deptLevelDTOs) == 0 {
		return
	}
	for _, d := range deptLevelDTOs {
		levelMap[d.Level] = append(levelMap[d.Level], d)
		if d.Level == util.RootLevel {
			rootDept = append(rootDept, d)
		}
	}
	sort.Slice(rootDept, func(i, j int) bool {
		return rootDept[i].Seq < rootDept[j].Seq
	})
	transformDeptTree(deptLevelDTOs, util.RootLevel, levelMap)
	return rootDept
}

func transformDeptTree(deptLevelDTOs []*contoller.DeptLevelDTO, level string, levelMap map[string][]*contoller.DeptLevelDTO) {
	for _, d := range deptLevelDTOs {
		nextLevel := util.CalLevel(level, d.Id)
		nextDeptLevelDTOs := levelMap[nextLevel]
		if len(nextDeptLevelDTOs) > 0 {
			sort.Slice(nextDeptLevelDTOs, func(i, j int) bool {
				return nextDeptLevelDTOs[i].Seq < nextDeptLevelDTOs[j].Seq
			})
			d.SubDepartments = nextDeptLevelDTOs
			transformDeptTree(nextDeptLevelDTOs, nextLevel, levelMap)
		}
	}
}
