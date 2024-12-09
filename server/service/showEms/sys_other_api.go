package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
	showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
)

type OtherApiService struct{}

// CreateOtherApi 创建OtherApi记录
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) CreateOtherApi(OtherApiSt *showEms.OtherApi) (err error) {
	err = global.GVA_DB.Create(OtherApiSt).Error
	return err
}

// DeleteOtherApi 删除OtherApi记录
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) DeleteOtherApi(ID string) (err error) {
	err = global.GVA_DB.Delete(&showEms.OtherApi{}, "id = ?", ID).Error
	return err
}

// DeleteOtherApiByIds 批量删除OtherApi记录
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) DeleteOtherApiByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]showEms.OtherApi{}, "id in ?", IDs).Error
	return err
}

// UpdateOtherApi 更新OtherApi记录
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) UpdateOtherApi(OtherApiSt showEms.OtherApi) (err error) {
	err = global.GVA_DB.Model(&showEms.OtherApi{}).Where("id = ?", OtherApiSt.ID).Updates(&OtherApiSt).Error
	return err
}

// GetOtherApi 根据ID获取OtherApi记录
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) GetOtherApi(ID string) (OtherApiSt showEms.OtherApi, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&OtherApiSt).Error
	return
}

// GetOtherApiInfoList 分页获取OtherApi记录
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) GetOtherApiInfoList(info showEmsReq.OtherApiSearch) (list []showEms.OtherApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	var OtherApiSts []showEms.OtherApi
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&OtherApiSts).Error
	return OtherApiSts, total, err
}
func (OtherApiStService *OtherApiService) GetOtherApiPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetConsulKey view.systemTools.enterMethodDescription
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) GetConsulKey() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	return db.Error
}

// SetConsulKey view.systemTools.enterMethodDescription
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) SetConsulKey() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	return db.Error
}

// SetJsonfile view.systemTools.enterMethodDescription
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) SetJsonfile() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	return db.Error
}

// GetJsonfile view.systemTools.enterMethodDescription
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) GetJsonfile() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	return db.Error
}

// GetRedisKey 获取RedisKey
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) GetRedisKey() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	return db.Error
}

// SetRedisKey 设置RedisKey
// Author [yourname](https://github.com/yourname)
func (OtherApiStService *OtherApiService) SetRedisKey() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&showEms.OtherApi{})
	return db.Error
}
