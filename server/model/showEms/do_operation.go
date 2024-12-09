// 自动生成模板OperationSetting
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// DO设置 结构体  OperationSetting
type OperationSetting struct {
	global.GVA_MODEL
	Interface       *string `json:"interface" form:"interface" gorm:"column:interface;comment:;" binding:"required"`                     //接口
	DOInitialStatus *bool   `json:"doInitialStatus" form:"doInitialStatus" gorm:"column:do_initial_status;comment:;" binding:"required"` //初始状态
	Function        *string `json:"function" form:"function" gorm:"column:function;comment:;" binding:"required"`                        //功能
	ActionMode      *string `json:"actionMode" form:"actionMode" gorm:"column:action_mode;comment:;" binding:"required"`                 //动作方式
	PulseDuration   *int    `json:"pulseDuration" form:"pulseDuration" gorm:"column:pulse_duration;comment:;" binding:"required"`        //脉冲持续时间
}

// TableName DO设置 OperationSetting自定义表名 do_settings
func (OperationSetting) TableName() string {
	return "do_settings"
}
