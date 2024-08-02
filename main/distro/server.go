package distro

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/telanflow/domain-deploy/web"
)

func LoadCore() *gin.Engine {
	// 生产环境
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	app.MaxMultipartMemory = 30 << 20 // 8 MiB

	// 中间件
	corsMiddleware := cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"*"},
	})

	// 路由加载
	basicGroup := app.Group("/")
	basicGroup.Use(corsMiddleware)
	{
		// 证书部署
		basicGroup.POST("/issueCertificate", web.IssueCertificateHandler)
	}

	return app
}
