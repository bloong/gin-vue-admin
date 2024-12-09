
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
)

type DITableService struct {}
// CreateDITable 创建接口信息记录
// Author [yourname](https://github.com/yourname)
func (DIService *DITableService) CreateDITable(DI *showEms.DITable) (err error) {
	err = global.GVA_DB.Create(DI).Error
	return err
}

// DeleteDITable 删除接口信息记录
// Author [yourname](https://github.com/yourname)
func (DIService *DITableService)DeleteDITable(ID string) (err error) {
	err = global.GVA_DB.Delete(&showEms.DITable{},"id = ?",ID).Error
	return err
}

// DeleteDITableByIds 批量删除接口信息记录
// Author [yourname](https://github.com/yourname)
func (DIService *DITableService)DeleteDITableByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]showEms.DITable{},"id in ?",IDs).Error
	return err
}

// UpdateDITable 更新接口信息记录
// Author [yourname](https://github.com/yourname)
func (DIService *DITableService)UpdateDITable(DI showEms.DITable) (err error) {
	err = global.GVA_DB.Model(&showEms.DITable{}).Where("id = ?",DI.ID).Updates(&DI).Error
	return err
}

// GetDITable 根据ID获取接口信息记录
// Author [yourname](https://github.com/yourname)
func (DIService *DITableService)GetDITable(ID string) (DI showEms.DITable, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&DI).Error
	return
}

// GetDITableInfoList 分页获取接口信息记录
// Author [yourname](https://github.com/yourname)
func (DIService *DITableService)GetDITableInfoList(info showEmsReq.DITableSearch) (list []showEms.DITable, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&showEms.DITable{})
    var DIs []showEms.DITable
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

	err = db.Find(&DIs).Error
	return  DIs, total, err
}
func (DIService *DITableService)GetDITablePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}