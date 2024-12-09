package showEms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type DITableApi struct {}



// CreateDITable 创建接口信息
// @Tags DITable
// @Summary 创建接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.DITable true "创建接口信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /DI/createDITable [post]
func (DIApi *DITableApi) CreateDITable(c *gin.Context) {
	var DI showEms.DITable
	err := c.ShouldBindJSON(&DI)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DIService.CreateDITable(&DI)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteDITable 删除接口信息
// @Tags DITable
// @Summary 删除接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.DITable true "删除接口信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /DI/deleteDITable [delete]
func (DIApi *DITableApi) DeleteDITable(c *gin.Context) {
	ID := c.Query("ID")
	err := DIService.DeleteDITable(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteDITableByIds 批量删除接口信息
// @Tags DITable
// @Summary 批量删除接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /DI/deleteDITableByIds [delete]
func (DIApi *DITableApi) DeleteDITableByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := DIService.DeleteDITableByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateDITable 更新接口信息
// @Tags DITable
// @Summary 更新接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.DITable true "更新接口信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /DI/updateDITable [put]
func (DIApi *DITableApi) UpdateDITable(c *gin.Context) {
	var DI showEms.DITable
	err := c.ShouldBindJSON(&DI)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = DIService.UpdateDITable(DI)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindDITable 用id查询接口信息
// @Tags DITable
// @Summary 用id查询接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.DITable true "用id查询接口信息"
// @Success 200 {object} response.Response{data=showEms.DITable,msg=string} "查询成功"
// @Router /DI/findDITable [get]
func (DIApi *DITableApi) FindDITable(c *gin.Context) {
	ID := c.Query("ID")
	reDI, err := DIService.GetDITable(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reDI, c)
}

// GetDITableList 分页获取接口信息列表
// @Tags DITable
// @Summary 分页获取接口信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.DITableSearch true "分页获取接口信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /DI/getDITableList [get]
func (DIApi *DITableApi) GetDITableList(c *gin.Context) {
	var pageInfo showEmsReq.DITableSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := DIService.GetDITableInfoList(pageInfo)
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

// GetDITablePublic 不需要鉴权的接口信息接口
// @Tags DITable
// @Summary 不需要鉴权的接口信息接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.DITableSearch true "分页获取接口信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DI/getDITablePublic [get]
func (DIApi *DITableApi) GetDITablePublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    DIService.GetDITablePublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的接口信息接口信息",
    }, "获取成功", c)
}
