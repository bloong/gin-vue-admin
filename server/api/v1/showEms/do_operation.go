package showEms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OperationSettingApi struct {}



// CreateOperationSetting 创建DO设置
// @Tags OperationSetting
// @Summary 创建DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.OperationSetting true "创建DO设置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /DO/createOperationSetting [post]
func (DOApi *OperationSettingApi) CreateOperationSetting(c *gin.Context) {
	var DO showEms.OperationSetting
	err := c.ShouldBindJSON(&DO)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DOService.CreateOperationSetting(&DO)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteOperationSetting 删除DO设置
// @Tags OperationSetting
// @Summary 删除DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.OperationSetting true "删除DO设置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /DO/deleteOperationSetting [delete]
func (DOApi *OperationSettingApi) DeleteOperationSetting(c *gin.Context) {
	ID := c.Query("ID")
	err := DOService.DeleteOperationSetting(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteOperationSettingByIds 批量删除DO设置
// @Tags OperationSetting
// @Summary 批量删除DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /DO/deleteOperationSettingByIds [delete]
func (DOApi *OperationSettingApi) DeleteOperationSettingByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := DOService.DeleteOperationSettingByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateOperationSetting 更新DO设置
// @Tags OperationSetting
// @Summary 更新DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.OperationSetting true "更新DO设置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /DO/updateOperationSetting [put]
func (DOApi *OperationSettingApi) UpdateOperationSetting(c *gin.Context) {
	var DO showEms.OperationSetting
	err := c.ShouldBindJSON(&DO)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DOService.UpdateOperationSetting(DO)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindOperationSetting 用id查询DO设置
// @Tags OperationSetting
// @Summary 用id查询DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.OperationSetting true "用id查询DO设置"
// @Success 200 {object} response.Response{data=showEms.OperationSetting,msg=string} "查询成功"
// @Router /DO/findOperationSetting [get]
func (DOApi *OperationSettingApi) FindOperationSetting(c *gin.Context) {
	ID := c.Query("ID")
	reDO, err := DOService.GetOperationSetting(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reDO, c)
}

// GetOperationSettingList 分页获取DO设置列表
// @Tags OperationSetting
// @Summary 分页获取DO设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OperationSettingSearch true "分页获取DO设置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /DO/getOperationSettingList [get]
func (DOApi *OperationSettingApi) GetOperationSettingList(c *gin.Context) {
	var pageInfo showEmsReq.OperationSettingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DOService.GetOperationSettingInfoList(pageInfo)
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

// GetOperationSettingPublic 不需要鉴权的DO设置接口
// @Tags OperationSetting
// @Summary 不需要鉴权的DO设置接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OperationSettingSearch true "分页获取DO设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DO/getOperationSettingPublic [get]
func (DOApi *OperationSettingApi) GetOperationSettingPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    DOService.GetOperationSettingPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的DO设置接口信息",
    }, "获取成功", c)
}
