// 自动生成模板SerialPortManagement
package showEms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 串口 结构体  SerialPortManagement
type SerialPortManagement struct {
	global.GVA_MODEL
	SerialPortName *string `json:"serialPortName" form:"serialPortName" gorm:"column:serial_port_name;comment:串口名称;" binding:"required"` //串口名称
	Usage          *string `json:"usage" form:"usage" gorm:"column:usage;comment:串口用途;" binding:"required"`                              //用途
	BaudRate       *int    `json:"baudRate" form:"baudRate" gorm:"column:baud_rate;comment:波特率;"`                                        //波特率
	DataBits       *int    `json:"dataBits" form:"dataBits" gorm:"column:data_bits;comment:数据位;"`                                        //数据位
	Parity         *string `json:"parity" form:"parity" gorm:"column:parity;comment:校验位;"`                                               //校验位
	StopBits       *int    `json:"stopBits" form:"stopBits" gorm:"column:stop_bits;comment:停止位;"`                                        //停止位
}

// TableName 串口 SerialPortManagement自定义表名 serial_port_management
func (SerialPortManagement) TableName() string {
	return "serial_port_management"
}
