package web

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

// 初始化服务
func InitApiServer() (err error) {
	var (
		mux           *http.ServeMux
		listener      net.Listener
		httpServer    *http.Server
		staticDir     http.Dir     // 静态文件根目录
		staticHandler http.Handler // 静态文件的 HTTP 回调
	)
	// 配置路由
	mux = http.NewServeMux()
	//mux.HandleFunc("/job/save", handleJobSave)
	//mux.HandleFunc("/job/delete", handleJobDelete)
	//mux.HandleFunc("/job/list", handleJobList)
	//mux.HandleFunc("/job/kill", handleJobKill)
	// 静态文件目录
	//staticDir = http.Dir(G_config.WebRoot)
	staticHandler = http.FileServer(staticDir)
	// /index.html -> index.html -> ./webroot/index.html
	mux.Handle("/", http.StripPrefix("/", staticHandler))
	// 启动 TCP 监听
	if listener, err = net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(8080)); err != nil {
		return err
	}
	// 创建 HTTP 服务
	httpServer = &http.Server{
		ReadHeaderTimeout: time.Duration(500) * time.Millisecond,
		WriteTimeout:      time.Duration(500) * time.Millisecond,
		Handler:           mux,
	}
	
	// 启动服务端
	go httpServer.Serve(listener)
	
	return nil
}
