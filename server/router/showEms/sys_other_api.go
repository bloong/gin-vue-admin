package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OtherApiRouter struct{}

func (s *OtherApiRouter) InitOtherApiRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	OtherApiStRouter := Router.Group("OtherApiSt").Use(middleware.OperationRecord())
	OtherApiStRouterWithoutRecord := Router.Group("OtherApiSt")
	OtherApiStRouterWithoutAuth := PublicRouter.Group("OtherApiSt")
	{
		OtherApiStRouter.POST("createOtherApi", OtherApiStApi.CreateOtherApi)
		OtherApiStRouter.DELETE("deleteOtherApi", OtherApiStApi.DeleteOtherApi)
		OtherApiStRouter.DELETE("deleteOtherApiByIds", OtherApiStApi.DeleteOtherApiByIds)
		OtherApiStRouter.PUT("updateOtherApi", OtherApiStApi.UpdateOtherApi)

		OtherApiStRouter.GET("GetConsulKey", OtherApiStApi.GetConsulKey)
		OtherApiStRouter.PUT("SetConsulKey", OtherApiStApi.SetConsulKey)
		OtherApiStRouter.POST("SetJsonfile", OtherApiStApi.SetJsonfile)
		OtherApiStRouter.GET("GetJsonfile", OtherApiStApi.GetJsonfile)
	}
	{
		OtherApiStRouterWithoutRecord.GET("findOtherApi", OtherApiStApi.FindOtherApi)
		OtherApiStRouterWithoutRecord.GET("getOtherApiList", OtherApiStApi.GetOtherApiList)
	}
	{
		OtherApiStRouterWithoutAuth.GET("getOtherApiPublic", OtherApiStApi.GetOtherApiPublic)
	}
}
