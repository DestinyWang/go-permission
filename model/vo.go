package model

// TODO validate
type DeptVO struct {
	Id int64
	Name string `validate:"require"`
	ParentId int64
	Seq int
	Remark string
}
