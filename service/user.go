package service

import (
	"github.com/DestinyWang/go-permission/database"
	"github.com/DestinyWang/go-permission/model"
	"github.com/DestinyWang/go-permission/util"
)

func AddUser(userVO *model.UserVO, operator string, operateTime int64, operateIp string) (err error) {
	user := &database.User{
		Username:     userVO.Username,
		Telephone:    userVO.Telephone,
		Mail:         userVO.Mail,
		Password:     util.GenMD5(userVO.Password),
		Remark:       userVO.Remark,
		DepartmentId: userVO.DeptId,
		Status:       1,
		Operator:     operator,
		OperateTime:  operateTime,
		OperateIp:    operateIp,
	}
	return database.AddUser(user)
}
