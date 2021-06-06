package protocol

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"os"
)

// Http3Run 注册https证书
func Http3Run(r *gin.Engine) {
	certFile := os.Getenv("CERT_FILE")
	keyFile := os.Getenv("KEY_FILE")

	if certFile != "" && keyFile != "" {
		r.Use(setHttp3Header)

		err := http3.ListenAndServe(getPort(),certFile,keyFile,r)
		if err != nil {
			log.Panicln("http3 服务器启动失败:",err)
		}
	} else {
		r.Run()
	}
}


// setHttp3Header 设置 http3 请求头
func setHttp3Header(c *gin.Context) {
	c.Header("Alternate-Protocol", "quic:8080")
	c.Next()
}

// getPort 获取环境变量中端口号（与gin默认保持一致）
func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":8080"
}