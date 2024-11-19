package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		showEmsRouter := router.RouterGroupApp.ShowEms
		showEmsRouter.InitPrimaryFreqCtrlRouter(privateGroup, publicGroup)
		showEmsRouter.InitEmsToolsRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		showEmsRouter.InitOtherApiRouter(privateGroup, publicGroup)
	}
}
