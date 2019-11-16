package model

// TODO validate
type DeptVO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" validate:"require"`
	ParentId int64  `json:"parentId"`
	Seq      int    `json:"seq"`
	Remark   string `json:"remark"`
}

type UserVO struct {
	Username string `json:"username"`
	Telephone string `json:"telephone"`
	Mail string `json:"mail"`
	Password string `json:"password"`
	DeptId int64 `json:"deptId"`
	Remark string `json:"remark"`
}
