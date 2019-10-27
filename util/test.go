package util

func InitTestConfig() {
	_ = InitConfig("../conf")
	_ = InitDatabase()
}
