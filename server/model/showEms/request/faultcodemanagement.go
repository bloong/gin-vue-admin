
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type FaultCodeSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    FaultID  *string `json:"faultID" form:"faultID" `
    FaultLevel  *string `json:"faultLevel" form:"faultLevel" `
    FaultDescription  *string `json:"faultDescription" form:"faultDescription" `
    Solution  *string `json:"solution" form:"solution" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
