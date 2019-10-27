package util

type JsonData struct {
	Succ bool        `json:"succ"` // 是否成功
	Msg  string      `json:"msg"`  // 错误信息
	Data interface{} `json:"data"` // 数据
}

func Success(data interface{}) *JsonData {
	jsonData := &JsonData{
		Succ: true,
		Data: data,
	}
	return jsonData
}

func Fail(msg string) *JsonData {
	jsonData := &JsonData{
		Succ: false,
		Msg:  msg,
	}
	return jsonData
}
