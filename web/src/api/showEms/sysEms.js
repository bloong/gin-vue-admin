import service from '@/utils/request'

// @Tags EmsTools
// @Summary 不需要鉴权的EMS通用API接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.EmsToolsSearch true "分页获取EMS通用API列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /Ems/getEmsToolsPublic [get]
export const getEmsToolsPublic = () => {
  return service({
    url: '/Ems/getEmsToolsPublic',
    method: 'get',
  })
}

 