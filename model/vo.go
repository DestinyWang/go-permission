package model

// TODO validate
type DeptVO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name",validate:"require"`
	ParentId int64  `json:"parentId"`
	Seq      int    `json:"seq"`
	Remark   string `json:"remark"`
}
