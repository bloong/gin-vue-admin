import service from '@/utils/request'
// @Tags NetworkPortManagement
// @Summary 创建网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NetworkPortManagement true "创建网口管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /NPM/createNetworkPortManagement [post]
export const createNetworkPortManagement = (data) => {
  return service({
    url: '/NPM/createNetworkPortManagement',
    method: 'post',
    data
  })
}

// @Tags NetworkPortManagement
// @Summary 删除网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NetworkPortManagement true "删除网口管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NPM/deleteNetworkPortManagement [delete]
export const deleteNetworkPortManagement = (params) => {
  return service({
    url: '/NPM/deleteNetworkPortManagement',
    method: 'delete',
    params
  })
}

// @Tags NetworkPortManagement
// @Summary 批量删除网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除网口管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /NPM/deleteNetworkPortManagement [delete]
export const deleteNetworkPortManagementByIds = (params) => {
  return service({
    url: '/NPM/deleteNetworkPortManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags NetworkPortManagement
// @Summary 更新网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.NetworkPortManagement true "更新网口管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /NPM/updateNetworkPortManagement [put]
export const updateNetworkPortManagement = (data) => {
  return service({
    url: '/NPM/updateNetworkPortManagement',
    method: 'put',
    data
  })
}

// @Tags NetworkPortManagement
// @Summary 用id查询网口管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.NetworkPortManagement true "用id查询网口管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /NPM/findNetworkPortManagement [get]
export const findNetworkPortManagement = (params) => {
  return service({
    url: '/NPM/findNetworkPortManagement',
    method: 'get',
    params
  })
}

// @Tags NetworkPortManagement
// @Summary 分页获取网口管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取网口管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /NPM/getNetworkPortManagementList [get]
export const getNetworkPortManagementList = (params) => {
  return service({
    url: '/NPM/getNetworkPortManagementList',
    method: 'get',
    params
  })
}

// @Tags NetworkPortManagement
// @Summary 不需要鉴权的网口管理接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.NetworkPortManagementSearch true "分页获取网口管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /NPM/getNetworkPortManagementPublic [get]
export const getNetworkPortManagementPublic = () => {
  return service({
    url: '/NPM/getNetworkPortManagementPublic',
    method: 'get',
  })
}
