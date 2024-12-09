
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
)

type SerialPortManagementService struct {}
// CreateSerialPortManagement 创建串口记录
// Author [yourname](https://github.com/yourname)
func (SPMService *SerialPortManagementService) CreateSerialPortManagement(SPM *showEms.SerialPortManagement) (err error) {
	err = global.GVA_DB.Create(SPM).Error
	return err
}

// DeleteSerialPortManagement 删除串口记录
// Author [yourname](https://github.com/yourname)
func (SPMService *SerialPortManagementService)DeleteSerialPortManagement(ID string) (err error) {
	err = global.GVA_DB.Delete(&showEms.SerialPortManagement{},"id = ?",ID).Error
	return err
}

// DeleteSerialPortManagementByIds 批量删除串口记录
// Author [yourname](https://github.com/yourname)
func (SPMService *SerialPortManagementService)DeleteSerialPortManagementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]showEms.SerialPortManagement{},"id in ?",IDs).Error
	return err
}

// UpdateSerialPortManagement 更新串口记录
// Author [yourname](https://github.com/yourname)
func (SPMService *SerialPortManagementService)UpdateSerialPortManagement(SPM showEms.SerialPortManagement) (err error) {
	err = global.GVA_DB.Model(&showEms.SerialPortManagement{}).Where("id = ?",SPM.ID).Updates(&SPM).Error
	return err
}

// GetSerialPortManagement 根据ID获取串口记录
// Author [yourname](https://github.com/yourname)
func (SPMService *SerialPortManagementService)GetSerialPortManagement(ID string) (SPM showEms.SerialPortManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&SPM).Error
	return
}

// GetSerialPortManagementInfoList 分页获取串口记录
// Author [yourname](https://github.com/yourname)
func (SPMService *SerialPortManagementService)GetSerialPortManagementInfoList(info showEmsReq.SerialPortManagementSearch) (list []showEms.SerialPortManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&showEms.SerialPortManagement{})
    var SPMs []showEms.SerialPortManagement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["usage"] = true
         	orderMap["baud_rate"] = true
         	orderMap["data_bits"] = true
         	orderMap["parity"] = true
         	orderMap["stop_bits"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&SPMs).Error
	return  SPMs, total, err
}
func (SPMService *SerialPortManagementService)GetSerialPortManagementPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}