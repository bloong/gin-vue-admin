// 自动生成模板OtherApi
package showEms
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// OtherApi 结构体  OtherApi
type OtherApi struct {
    global.GVA_MODEL
    Test  string `json:"test" form:"test" gorm:"column:test;comment:;"`  //test 
}


// TableName OtherApi OtherApi自定义表名 other_api
func (OtherApi) TableName() string {
    return "other_api"
}
