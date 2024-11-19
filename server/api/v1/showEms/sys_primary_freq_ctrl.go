package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
	showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PrimaryFreqCtrlApi struct{}

// CreatePrimaryFreqCtrl 创建一次调频API
// @Tags PrimaryFreqCtrl
// @Summary 创建一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.PrimaryFreqCtrl true "创建一次调频API"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /PriFreqCtrl/createPrimaryFreqCtrl [post]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) CreatePrimaryFreqCtrl(c *gin.Context) {
	var PriFreqCtrl showEms.PrimaryFreqCtrl
	err := c.ShouldBindJSON(&PriFreqCtrl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	PriFreqCtrl.CreatedBy = utils.GetUserID(c)
	err = PriFreqCtrlService.CreatePrimaryFreqCtrl(&PriFreqCtrl)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccss"), c)
}

// DeletePrimaryFreqCtrl 删除一次调频API
// @Tags PrimaryFreqCtrl
// @Summary 删除一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.PrimaryFreqCtrl true "删除一次调频API"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /PriFreqCtrl/deletePrimaryFreqCtrl [delete]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) DeletePrimaryFreqCtrl(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := PriFreqCtrlService.DeletePrimaryFreqCtrl(ID, userID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deletFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeletePrimaryFreqCtrlByIds 批量删除一次调频API
// @Tags PrimaryFreqCtrl
// @Summary 批量删除一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /PriFreqCtrl/deletePrimaryFreqCtrlByIds [delete]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) DeletePrimaryFreqCtrlByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := PriFreqCtrlService.DeletePrimaryFreqCtrlByIds(IDs, userID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("sys_operation_record.batchDeleteSuccess"), c)
}

// UpdatePrimaryFreqCtrl 更新一次调频API
// @Tags PrimaryFreqCtrl
// @Summary 更新一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.PrimaryFreqCtrl true "更新一次调频API"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /PriFreqCtrl/updatePrimaryFreqCtrl [put]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) UpdatePrimaryFreqCtrl(c *gin.Context) {
	var PriFreqCtrl showEms.PrimaryFreqCtrl
	err := c.ShouldBindJSON(&PriFreqCtrl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	PriFreqCtrl.UpdatedBy = utils.GetUserID(c)
	err = PriFreqCtrlService.UpdatePrimaryFreqCtrl(PriFreqCtrl)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindPrimaryFreqCtrl 用id查询一次调频API
// @Tags PrimaryFreqCtrl
// @Summary 用id查询一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.PrimaryFreqCtrl true "用id查询一次调频API"
// @Success 200 {object} response.Response{data=showEms.PrimaryFreqCtrl,msg=string} "查询成功"
// @Router /PriFreqCtrl/findPrimaryFreqCtrl [get]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) FindPrimaryFreqCtrl(c *gin.Context) {
	ID := c.Query("ID")
	rePriFreqCtrl, err := PriFreqCtrlService.GetPrimaryFreqCtrl(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(rePriFreqCtrl, c)
}

// GetPrimaryFreqCtrlList 分页获取一次调频API列表
// @Tags PrimaryFreqCtrl
// @Summary 分页获取一次调频API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.PrimaryFreqCtrlSearch true "分页获取一次调频API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /PriFreqCtrl/getPrimaryFreqCtrlList [get]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) GetPrimaryFreqCtrlList(c *gin.Context) {
	var pageInfo showEmsReq.PrimaryFreqCtrlSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := PriFreqCtrlService.GetPrimaryFreqCtrlInfoList(pageInfo)
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

// GetPrimaryFreqCtrlPublic 不需要鉴权的一次调频API接口
// @Tags PrimaryFreqCtrl
// @Summary 不需要鉴权的一次调频API接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.PrimaryFreqCtrlSearch true "分页获取一次调频API列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PriFreqCtrl/getPrimaryFreqCtrlPublic [get]
func (PriFreqCtrlApi *PrimaryFreqCtrlApi) GetPrimaryFreqCtrlPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	PriFreqCtrlService.GetPrimaryFreqCtrlPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的一次调频API接口信息",
	}, "获取成功", c)
}
