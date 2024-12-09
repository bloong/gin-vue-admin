package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FaultCodeRouter struct {}

// InitFaultCodeRouter 初始化 故障码 路由信息
func (s *FaultCodeRouter) InitFaultCodeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	FCMRouter := Router.Group("FCM").Use(middleware.OperationRecord())
	FCMRouterWithoutRecord := Router.Group("FCM")
	FCMRouterWithoutAuth := PublicRouter.Group("FCM")
	{
		FCMRouter.POST("createFaultCode", FCMApi.CreateFaultCode)   // 新建故障码
		FCMRouter.DELETE("deleteFaultCode", FCMApi.DeleteFaultCode) // 删除故障码
		FCMRouter.DELETE("deleteFaultCodeByIds", FCMApi.DeleteFaultCodeByIds) // 批量删除故障码
		FCMRouter.PUT("updateFaultCode", FCMApi.UpdateFaultCode)    // 更新故障码
	}
	{
		FCMRouterWithoutRecord.GET("findFaultCode", FCMApi.FindFaultCode)        // 根据ID获取故障码
		FCMRouterWithoutRecord.GET("getFaultCodeList", FCMApi.GetFaultCodeList)  // 获取故障码列表
	}
	{
	    FCMRouterWithoutAuth.GET("getFaultCodePublic", FCMApi.GetFaultCodePublic)  // 故障码开放接口
	}
}
