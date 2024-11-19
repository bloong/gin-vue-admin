package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PrimaryFreqCtrlRouter struct{}

// InitPrimaryFreqCtrlRouter 初始化 一次调频API 路由信息
func (s *PrimaryFreqCtrlRouter) InitPrimaryFreqCtrlRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	PriFreqCtrlRouter := Router.Group("PriFreqCtrl").Use(middleware.OperationRecord())
	PriFreqCtrlRouterWithoutRecord := Router.Group("PriFreqCtrl")
	PriFreqCtrlRouterWithoutAuth := PublicRouter.Group("PriFreqCtrl")
	{
		PriFreqCtrlRouter.POST("createPrimaryFreqCtrl", PriFreqCtrlApi.CreatePrimaryFreqCtrl)             // 新建一次调频API
		PriFreqCtrlRouter.DELETE("deletePrimaryFreqCtrl", PriFreqCtrlApi.DeletePrimaryFreqCtrl)           // 删除一次调频API
		PriFreqCtrlRouter.DELETE("deletePrimaryFreqCtrlByIds", PriFreqCtrlApi.DeletePrimaryFreqCtrlByIds) // 批量删除一次调频API
		PriFreqCtrlRouter.PUT("updatePrimaryFreqCtrl", PriFreqCtrlApi.UpdatePrimaryFreqCtrl)              // 更新一次调频API
	}
	{
		PriFreqCtrlRouterWithoutRecord.GET("findPrimaryFreqCtrl", PriFreqCtrlApi.FindPrimaryFreqCtrl)       // 根据ID获取一次调频API
		PriFreqCtrlRouterWithoutRecord.GET("getPrimaryFreqCtrlList", PriFreqCtrlApi.GetPrimaryFreqCtrlList) // 获取一次调频API列表
	}
	{
		PriFreqCtrlRouterWithoutAuth.GET("getPrimaryFreqCtrlPublic", PriFreqCtrlApi.GetPrimaryFreqCtrlPublic) // 一次调频API开放接口
	}
}
