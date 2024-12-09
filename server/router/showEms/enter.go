package showEms

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	PrimaryFreqCtrlRouter
	EmsToolsRouter
	OtherApiRouter
	FaultCodeRouter
	SerialPortManagementRouter
	NetworkPortManagementRouter
	DITableRouter
	OperationSettingRouter
}

var (
	PriFreqCtrlApi = api.ApiGroupApp.ShowEmsApiGroup.PrimaryFreqCtrlApi
	EmsApi         = api.ApiGroupApp.ShowEmsApiGroup.EmsToolsApi
	OtherApiStApi  = api.ApiGroupApp.ShowEmsApiGroup.OtherApiApi
	FCMApi         = api.ApiGroupApp.ShowEmsApiGroup.FaultCodeApi
	SPMApi         = api.ApiGroupApp.ShowEmsApiGroup.SerialPortManagementApi
	NPMApi         = api.ApiGroupApp.ShowEmsApiGroup.NetworkPortManagementApi
	DIApi          = api.ApiGroupApp.ShowEmsApiGroup.DITableApi
	DOApi          = api.ApiGroupApp.ShowEmsApiGroup.OperationSettingApi
)
