package database

type RoleAcl struct {
	Id          int64  `json:"id",gorm:"id"`                    //
	UserId      int64  `json:"userId",gorm:"user_id"`           // 用户 id
	AclId       int64  `json:"aclId",gorm:"acl_id"`             // 角色 id
	Operator    string `json:"operator",gorm:"operator"`        // 操作者
	OperateTime int64  `json:"operateTime",gorm:"operate_time"` // 操作时间
	OperateIp   string `json:"operateIp",gorm:"operate_ip"`     // 操作者 ip
}
