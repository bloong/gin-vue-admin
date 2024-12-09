package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DITableRouter struct {}

// InitDITableRouter 初始化 接口信息 路由信息
func (s *DITableRouter) InitDITableRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	DIRouter := Router.Group("DI").Use(middleware.OperationRecord())
	DIRouterWithoutRecord := Router.Group("DI")
	DIRouterWithoutAuth := PublicRouter.Group("DI")
	{
		DIRouter.POST("createDITable", DIApi.CreateDITable)   // 新建接口信息
		DIRouter.DELETE("deleteDITable", DIApi.DeleteDITable) // 删除接口信息
		DIRouter.DELETE("deleteDITableByIds", DIApi.DeleteDITableByIds) // 批量删除接口信息
		DIRouter.PUT("updateDITable", DIApi.UpdateDITable)    // 更新接口信息
	}
	{
		DIRouterWithoutRecord.GET("findDITable", DIApi.FindDITable)        // 根据ID获取接口信息
		DIRouterWithoutRecord.GET("getDITableList", DIApi.GetDITableList)  // 获取接口信息列表
	}
	{
	    DIRouterWithoutAuth.GET("getDITablePublic", DIApi.GetDITablePublic)  // 接口信息开放接口
	}
}
