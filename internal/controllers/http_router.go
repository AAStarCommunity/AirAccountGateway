package controllers

import (
	"AirAccountGateway/docs"
	"AirAccountGateway/internal/controllers/actions"
	"AirAccountGateway/internal/models/webapi/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// SetRouters 设置API路由
func SetRouters() (routers *gin.Engine) {
	routers = gin.New()

	// 生产模式配置
	//if conf.Environment.IsProduction() {
	//	gin.SetMode(gin.ReleaseMode)   // 生产模式
	//	gin.DefaultWriter = io.Discard // 禁用 gin 输出接口访问日志
	//}

	// 开发模式配置
	//if conf.Environment.IsDevelopment() {
	gin.SetMode(gin.DebugMode) // 调试模式
	buildSwagger(routers)      // 构建swagger
	//}

	// 构建路由
	buildRouters(routers) // 构建http路由

	routers.NoRoute(func(ctx *gin.Context) {
		response.GetResponse().SetHttpCode(http.StatusNotFound).FailCode(ctx, http.StatusNotFound)
	})

	return
}

// buildRouters 构建路由表
func buildRouters(router *gin.Engine) {
	r := router.Group("/api/instructions")
	r.GET("/balance", actions.Balance)
	r.POST("/bind", actions.Bind)
	r.GET("/transfer/check", actions.TransferCheck)
	r.POST("/transfer", actions.Transfer)
}

// buildSwagger 创建swagger文档
func buildSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
