package showEms

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PrimaryFreqCtrlApi
	EmsToolsApi
	OtherApiApi
}

var (
	PriFreqCtrlService = service.ServiceGroupApp.ShowEmsServiceGroup.PrimaryFreqCtrlService
	EmsService         = service.ServiceGroupApp.ShowEmsServiceGroup.EmsToolsService
	OtherApiStService  = service.ServiceGroupApp.ShowEmsServiceGroup.OtherApiService
)
