package core

import (
	"github.com/huerni/gmitex/core/gateway/config"
	"github.com/huerni/gmitex/core/logger"
	"net/http"
	"strconv"
	"time"
)

// APIGatewayServer API网关服务启动
type APIGatewayServer struct {
	timeout     time.Duration
	host        string
	port        int
	contextPath string
}

// NewAPIGatewayServer 创建 API Gateway 服务对象
func NewAPIGatewayServer(conf *config.ServerConfig) *APIGatewayServer {
	return &APIGatewayServer{
		timeout:     time.Second * conf.Timeout,
		host:        conf.Host,
		port:        conf.Port,
		contextPath: conf.ContextPath,
	}
}

// Start 启动 API 服务
// 所有的请求都会经过此处进行分发
func (r *APIGatewayServer) Start(routers *config.Routers) {
	gatewayProxy := NewGatewayProxy(routers)
	http.HandleFunc(r.contextPath, gatewayProxy.dispatch)
	//http.Handle("/favicon.ico", http.FileServer(http.Dir("static")))

	// 服务监听地址
	addr := r.host + ":" + strconv.Itoa(r.port)
	logger.Info("服务启动中，服务绑定端口号：", r.port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Fatal("API 网关服务启动失败", err)
	}
}
