
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
)

type OperationSettingService struct {}
// CreateOperationSetting 创建DO设置记录
// Author [yourname](https://github.com/yourname)
func (DOService *OperationSettingService) CreateOperationSetting(DO *showEms.OperationSetting) (err error) {
	err = global.GVA_DB.Create(DO).Error
	return err
}

// DeleteOperationSetting 删除DO设置记录
// Author [yourname](https://github.com/yourname)
func (DOService *OperationSettingService)DeleteOperationSetting(ID string) (err error) {
	err = global.GVA_DB.Delete(&showEms.OperationSetting{},"id = ?",ID).Error
	return err
}

// DeleteOperationSettingByIds 批量删除DO设置记录
// Author [yourname](https://github.com/yourname)
func (DOService *OperationSettingService)DeleteOperationSettingByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]showEms.OperationSetting{},"id in ?",IDs).Error
	return err
}

// UpdateOperationSetting 更新DO设置记录
// Author [yourname](https://github.com/yourname)
func (DOService *OperationSettingService)UpdateOperationSetting(DO showEms.OperationSetting) (err error) {
	err = global.GVA_DB.Model(&showEms.OperationSetting{}).Where("id = ?",DO.ID).Updates(&DO).Error
	return err
}

// GetOperationSetting 根据ID获取DO设置记录
// Author [yourname](https://github.com/yourname)
func (DOService *OperationSettingService)GetOperationSetting(ID string) (DO showEms.OperationSetting, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&DO).Error
	return
}

// GetOperationSettingInfoList 分页获取DO设置记录
// Author [yourname](https://github.com/yourname)
func (DOService *OperationSettingService)GetOperationSettingInfoList(info showEmsReq.OperationSettingSearch) (list []showEms.OperationSetting, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&showEms.OperationSetting{})
    var DOs []showEms.OperationSetting
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&DOs).Error
	return  DOs, total, err
}
func (DOService *OperationSettingService)GetOperationSettingPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}