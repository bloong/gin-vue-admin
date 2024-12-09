import service from '@/utils/request'
// @Tags SerialPortManagement
// @Summary 创建串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SerialPortManagement true "创建串口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /SPM/createSerialPortManagement [post]
export const createSerialPortManagement = (data) => {
  return service({
    url: '/SPM/createSerialPortManagement',
    method: 'post',
    data
  })
}

// @Tags SerialPortManagement
// @Summary 删除串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SerialPortManagement true "删除串口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SPM/deleteSerialPortManagement [delete]
export const deleteSerialPortManagement = (params) => {
  return service({
    url: '/SPM/deleteSerialPortManagement',
    method: 'delete',
    params
  })
}

// @Tags SerialPortManagement
// @Summary 批量删除串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除串口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SPM/deleteSerialPortManagement [delete]
export const deleteSerialPortManagementByIds = (params) => {
  return service({
    url: '/SPM/deleteSerialPortManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags SerialPortManagement
// @Summary 更新串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SerialPortManagement true "更新串口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SPM/updateSerialPortManagement [put]
export const updateSerialPortManagement = (data) => {
  return service({
    url: '/SPM/updateSerialPortManagement',
    method: 'put',
    data
  })
}

// @Tags SerialPortManagement
// @Summary 用id查询串口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SerialPortManagement true "用id查询串口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SPM/findSerialPortManagement [get]
export const findSerialPortManagement = (params) => {
  return service({
    url: '/SPM/findSerialPortManagement',
    method: 'get',
    params
  })
}

// @Tags SerialPortManagement
// @Summary 分页获取串口列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取串口列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SPM/getSerialPortManagementList [get]
export const getSerialPortManagementList = (params) => {
  return service({
    url: '/SPM/getSerialPortManagementList',
    method: 'get',
    params
  })
}

// @Tags SerialPortManagement
// @Summary 不需要鉴权的串口接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.SerialPortManagementSearch true "分页获取串口列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /SPM/getSerialPortManagementPublic [get]
export const getSerialPortManagementPublic = () => {
  return service({
    url: '/SPM/getSerialPortManagementPublic',
    method: 'get',
  })
}
