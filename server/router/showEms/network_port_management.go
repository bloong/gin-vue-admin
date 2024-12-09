package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NetworkPortManagementRouter struct {}

// InitNetworkPortManagementRouter 初始化 网口管理 路由信息
func (s *NetworkPortManagementRouter) InitNetworkPortManagementRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	NPMRouter := Router.Group("NPM").Use(middleware.OperationRecord())
	NPMRouterWithoutRecord := Router.Group("NPM")
	NPMRouterWithoutAuth := PublicRouter.Group("NPM")
	{
		NPMRouter.POST("createNetworkPortManagement", NPMApi.CreateNetworkPortManagement)   // 新建网口管理
		NPMRouter.DELETE("deleteNetworkPortManagement", NPMApi.DeleteNetworkPortManagement) // 删除网口管理
		NPMRouter.DELETE("deleteNetworkPortManagementByIds", NPMApi.DeleteNetworkPortManagementByIds) // 批量删除网口管理
		NPMRouter.PUT("updateNetworkPortManagement", NPMApi.UpdateNetworkPortManagement)    // 更新网口管理
	}
	{
		NPMRouterWithoutRecord.GET("findNetworkPortManagement", NPMApi.FindNetworkPortManagement)        // 根据ID获取网口管理
		NPMRouterWithoutRecord.GET("getNetworkPortManagementList", NPMApi.GetNetworkPortManagementList)  // 获取网口管理列表
	}
	{
	    NPMRouterWithoutAuth.GET("getNetworkPortManagementPublic", NPMApi.GetNetworkPortManagementPublic)  // 网口管理开放接口
	}
}
