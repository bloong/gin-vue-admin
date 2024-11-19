import service from '@/utils/request'
// @Tags OtherApi
// @Summary 创建OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OtherApi true "创建OtherApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /OtherApiSt/createOtherApi [post]
export const createOtherApi = (data) => {
  return service({
    url: '/OtherApiSt/createOtherApi',
    method: 'post',
    data
  })
}

// @Tags OtherApi
// @Summary 删除OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OtherApi true "删除OtherApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /OtherApiSt/deleteOtherApi [delete]
export const deleteOtherApi = (params) => {
  return service({
    url: '/OtherApiSt/deleteOtherApi',
    method: 'delete',
    params
  })
}

// @Tags OtherApi
// @Summary 批量删除OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OtherApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /OtherApiSt/deleteOtherApi [delete]
export const deleteOtherApiByIds = (params) => {
  return service({
    url: '/OtherApiSt/deleteOtherApiByIds',
    method: 'delete',
    params
  })
}

// @Tags OtherApi
// @Summary 更新OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OtherApi true "更新OtherApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /OtherApiSt/updateOtherApi [put]
export const updateOtherApi = (data) => {
  return service({
    url: '/OtherApiSt/updateOtherApi',
    method: 'put',
    data
  })
}

// @Tags OtherApi
// @Summary 用id查询OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OtherApi true "用id查询OtherApi"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /OtherApiSt/findOtherApi [get]
export const findOtherApi = (params) => {
  return service({
    url: '/OtherApiSt/findOtherApi',
    method: 'get',
    params
  })
}

// @Tags OtherApi
// @Summary 分页获取OtherApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OtherApi列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /OtherApiSt/getOtherApiList [get]
export const getOtherApiList = (params) => {
  return service({
    url: '/OtherApiSt/getOtherApiList',
    method: 'get',
    params
  })
}

// @Tags OtherApi
// @Summary 不需要鉴权的OtherApi接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "分页获取OtherApi列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /OtherApiSt/getOtherApiPublic [get]
export const getOtherApiPublic = () => {
  return service({
    url: '/OtherApiSt/getOtherApiPublic',
    method: 'get',
  })
}
//  view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/GetConsulKey [GET]
export const GetConsulKey = (params) => {
  return service({
    url: '/OtherApiSt/GetConsulKey',
    method: 'GET',
    params
  })
}
// SetConsulKey view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/SetConsulKey [PUT]
export const SetConsulKey = (data) => {
  return service({
    url: '/OtherApiSt/SetConsulKey',
    method: 'PUT',
    data
  })
}
// SetJsonfile view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/SetJsonfile [POST]
export const SetJsonfile = () => {
  return service({
    url: '/OtherApiSt/SetJsonfile',
    method: 'POST'
  })
}
// GetJsonfile view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/GetJsonfile [GET]
export const GetJsonfile = () => {
  return service({
    url: '/OtherApiSt/GetJsonfile',
    method: 'GET'
  })
}
