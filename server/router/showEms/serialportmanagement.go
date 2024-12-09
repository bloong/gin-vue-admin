package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SerialPortManagementRouter struct {}

// InitSerialPortManagementRouter 初始化 串口 路由信息
func (s *SerialPortManagementRouter) InitSerialPortManagementRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	SPMRouter := Router.Group("SPM").Use(middleware.OperationRecord())
	SPMRouterWithoutRecord := Router.Group("SPM")
	SPMRouterWithoutAuth := PublicRouter.Group("SPM")
	{
		SPMRouter.POST("createSerialPortManagement", SPMApi.CreateSerialPortManagement)   // 新建串口
		SPMRouter.DELETE("deleteSerialPortManagement", SPMApi.DeleteSerialPortManagement) // 删除串口
		SPMRouter.DELETE("deleteSerialPortManagementByIds", SPMApi.DeleteSerialPortManagementByIds) // 批量删除串口
		SPMRouter.PUT("updateSerialPortManagement", SPMApi.UpdateSerialPortManagement)    // 更新串口
	}
	{
		SPMRouterWithoutRecord.GET("findSerialPortManagement", SPMApi.FindSerialPortManagement)        // 根据ID获取串口
		SPMRouterWithoutRecord.GET("getSerialPortManagementList", SPMApi.GetSerialPortManagementList)  // 获取串口列表
	}
	{
	    SPMRouterWithoutAuth.GET("getSerialPortManagementPublic", SPMApi.GetSerialPortManagementPublic)  // 串口开放接口
	}
}
