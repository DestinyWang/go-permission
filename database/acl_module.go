package database

type AclModule struct {
	Id          int64  `json:"id",gorm:"id"`// 主键
	Name        string `json:"name",gorm:"name"`// 权限模块名称
	ParentId    int64  `json:"parentId",gorm:"parent_id"`// 父模块名称
	Level       int    `json:"level",gorm:"level"`// 模块级别
	Status      int    `json:"status",gorm:"status"`// 状态
	Seq         int    `json:"seq",grom:"seq"`// 同级别所有权限模块的顺序
	Remark      string `json:"remark",gorm:"remark"`// 备注
	Operator    string `json:"operator",gorm:"operator"`// 操作者
	OperateTime int64  `json:"operate_time",gorm:"operate_time"`// 操作时间
	OperateIp   string `json:"operate_ip",gorm:"operate_ip"`// 操作者 ip
}
