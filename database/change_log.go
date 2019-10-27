package database

type ChangeLog struct {
	Id          int64  `json:"id",gorm:"id"`                    //
	Type        int    `json:"type",gorm:"type"`                // 其他所有表的类型
	TargetId    int    `json:"targetId",gorm:"target_id"`       // 所在表的主键
	OldValue    string `json:"oldValue",gorm:"old_value"`       // 旧值
	NewValue    string `json:"newValue",gorm:"new_value"`       // 新值
	Operator    string `json:"operator",gorm:"operator"`        // 操作者
	OperateTime int64  `json:"operateTime",gorm:"operate_time"` // 操作时间
	OperateIp   string `json:"operateIp",gorm:"operate_ip"`     // 操作者 ip
}
