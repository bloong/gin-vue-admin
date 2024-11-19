// 自动生成模板PrimaryFreqCtrl
package showEms
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 一次调频API 结构体  PrimaryFreqCtrl
type PrimaryFreqCtrl struct {
    global.GVA_MODEL
    Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:名称;size:64;"`  //sex 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 一次调频API PrimaryFreqCtrl自定义表名 primary_freq_ctrl
func (PrimaryFreqCtrl) TableName() string {
    return "primary_freq_ctrl"
}
