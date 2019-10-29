package main

import (
	"github.com/DestinyWang/go-permission/model"
	"github.com/DestinyWang/go-permission/database"
	"github.com/DestinyWang/go-permission/service"
	"github.com/DestinyWang/go-permission/util"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"net/http"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.GET("/hello", hello)
	router.GET("/dept/tree.json", deptTree)
	router.POST("/dept/add.json", addDept)
	router.POST("/dept/update.json", deptTree)
	return router
}

// test
func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func addDept(c *gin.Context) {
	var (
		deptVO  *model.DeptVO
		reqBody []byte
		err     error
	)
	reqBody, err = util.LogReq(c)
	if err = jsoniter.Unmarshal(reqBody, &deptVO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	logrus.WithField("deptVO", deptVO).Info("add dept")
	operator, operateTime, operateIp := util.Operator(c)
	dept := &database.Department{
		Name:        deptVO.Name,
		Seq:         deptVO.Seq,
		Remark:      deptVO.Remark,
		ParentId:    deptVO.ParentId,
		Operator:    operator,
		OperateTime: operateTime,
		OperateIp:   operateIp,
	}
	if err = service.AddDepartment(dept); err != nil {
		logrus.WithError(err).Errorf("add service fail: deptVO=[%+v]", deptVO)
		c.JSON(http.StatusInternalServerError, util.Fail(err.Error()))
	} else {
		c.JSON(http.StatusOK, util.Success("add service success"))
	}
}

func deptTree(c *gin.Context) {
	var (
		//reqBody []byte
		err           error
		deptLevelDTOs []*model.DeptLevelDTO
	)
	_, _ = util.LogReq(c)
	if deptLevelDTOs, err = service.DepartmentTree(); err != nil {
		logrus.WithError(err).Error("get service tree fail")
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, deptLevelDTOs)
}

//func updateDept
