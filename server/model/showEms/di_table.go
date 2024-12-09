// 自动生成模板DITable
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 接口信息 结构体  DITable
type DITable struct {
	global.GVA_MODEL
	InterfaceName *string `json:"interfaceName" form:"interfaceName" gorm:"column:interface_name;comment:;" binding:"required"` //接口
	RealStatus    *bool   `json:"realStatus" form:"realStatus" gorm:"column:real_status;comment:;" binding:"required"`          //实时状态
	UseFor        *string `json:"useFor" form:"useFor" gorm:"column:use_for;comment:;" binding:"required"`                      //用途
}

// TableName 接口信息 DITable自定义表名 di_table
func (DITable) TableName() string {
	return "di_table"
}
