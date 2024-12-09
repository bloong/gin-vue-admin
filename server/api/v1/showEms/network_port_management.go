package showEms

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type NetworkPortManagementApi struct {}



// CreateNetworkPortManagement 创建网口管理
// @Tags NetworkPortManagement
// @Summary 创建网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.NetworkPortManagement true "创建网口管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /NPM/createNetworkPortManagement [post]
func (NPMApi *NetworkPortManagementApi) CreateNetworkPortManagement(c *gin.Context) {
	var NPM showEms.NetworkPortManagement
	err := c.ShouldBindJSON(&NPM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = NPMService.CreateNetworkPortManagement(&NPM)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
    response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteNetworkPortManagement 删除网口管理
// @Tags NetworkPortManagement
// @Summary 删除网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.NetworkPortManagement true "删除网口管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /NPM/deleteNetworkPortManagement [delete]
func (NPMApi *NetworkPortManagementApi) DeleteNetworkPortManagement(c *gin.Context) {
	ID := c.Query("ID")
	err := NPMService.DeleteNetworkPortManagement(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteNetworkPortManagementByIds 批量删除网口管理
// @Tags NetworkPortManagement
// @Summary 批量删除网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /NPM/deleteNetworkPortManagementByIds [delete]
func (NPMApi *NetworkPortManagementApi) DeleteNetworkPortManagementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := NPMService.DeleteNetworkPortManagementByIds(IDs)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateNetworkPortManagement 更新网口管理
// @Tags NetworkPortManagement
// @Summary 更新网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.NetworkPortManagement true "更新网口管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /NPM/updateNetworkPortManagement [put]
func (NPMApi *NetworkPortManagementApi) UpdateNetworkPortManagement(c *gin.Context) {
	var NPM showEms.NetworkPortManagement
	err := c.ShouldBindJSON(&NPM)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = NPMService.UpdateNetworkPortManagement(NPM)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindNetworkPortManagement 用id查询网口管理
// @Tags NetworkPortManagement
// @Summary 用id查询网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.NetworkPortManagement true "用id查询网口管理"
// @Success 200 {object} response.Response{data=showEms.NetworkPortManagement,msg=string} "查询成功"
// @Router /NPM/findNetworkPortManagement [get]
func (NPMApi *NetworkPortManagementApi) FindNetworkPortManagement(c *gin.Context) {
	ID := c.Query("ID")
	reNPM, err := NPMService.GetNetworkPortManagement(ID)
	if err != nil {
        global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reNPM, c)
}

// GetNetworkPortManagementList 分页获取网口管理列表
// @Tags NetworkPortManagement
// @Summary 分页获取网口管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.NetworkPortManagementSearch true "分页获取网口管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /NPM/getNetworkPortManagementList [get]
func (NPMApi *NetworkPortManagementApi) GetNetworkPortManagementList(c *gin.Context) {
	var pageInfo showEmsReq.NetworkPortManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := NPMService.GetNetworkPortManagementInfoList(pageInfo)
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

// GetNetworkPortManagementPublic 不需要鉴权的网口管理接口
// @Tags NetworkPortManagement
// @Summary 不需要鉴权的网口管理接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.NetworkPortManagementSearch true "分页获取网口管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NPM/getNetworkPortManagementPublic [get]
func (NPMApi *NetworkPortManagementApi) GetNetworkPortManagementPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    NPMService.GetNetworkPortManagementPublic()
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的网口管理接口信息",
    }, "获取成功", c)
}
