package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OperationSettingRouter struct {}

// InitOperationSettingRouter 初始化 DO设置 路由信息
func (s *OperationSettingRouter) InitOperationSettingRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	DORouter := Router.Group("DO").Use(middleware.OperationRecord())
	DORouterWithoutRecord := Router.Group("DO")
	DORouterWithoutAuth := PublicRouter.Group("DO")
	{
		DORouter.POST("createOperationSetting", DOApi.CreateOperationSetting)   // 新建DO设置
		DORouter.DELETE("deleteOperationSetting", DOApi.DeleteOperationSetting) // 删除DO设置
		DORouter.DELETE("deleteOperationSettingByIds", DOApi.DeleteOperationSettingByIds) // 批量删除DO设置
		DORouter.PUT("updateOperationSetting", DOApi.UpdateOperationSetting)    // 更新DO设置
	}
	{
		DORouterWithoutRecord.GET("findOperationSetting", DOApi.FindOperationSetting)        // 根据ID获取DO设置
		DORouterWithoutRecord.GET("getOperationSettingList", DOApi.GetOperationSettingList)  // 获取DO设置列表
	}
	{
	    DORouterWithoutAuth.GET("getOperationSettingPublic", DOApi.GetOperationSettingPublic)  // DO设置开放接口
	}
}
