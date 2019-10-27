package database

type Role struct {
	Id          int64  `json:"id",gorm:"id"`
	Name        string `json:"name",gorm:"name"`                // 角色名称
	Type        int    `json:"type",gorm:"type"`                // 角色类型
	Status      int    `json:"status",gorm:"status"`            // 角色状态: 0-停用, 1-生效
	Remark      string `json:"remark",gorm:"remark"`            // 备注
	Operator    string `json:"operator",gorm:"operator"`        // 操作者
	OperateTime int64  `json:"operateTime",gorm:"operate_time"` // 操作时间
	OperateIp   string `json:"operateIp",gorm:"operate_ip"`     // 操作者 ip
}
