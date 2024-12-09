
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
)

type FaultCodeService struct {}
// CreateFaultCode 创建故障码记录
// Author [yourname](https://github.com/yourname)
func (FCMService *FaultCodeService) CreateFaultCode(FCM *showEms.FaultCode) (err error) {
	err = global.GVA_DB.Create(FCM).Error
	return err
}

// DeleteFaultCode 删除故障码记录
// Author [yourname](https://github.com/yourname)
func (FCMService *FaultCodeService)DeleteFaultCode(ID string) (err error) {
	err = global.GVA_DB.Delete(&showEms.FaultCode{},"id = ?",ID).Error
	return err
}

// DeleteFaultCodeByIds 批量删除故障码记录
// Author [yourname](https://github.com/yourname)
func (FCMService *FaultCodeService)DeleteFaultCodeByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]showEms.FaultCode{},"id in ?",IDs).Error
	return err
}

// UpdateFaultCode 更新故障码记录
// Author [yourname](https://github.com/yourname)
func (FCMService *FaultCodeService)UpdateFaultCode(FCM showEms.FaultCode) (err error) {
	err = global.GVA_DB.Model(&showEms.FaultCode{}).Where("id = ?",FCM.ID).Updates(&FCM).Error
	return err
}

// GetFaultCode 根据ID获取故障码记录
// Author [yourname](https://github.com/yourname)
func (FCMService *FaultCodeService)GetFaultCode(ID string) (FCM showEms.FaultCode, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&FCM).Error
	return
}

// GetFaultCodeInfoList 分页获取故障码记录
// Author [yourname](https://github.com/yourname)
func (FCMService *FaultCodeService)GetFaultCodeInfoList(info showEmsReq.FaultCodeSearch) (list []showEms.FaultCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&showEms.FaultCode{})
    var FCMs []showEms.FaultCode
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.FaultID != nil && *info.FaultID != "" {
        db = db.Where("fault_id LIKE ?","%"+*info.FaultID+"%")
    }
    if info.FaultLevel != nil && *info.FaultLevel != "" {
        db = db.Where("fault_level = ?",*info.FaultLevel)
    }
    if info.FaultDescription != nil && *info.FaultDescription != "" {
        db = db.Where("fault_description LIKE ?","%"+*info.FaultDescription+"%")
    }
    if info.Solution != nil && *info.Solution != "" {
        db = db.Where("solution LIKE ?","%"+*info.Solution+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["fault_id"] = true
         	orderMap["fault_level"] = true
         	orderMap["fault_description"] = true
         	orderMap["solution"] = true
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

	err = db.Find(&FCMs).Error
	return  FCMs, total, err
}
func (FCMService *FaultCodeService)GetFaultCodePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}