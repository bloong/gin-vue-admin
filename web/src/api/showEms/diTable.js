import service from '@/utils/request'
// @Tags DITable
// @Summary 创建接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DITable true "创建接口信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /DI/createDITable [post]
export const createDITable = (data) => {
  return service({
    url: '/DI/createDITable',
    method: 'post',
    data
  })
}

// @Tags DITable
// @Summary 删除接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DITable true "删除接口信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DI/deleteDITable [delete]
export const deleteDITable = (params) => {
  return service({
    url: '/DI/deleteDITable',
    method: 'delete',
    params
  })
}

// @Tags DITable
// @Summary 批量删除接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除接口信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /DI/deleteDITable [delete]
export const deleteDITableByIds = (params) => {
  return service({
    url: '/DI/deleteDITableByIds',
    method: 'delete',
    params
  })
}

// @Tags DITable
// @Summary 更新接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DITable true "更新接口信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /DI/updateDITable [put]
export const updateDITable = (data) => {
  return service({
    url: '/DI/updateDITable',
    method: 'put',
    data
  })
}

// @Tags DITable
// @Summary 用id查询接口信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DITable true "用id查询接口信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /DI/findDITable [get]
export const findDITable = (params) => {
  return service({
    url: '/DI/findDITable',
    method: 'get',
    params
  })
}

// @Tags DITable
// @Summary 分页获取接口信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取接口信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /DI/getDITableList [get]
export const getDITableList = (params) => {
  return service({
    url: '/DI/getDITableList',
    method: 'get',
    params
  })
}

// @Tags DITable
// @Summary 不需要鉴权的接口信息接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.DITableSearch true "分页获取接口信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /DI/getDITablePublic [get]
export const getDITablePublic = () => {
  return service({
    url: '/DI/getDITablePublic',
    method: 'get',
  })
}
