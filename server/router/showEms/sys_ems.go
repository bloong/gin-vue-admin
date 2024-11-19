package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EmsToolsRouter struct{}

// InitEmsToolsRouter 初始化 EMS通用API 路由信息
func (s *EmsToolsRouter) InitEmsToolsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	EmsRouter := Router.Group("Ems").Use(middleware.OperationRecord())
	//EmsRouterWithoutAuth := PublicRouter.Group("Ems")
	{
		EmsRouter.GET("getEmsToolsPublic", EmsApi.GetEmsToolsPublic) // EMS通用API开放接口
	}
}
