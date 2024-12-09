package showEms

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	PrimaryFreqCtrlApi
	EmsToolsApi
	OtherApiApi
	FaultCodeApi
	SerialPortManagementApi
	NetworkPortManagementApi
	DITableApi
	OperationSettingApi
}

var (
	PriFreqCtrlService = service.ServiceGroupApp.ShowEmsServiceGroup.PrimaryFreqCtrlService
	EmsService         = service.ServiceGroupApp.ShowEmsServiceGroup.EmsToolsService
	OtherApiStService  = service.ServiceGroupApp.ShowEmsServiceGroup.OtherApiService
	FCMService         = service.ServiceGroupApp.ShowEmsServiceGroup.FaultCodeService
	SPMService         = service.ServiceGroupApp.ShowEmsServiceGroup.SerialPortManagementService
	NPMService         = service.ServiceGroupApp.ShowEmsServiceGroup.NetworkPortManagementService
	DIService          = service.ServiceGroupApp.ShowEmsServiceGroup.DITableService
	DOService          = service.ServiceGroupApp.ShowEmsServiceGroup.OperationSettingService
)
