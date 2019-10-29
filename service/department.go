package service

import (
	"errors"
	"fmt"
	"github.com/DestinyWang/go-permission/database"
	"github.com/DestinyWang/go-permission/model"
	"github.com/DestinyWang/go-permission/util"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

var (
	DepartmentNotFound = errors.New("部门不存在")
)

func AddDepartment(department *database.Department) (err error) {
	var (
		level string
	)
	if database.CheckExistDB(department.ParentId, department.Name, department.Id) {
		return fmt.Errorf("service name [%s] already exists", department.Name)
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
	if parentDept, err = database.GetDeptById(parentId); err != nil {
		logrus.WithError(err).Errorf("get level fail: id=[%d]", parentId)
		return
	}
	if parentDept == nil {
		return "", nil
	}
	return util.CalLevel(parentDept.Level, parentId), nil
}

func DepartmentTree() (deptLevelDTOs []*model.DeptLevelDTO, err error) {
	var (
		deptList []*database.Department
	)
	if deptList, err = database.GetAllDept(); err != nil {
		logrus.WithError(err).Error("get all service fail")
		return deptLevelDTOs, err
	}
	for _, d := range deptList {
		deptLevelDTOs = append(deptLevelDTOs, &model.DeptLevelDTO{
			Department:     d,
			SubDepartments: nil,
		})
	}
	return deptListToTree(deptLevelDTOs), nil
}

// 将线性结构转换成树形结构
func deptListToTree(deptLevelDTOs []*model.DeptLevelDTO) (deptDTOTree []*model.DeptLevelDTO) {
	var (
		levelMap = make(map[string][]*model.DeptLevelDTO)
		rootDept []*model.DeptLevelDTO
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

func transformDeptTree(deptLevelDTOs []*model.DeptLevelDTO, level string, levelMap map[string][]*model.DeptLevelDTO) {
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

func checkExist(parentId int64, deptName string, deptId int64) (exists bool, err error) {
	var cnt int
	if cnt, err = database.CountByParentIdAndName(parentId, deptName, deptId); err != nil {
		return exists, err
	}
	return cnt > 0, nil
}

func UpdateDept(vo *model.DeptVO, operator string, operateIp string) (err error) {
	var (
		department  *database.Department
		parentLevel string
	)
	if department, err = database.GetDeptById(vo.Id); err != nil {
		return err
	}
	if department == nil {
		return DepartmentNotFound
	}
	// 判断当前层级是否存在相同部门
	if parentLevel, err = getLevel(department.ParentId); err != nil {
		return err
	}
	newDept := &database.Department{
		Id:          vo.Id,
		Name:        vo.Name,
		Level:       util.CalLevel(parentLevel, department.ParentId),
		Seq:         vo.Seq,
		Remark:      vo.Remark,
		ParentId:    vo.ParentId,
		Operator:    operator,
		OperateTime: util.CurrMillSecond(),
		OperateIp:   operateIp,
	}
	return updateWithChild(department, newDept)
}

func updateWithChild(before *database.Department, after *database.Department) (err error) {
	var (
		//	tx *gorm.DB
		newLevelPrefix = after.Level
		oldLevelPrefix = before.Level
		beforeChildDepts []*database.Department
	)
	//tx = util.Db.Begin()
	if err = database.UpdateDept(after); err != nil {
		return
	}
	if after.Level != before.Level {
		// 如果更新后部门 level 发生改变, 需要更新所有子部门
		if beforeChildDepts, err = database.GetChildDeptByLevel(before.Level); err != nil {
			return err
		}
		if len(beforeChildDepts) > 0 {
			for _, d := range beforeChildDepts {
				currLevel := d.Level
				if strings.Index(currLevel, oldLevelPrefix) == 0 {
					// 如果当前部门是原部门的子部门
					currLevel = newLevelPrefix + currLevel[len(oldLevelPrefix):]
					d.Level = currLevel
				}
				// TODO batch update
				if err = database.UpdateDept(d);err != nil {
					return err
				}
			}
		}
	}
	if err = database.UpdateDept(after); err != nil {
		return err
	}
	return nil
}
