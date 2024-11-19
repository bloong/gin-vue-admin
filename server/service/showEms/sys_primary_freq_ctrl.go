package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
    showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
    "gorm.io/gorm"
)

type PrimaryFreqCtrlService struct {}
// CreatePrimaryFreqCtrl 创建一次调频API记录
// Author [yourname](https://github.com/yourname)
func (PriFreqCtrlService *PrimaryFreqCtrlService) CreatePrimaryFreqCtrl(PriFreqCtrl *showEms.PrimaryFreqCtrl) (err error) {
	err = global.GVA_DB.Create(PriFreqCtrl).Error
	return err
}

// DeletePrimaryFreqCtrl 删除一次调频API记录
// Author [yourname](https://github.com/yourname)
func (PriFreqCtrlService *PrimaryFreqCtrlService)DeletePrimaryFreqCtrl(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&showEms.PrimaryFreqCtrl{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&showEms.PrimaryFreqCtrl{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeletePrimaryFreqCtrlByIds 批量删除一次调频API记录
// Author [yourname](https://github.com/yourname)
func (PriFreqCtrlService *PrimaryFreqCtrlService)DeletePrimaryFreqCtrlByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&showEms.PrimaryFreqCtrl{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&showEms.PrimaryFreqCtrl{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdatePrimaryFreqCtrl 更新一次调频API记录
// Author [yourname](https://github.com/yourname)
func (PriFreqCtrlService *PrimaryFreqCtrlService)UpdatePrimaryFreqCtrl(PriFreqCtrl showEms.PrimaryFreqCtrl) (err error) {
	err = global.GVA_DB.Model(&showEms.PrimaryFreqCtrl{}).Where("id = ?",PriFreqCtrl.ID).Updates(&PriFreqCtrl).Error
	return err
}

// GetPrimaryFreqCtrl 根据ID获取一次调频API记录
// Author [yourname](https://github.com/yourname)
func (PriFreqCtrlService *PrimaryFreqCtrlService)GetPrimaryFreqCtrl(ID string) (PriFreqCtrl showEms.PrimaryFreqCtrl, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&PriFreqCtrl).Error
	return
}

// GetPrimaryFreqCtrlInfoList 分页获取一次调频API记录
// Author [yourname](https://github.com/yourname)
func (PriFreqCtrlService *PrimaryFreqCtrlService)GetPrimaryFreqCtrlInfoList(info showEmsReq.PrimaryFreqCtrlSearch) (list []showEms.PrimaryFreqCtrl, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&showEms.PrimaryFreqCtrl{})
    var PriFreqCtrls []showEms.PrimaryFreqCtrl
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

	err = db.Find(&PriFreqCtrls).Error
	return  PriFreqCtrls, total, err
}
func (PriFreqCtrlService *PrimaryFreqCtrlService)GetPrimaryFreqCtrlPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
