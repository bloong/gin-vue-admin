package showEms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type FaultCodeApi struct {}



// CreateFaultCode 创建故障码
// @Tags FaultCode
// @Summary 创建故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.FaultCode true "创建故障码"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /FCM/createFaultCode [post]
func (FCMApi *FaultCodeApi) CreateFaultCode(c *gin.Context) {
	var FCM showEms.FaultCode
	err := c.ShouldBindJSON(&FCM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FCMService.CreateFaultCode(&FCM)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteFaultCode 删除故障码
// @Tags FaultCode
// @Summary 删除故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.FaultCode true "删除故障码"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /FCM/deleteFaultCode [delete]
func (FCMApi *FaultCodeApi) DeleteFaultCode(c *gin.Context) {
	ID := c.Query("ID")
	err := FCMService.DeleteFaultCode(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteFaultCodeByIds 批量删除故障码
// @Tags FaultCode
// @Summary 批量删除故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /FCM/deleteFaultCodeByIds [delete]
func (FCMApi *FaultCodeApi) DeleteFaultCodeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := FCMService.DeleteFaultCodeByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateFaultCode 更新故障码
// @Tags FaultCode
// @Summary 更新故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.FaultCode true "更新故障码"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /FCM/updateFaultCode [put]
func (FCMApi *FaultCodeApi) UpdateFaultCode(c *gin.Context) {
	var FCM showEms.FaultCode
	err := c.ShouldBindJSON(&FCM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FCMService.UpdateFaultCode(FCM)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindFaultCode 用id查询故障码
// @Tags FaultCode
// @Summary 用id查询故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.FaultCode true "用id查询故障码"
// @Success 200 {object} response.Response{data=showEms.FaultCode,msg=string} "查询成功"
// @Router /FCM/findFaultCode [get]
func (FCMApi *FaultCodeApi) FindFaultCode(c *gin.Context) {
	ID := c.Query("ID")
	reFCM, err := FCMService.GetFaultCode(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reFCM, c)
}

// GetFaultCodeList 分页获取故障码列表
// @Tags FaultCode
// @Summary 分页获取故障码列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.FaultCodeSearch true "分页获取故障码列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /FCM/getFaultCodeList [get]
func (FCMApi *FaultCodeApi) GetFaultCodeList(c *gin.Context) {
	var pageInfo showEmsReq.FaultCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := FCMService.GetFaultCodeInfoList(pageInfo)
	if err != nil {
	    global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
        response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, global.Translate("general.getDataSuccess"), c)
}

// GetFaultCodePublic 不需要鉴权的故障码接口
// @Tags FaultCode
// @Summary 不需要鉴权的故障码接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.FaultCodeSearch true "分页获取故障码列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /FCM/getFaultCodePublic [get]
func (FCMApi *FaultCodeApi) GetFaultCodePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    FCMService.GetFaultCodePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的故障码接口信息",
    }, "获取成功", c)
}
