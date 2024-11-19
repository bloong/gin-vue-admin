package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type EmsToolsApi struct{}

// GetEmsToolsPublic 不需要鉴权的EMS通用API接口
// @Tags EmsTools
// @Summary 不需要鉴权的EMS通用API接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.EmsToolsSearch true "分页获取EMS通用API列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Ems/getEmsToolsPublic [get]
func (EmsApi *EmsToolsApi) GetEmsToolsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	EmsService.GetEmsToolsPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的EMS通用API接口信息",
	}, "获取成功", c)
}
