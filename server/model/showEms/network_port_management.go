// 自动生成模板NetworkPortManagement
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 网口管理 结构体  NetworkPortManagement
type NetworkPortManagement struct {
	global.GVA_MODEL
	Port    *string `json:"port" form:"port" gorm:"column:port;comment:;" binding:"required"`          //网口
	Enabled *bool   `json:"enabled" form:"enabled" gorm:"column:enabled;comment:;" binding:"required"` //启用
	DHCP    *bool   `json:"dhcp" form:"dhcp" gorm:"column:dhcp;comment:;" binding:"required"`          //DHCP
	IP      *string `json:"ip" form:"ip" gorm:"column:ip;comment:;" binding:"required"`                //IP
	Netmask *string `json:"netmask" form:"netmask" gorm:"column:netmask;comment:;" binding:"required"` //子网掩码
	Gateway *string `json:"gateway" form:"gateway" gorm:"column:gateway;comment:;" binding:"required"` //网关
	DNS1    *string `json:"dns1" form:"dns1" gorm:"column:dns1;comment:;" binding:"required"`          //DNS1
	DNS2    *string `json:"dns2" form:"dns2" gorm:"column:dns2;comment:;" binding:"required"`          //DNS2
}

// TableName 网口管理 NetworkPortManagement自定义表名 NetworkPortManagement
func (NetworkPortManagement) TableName() string {
	return "NetworkPortManagement"
}
