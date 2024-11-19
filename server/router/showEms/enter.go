package showEms

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	PrimaryFreqCtrlRouter
	EmsToolsRouter
	OtherApiRouter
}

var (
	PriFreqCtrlApi = api.ApiGroupApp.ShowEmsApiGroup.PrimaryFreqCtrlApi
	EmsApi         = api.ApiGroupApp.ShowEmsApiGroup.EmsToolsApi
	OtherApiStApi  = api.ApiGroupApp.ShowEmsApiGroup.OtherApiApi
)
