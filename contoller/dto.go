package contoller

import "github.com/DestinyWang/go-permission/database"

// 部门层级
type DeptLevelDTO struct {
	*database.Department                 // 当前部门信息
	SubDepartments       []*DeptLevelDTO // 当前部门下的所有子部门
}
