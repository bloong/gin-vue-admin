// 自动生成模板FaultCode
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 故障码 结构体  FaultCode
type FaultCode struct {
	global.GVA_MODEL
	FaultID          *string `json:"faultID" form:"faultID" gorm:"column:fault_id;comment:;" binding:"required"`                            //故障ID
	FaultLevel       *string `json:"faultLevel" form:"faultLevel" gorm:"column:fault_level;comment:;" binding:"required"`                   //故障等级
	FaultDescription *string `json:"faultDescription" form:"faultDescription" gorm:"column:fault_description;comment:;" binding:"required"` //故障描述
	Solution         *string `json:"solution" form:"solution" gorm:"column:solution;comment:;" binding:"required"`                          //解决方案
}

// TableName 故障码 FaultCode自定义表名 FaultCodeManagement
func (FaultCode) TableName() string {
	return "FaultCodeManagement"
}
