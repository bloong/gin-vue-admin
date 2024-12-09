package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(showEms.PrimaryFreqCtrl{}, showEms.OtherApi{}, showEms.FaultCode{}, showEms.SerialPortManagement{}, showEms.NetworkPortManagement{}, showEms.DITable{}, showEms.OperationSetting{})
	if err != nil {
		return err
	}
	return nil
}
