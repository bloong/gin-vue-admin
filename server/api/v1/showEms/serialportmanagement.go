package showEms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SerialPortManagementApi struct {}



// CreateSerialPortManagement 创建串口
// @Tags SerialPortManagement
// @Summary 创建串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.SerialPortManagement true "创建串口"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /SPM/createSerialPortManagement [post]
func (SPMApi *SerialPortManagementApi) CreateSerialPortManagement(c *gin.Context) {
	var SPM showEms.SerialPortManagement
	err := c.ShouldBindJSON(&SPM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SPMService.CreateSerialPortManagement(&SPM)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteSerialPortManagement 删除串口
// @Tags SerialPortManagement
// @Summary 删除串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.SerialPortManagement true "删除串口"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /SPM/deleteSerialPortManagement [delete]
func (SPMApi *SerialPortManagementApi) DeleteSerialPortManagement(c *gin.Context) {
	ID := c.Query("ID")
	err := SPMService.DeleteSerialPortManagement(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteSerialPortManagementByIds 批量删除串口
// @Tags SerialPortManagement
// @Summary 批量删除串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /SPM/deleteSerialPortManagementByIds [delete]
func (SPMApi *SerialPortManagementApi) DeleteSerialPortManagementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := SPMService.DeleteSerialPortManagementByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateSerialPortManagement 更新串口
// @Tags SerialPortManagement
// @Summary 更新串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.SerialPortManagement true "更新串口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /SPM/updateSerialPortManagement [put]
func (SPMApi *SerialPortManagementApi) UpdateSerialPortManagement(c *gin.Context) {
	var SPM showEms.SerialPortManagement
	err := c.ShouldBindJSON(&SPM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = SPMService.UpdateSerialPortManagement(SPM)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindSerialPortManagement 用id查询串口
// @Tags SerialPortManagement
// @Summary 用id查询串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.SerialPortManagement true "用id查询串口"
// @Success 200 {object} response.Response{data=showEms.SerialPortManagement,msg=string} "查询成功"
// @Router /SPM/findSerialPortManagement [get]
func (SPMApi *SerialPortManagementApi) FindSerialPortManagement(c *gin.Context) {
	ID := c.Query("ID")
	reSPM, err := SPMService.GetSerialPortManagement(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reSPM, c)
}

// GetSerialPortManagementList 分页获取串口列表
// @Tags SerialPortManagement
// @Summary 分页获取串口列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.SerialPortManagementSearch true "分页获取串口列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /SPM/getSerialPortManagementList [get]
func (SPMApi *SerialPortManagementApi) GetSerialPortManagementList(c *gin.Context) {
	var pageInfo showEmsReq.SerialPortManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := SPMService.GetSerialPortManagementInfoList(pageInfo)
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

// GetSerialPortManagementPublic 不需要鉴权的串口接口
// @Tags SerialPortManagement
// @Summary 不需要鉴权的串口接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.SerialPortManagementSearch true "分页获取串口列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /SPM/getSerialPortManagementPublic [get]
func (SPMApi *SerialPortManagementApi) GetSerialPortManagementPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    SPMService.GetSerialPortManagementPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的串口接口信息",
    }, "获取成功", c)
}
