package util

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogReq(c *gin.Context) {
	//var (
	//	err error
	//	reqBody []byte
	//)
	//if reqBody, err = ioutil.ReadAll(c.Request.Body); err != nil {
	//	logrus.WithError(err).Errorf("read req body fail: reqBody=[%s]", c.Request.Body)
	//	return nil
	//}
	logrus.WithFields(logrus.Fields{
		"Method": c.Request.Method,
		"RequestHost": c.Request.Host,
		"RequestHeader": c.Request.Header,
		"RequestForm": c.Request.Form.Encode(),
		//"RequestBody": string(reqBody),
		"Cookies": c.Request.Cookies(),
		"RequestPostForm": c.Request.PostForm,
	}).Infof("url=[%s]", c.Request.URL.String())
	//return reqBody
}

func Operator(c *gin.Context) (operator string, operateTime int64, operateIp string) {
	operator = "admin"
	operateTime = CurrMillSecond()
	operateIp = c.Request.Host
	return
}
