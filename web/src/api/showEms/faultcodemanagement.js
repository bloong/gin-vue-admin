import service from '@/utils/request'
// @Tags FaultCode
// @Summary 创建故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaultCode true "创建故障码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /FCM/createFaultCode [post]
export const createFaultCode = (data) => {
  return service({
    url: '/FCM/createFaultCode',
    method: 'post',
    data
  })
}

// @Tags FaultCode
// @Summary 删除故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaultCode true "删除故障码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FCM/deleteFaultCode [delete]
export const deleteFaultCode = (params) => {
  return service({
    url: '/FCM/deleteFaultCode',
    method: 'delete',
    params
  })
}

// @Tags FaultCode
// @Summary 批量删除故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除故障码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FCM/deleteFaultCode [delete]
export const deleteFaultCodeByIds = (params) => {
  return service({
    url: '/FCM/deleteFaultCodeByIds',
    method: 'delete',
    params
  })
}

// @Tags FaultCode
// @Summary 更新故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FaultCode true "更新故障码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FCM/updateFaultCode [put]
export const updateFaultCode = (data) => {
  return service({
    url: '/FCM/updateFaultCode',
    method: 'put',
    data
  })
}

// @Tags FaultCode
// @Summary 用id查询故障码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FaultCode true "用id查询故障码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FCM/findFaultCode [get]
export const findFaultCode = (params) => {
  return service({
    url: '/FCM/findFaultCode',
    method: 'get',
    params
  })
}

// @Tags FaultCode
// @Summary 分页获取故障码列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取故障码列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FCM/getFaultCodeList [get]
export const getFaultCodeList = (params) => {
  return service({
    url: '/FCM/getFaultCodeList',
    method: 'get',
    params
  })
}

// @Tags FaultCode
// @Summary 不需要鉴权的故障码接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.FaultCodeSearch true "分页获取故障码列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /FCM/getFaultCodePublic [get]
export const getFaultCodePublic = () => {
  return service({
    url: '/FCM/getFaultCodePublic',
    method: 'get',
  })
}
