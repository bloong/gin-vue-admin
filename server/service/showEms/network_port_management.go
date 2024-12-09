
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
)

type NetworkPortManagementService struct {}
// CreateNetworkPortManagement 创建网口管理记录
// Author [yourname](https://github.com/yourname)
func (NPMService *NetworkPortManagementService) CreateNetworkPortManagement(NPM *showEms.NetworkPortManagement) (err error) {
	err = global.GVA_DB.Create(NPM).Error
	return err
}

// DeleteNetworkPortManagement 删除网口管理记录
// Author [yourname](https://github.com/yourname)
func (NPMService *NetworkPortManagementService)DeleteNetworkPortManagement(ID string) (err error) {
	err = global.GVA_DB.Delete(&showEms.NetworkPortManagement{},"id = ?",ID).Error
	return err
}

// DeleteNetworkPortManagementByIds 批量删除网口管理记录
// Author [yourname](https://github.com/yourname)
func (NPMService *NetworkPortManagementService)DeleteNetworkPortManagementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]showEms.NetworkPortManagement{},"id in ?",IDs).Error
	return err
}

// UpdateNetworkPortManagement 更新网口管理记录
// Author [yourname](https://github.com/yourname)
func (NPMService *NetworkPortManagementService)UpdateNetworkPortManagement(NPM showEms.NetworkPortManagement) (err error) {
	err = global.GVA_DB.Model(&showEms.NetworkPortManagement{}).Where("id = ?",NPM.ID).Updates(&NPM).Error
	return err
}

// GetNetworkPortManagement 根据ID获取网口管理记录
// Author [yourname](https://github.com/yourname)
func (NPMService *NetworkPortManagementService)GetNetworkPortManagement(ID string) (NPM showEms.NetworkPortManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&NPM).Error
	return
}

// GetNetworkPortManagementInfoList 分页获取网口管理记录
// Author [yourname](https://github.com/yourname)
func (NPMService *NetworkPortManagementService)GetNetworkPortManagementInfoList(info showEmsReq.NetworkPortManagementSearch) (list []showEms.NetworkPortManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&showEms.NetworkPortManagement{})
    var NPMs []showEms.NetworkPortManagement
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

	err = db.Find(&NPMs).Error
	return  NPMs, total, err
}
func (NPMService *NetworkPortManagementService)GetNetworkPortManagementPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}