import service from '@/utils/request'
// @Tags PrimaryFreqCtrl
// @Summary 创建一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PrimaryFreqCtrl true "创建一次调频API"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /PriFreqCtrl/createPrimaryFreqCtrl [post]
export const createPrimaryFreqCtrl = (data) => {
  return service({
    url: '/PriFreqCtrl/createPrimaryFreqCtrl',
    method: 'post',
    data
  })
}

// @Tags PrimaryFreqCtrl
// @Summary 删除一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PrimaryFreqCtrl true "删除一次调频API"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PriFreqCtrl/deletePrimaryFreqCtrl [delete]
export const deletePrimaryFreqCtrl = (params) => {
  return service({
    url: '/PriFreqCtrl/deletePrimaryFreqCtrl',
    method: 'delete',
    params
  })
}

// @Tags PrimaryFreqCtrl
// @Summary 批量删除一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除一次调频API"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /PriFreqCtrl/deletePrimaryFreqCtrl [delete]
export const deletePrimaryFreqCtrlByIds = (params) => {
  return service({
    url: '/PriFreqCtrl/deletePrimaryFreqCtrlByIds',
    method: 'delete',
    params
  })
}

// @Tags PrimaryFreqCtrl
// @Summary 更新一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PrimaryFreqCtrl true "更新一次调频API"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /PriFreqCtrl/updatePrimaryFreqCtrl [put]
export const updatePrimaryFreqCtrl = (data) => {
  return service({
    url: '/PriFreqCtrl/updatePrimaryFreqCtrl',
    method: 'put',
    data
  })
}

// @Tags PrimaryFreqCtrl
// @Summary 用id查询一次调频API
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.PrimaryFreqCtrl true "用id查询一次调频API"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /PriFreqCtrl/findPrimaryFreqCtrl [get]
export const findPrimaryFreqCtrl = (params) => {
  return service({
    url: '/PriFreqCtrl/findPrimaryFreqCtrl',
    method: 'get',
    params
  })
}

// @Tags PrimaryFreqCtrl
// @Summary 分页获取一次调频API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取一次调频API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /PriFreqCtrl/getPrimaryFreqCtrlList [get]
export const getPrimaryFreqCtrlList = (params) => {
  return service({
    url: '/PriFreqCtrl/getPrimaryFreqCtrlList',
    method: 'get',
    params
  })
}

// @Tags PrimaryFreqCtrl
// @Summary 不需要鉴权的一次调频API接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.PrimaryFreqCtrlSearch true "分页获取一次调频API列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /PriFreqCtrl/getPrimaryFreqCtrlPublic [get]
export const getPrimaryFreqCtrlPublic = () => {
  return service({
    url: '/PriFreqCtrl/getPrimaryFreqCtrlPublic',
    method: 'get',
  })
}
