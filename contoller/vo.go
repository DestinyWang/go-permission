package contoller

// TODO validate
type DeptVO struct {
	Name string `validate:"require"`
	ParentId int64
	Seq int
	Remark string
}
