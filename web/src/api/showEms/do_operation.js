import service from '@/utils/request'
// @Tags OperationSetting
// @Summary 创建DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperationSetting true "创建DO设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /DO/createOperationSetting [post]
export const createOperationSetting = (data) => {
  return service({
    url: '/DO/createOperationSetting',
    method: 'post',
    data
  })
}

// @Tags OperationSetting
// @Summary 删除DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperationSetting true "删除DO设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DO/deleteOperationSetting [delete]
export const deleteOperationSetting = (params) => {
  return service({
    url: '/DO/deleteOperationSetting',
    method: 'delete',
    params
  })
}

// @Tags OperationSetting
// @Summary 批量删除DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DO设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DO/deleteOperationSetting [delete]
export const deleteOperationSettingByIds = (params) => {
  return service({
    url: '/DO/deleteOperationSettingByIds',
    method: 'delete',
    params
  })
}

// @Tags OperationSetting
// @Summary 更新DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OperationSetting true "更新DO设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DO/updateOperationSetting [put]
export const updateOperationSetting = (data) => {
  return service({
    url: '/DO/updateOperationSetting',
    method: 'put',
    data
  })
}

// @Tags OperationSetting
// @Summary 用id查询DO设置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OperationSetting true "用id查询DO设置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DO/findOperationSetting [get]
export const findOperationSetting = (params) => {
  return service({
    url: '/DO/findOperationSetting',
    method: 'get',
    params
  })
}

// @Tags OperationSetting
// @Summary 分页获取DO设置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DO设置列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DO/getOperationSettingList [get]
export const getOperationSettingList = (params) => {
  return service({
    url: '/DO/getOperationSettingList',
    method: 'get',
    params
  })
}

// @Tags OperationSetting
// @Summary 不需要鉴权的DO设置接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OperationSettingSearch true "分页获取DO设置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DO/getOperationSettingPublic [get]
export const getOperationSettingPublic = () => {
  return service({
    url: '/DO/getOperationSettingPublic',
    method: 'get',
  })
}
