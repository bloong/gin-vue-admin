package showEms

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/showEms"
	showEmsReq "github.com/flipped-aurora/gin-vue-admin/server/model/showEms/request"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type OtherApiApi struct{}

// CreateOtherApi 创建OtherApi
// @Tags OtherApi
// @Summary 创建OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.OtherApi true "创建OtherApi"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /OtherApiSt/createOtherApi [post]
func (OtherApiStApi *OtherApiApi) CreateOtherApi(c *gin.Context) {
	var OtherApiSt showEms.OtherApi
	err := c.ShouldBindJSON(&OtherApiSt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = OtherApiStService.CreateOtherApi(&OtherApiSt)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.creationFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.creationFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.createSuccess"), c)
}

// DeleteOtherApi 删除OtherApi
// @Tags OtherApi
// @Summary 删除OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.OtherApi true "删除OtherApi"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /OtherApiSt/deleteOtherApi [delete]
func (OtherApiStApi *OtherApiApi) DeleteOtherApi(c *gin.Context) {
	ID := c.Query("ID")
	err := OtherApiStService.DeleteOtherApi(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.deleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.deleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.deleteSuccess"), c)
}

// DeleteOtherApiByIds 批量删除OtherApi
// @Tags OtherApi
// @Summary 批量删除OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /OtherApiSt/deleteOtherApiByIds [delete]
func (OtherApiStApi *OtherApiApi) DeleteOtherApiByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := OtherApiStService.DeleteOtherApiByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("system.sys_operation_record.batchDeleteFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("system.sys_operation_record.batchDeleteFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("system.sys_operation_record.batchDeleteSuccess"), c)
}

// UpdateOtherApi 更新OtherApi
// @Tags OtherApi
// @Summary 更新OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body showEms.OtherApi true "更新OtherApi"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /OtherApiSt/updateOtherApi [put]
func (OtherApiStApi *OtherApiApi) UpdateOtherApi(c *gin.Context) {
	var OtherApiSt showEms.OtherApi
	err := c.ShouldBindJSON(&OtherApiSt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = OtherApiStService.UpdateOtherApi(OtherApiSt)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.updateFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.updateFailErr"), c)
		return
	}
	response.OkWithMessage(global.Translate("general.updateSuccess"), c)
}

// FindOtherApi 用id查询OtherApi
// @Tags OtherApi
// @Summary 用id查询OtherApi
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEms.OtherApi true "用id查询OtherApi"
// @Success 200 {object} response.Response{data=showEms.OtherApi,msg=string} "查询成功"
// @Router /OtherApiSt/findOtherApi [get]
func (OtherApiStApi *OtherApiApi) FindOtherApi(c *gin.Context) {
	ID := c.Query("ID")
	reOtherApiSt, err := OtherApiStService.GetOtherApi(ID)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.queryFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.queryFailErr"), c)
		return
	}
	response.OkWithData(reOtherApiSt, c)
}

// GetOtherApiList 分页获取OtherApi列表
// @Tags OtherApi
// @Summary 分页获取OtherApi列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "分页获取OtherApi列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /OtherApiSt/getOtherApiList [get]
func (OtherApiStApi *OtherApiApi) GetOtherApiList(c *gin.Context) {
	var pageInfo showEmsReq.OtherApiSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := OtherApiStService.GetOtherApiInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error(global.Translate("general.getDataFail"), zap.Error(err))
		response.FailWithMessage(global.Translate("general.getDataFailErr"), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, global.Translate("general.getDataSuccess"), c)
}

// GetOtherApiPublic 不需要鉴权的OtherApi接口
// @Tags OtherApi
// @Summary 不需要鉴权的OtherApi接口
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "分页获取OtherApi列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /OtherApiSt/getOtherApiPublic [get]
func (OtherApiStApi *OtherApiApi) GetOtherApiPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	OtherApiStService.GetOtherApiPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的OtherApi接口信息",
	}, "获取成功", c)
}

// GetConsulKey view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/GetConsulKey [GET]
func (OtherApiStApi *OtherApiApi) GetConsulKey(c *gin.Context) {
	KEY := c.Query("KEY")

	// 记录日志
	//global.GVA_LOG.Error("KEY: " + KEY)

	// 请添加自己的业务逻辑
	config := api.DefaultConfig()
	config.Address = global.GVA_CONFIG.Consul.Addr
	config.Token = global.GVA_CONFIG.Consul.Token

	// 初始化Consul客户端
	client, err := api.NewClient(config)
	if err != nil {
		global.GVA_LOG.Error("Error creating client: ")
	}

	// 获取KV存储的值
	key := KEY
	value, _, err := client.KV().Get(key, nil)
	if err != nil {
		global.GVA_LOG.Error("Error fetching key " + KEY)
		response.FailWithMessage("Error fetching key "+KEY, c)
		return
	}

	if value == nil {
		global.GVA_LOG.Error(KEY + " not found")
		response.FailWithMessage(KEY+" not found", c)
		return
	}

	// 打印获取到的值
	global.GVA_LOG.Debug(KEY + " Value is " + string(value.Value))

	response.OkWithData(string(value.Value), c)
}

// SetConsulKey view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/SetConsulKey [PUT]
func (OtherApiStApi *OtherApiApi) SetConsulKey(c *gin.Context) {

	type Data struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	}

	var ConsulKvStruct Data

	err := c.ShouldBindJSON(&ConsulKvStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	logrus.Infof("Received data: %+v", ConsulKvStruct)

	config := api.DefaultConfig()
	config.Address = global.GVA_CONFIG.Consul.Addr
	config.Token = global.GVA_CONFIG.Consul.Token

	//global.GVA_LOG.Error(global.GVA_CONFIG.Consul.Addr)
	//global.GVA_LOG.Error(global.GVA_CONFIG.Consul.Token)

	// 初始化Consul客户端
	client, err := api.NewClient(config)
	if err != nil {
		global.GVA_LOG.Error("Error creating client:" + err.Error())
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	valueStr := fmt.Sprintf("%s", ConsulKvStruct.Value)

	// 获取KV存储的值
	key := ConsulKvStruct.Key
	value := []byte(valueStr)

	// 创建 PUT 请求
	kv := client.KV()
	pair := &api.KVPair{Key: key, Value: value}

	// 设置键值对
	if _, err := kv.Put(pair, nil); err != nil {
		global.GVA_LOG.Error("Error setting key-value pair: " + err.Error())
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// SetJsonfile view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/SetJsonfile [POST]
func (OtherApiStApi *OtherApiApi) SetJsonfile(c *gin.Context) {
	type Data struct {
		Input12931     string   `json:"input12931"`
		Switch96070    bool     `json:"switch96070"`
		Checkbox63174  []int    `json:"checkbox63174"`
		Timerange47503 []string `json:"timerange47503"`
		Slider54714    int      `json:"slider54714"`
		Textarea64794  string   `json:"textarea64794"`
	}

	var TestmyStruct Data

	err := c.ShouldBindJSON(&TestmyStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 将Person实例序列化为JSON
	jsonData, err := json.Marshal(TestmyStruct)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	// 将JSON数据写入文件
	err = os.WriteFile("person.json", jsonData, 0644)
	if err != nil {
		fmt.Println("写入文件失败:", err)
		return
	}

	logrus.Infof("Received data: %+v", TestmyStruct)

	global.GVA_LOG.Info("TestStruct NAME " + TestmyStruct.Input12931)

}

// GetJsonfile view.systemTools.enterMethodDescription
// @Tags OtherApi
// @Summary view.systemTools.enterMethodDescription
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/GetJsonfile [GET]
func (OtherApiStApi *OtherApiApi) GetJsonfile(c *gin.Context) {

	type Data struct {
		Input12931     string   `json:"input12931"`
		Switch96070    bool     `json:"switch96070"`
		Checkbox63174  []int    `json:"checkbox63174"`
		Timerange47503 []string `json:"timerange47503"`
		Slider54714    int      `json:"slider54714"`
		Textarea64794  string   `json:"textarea64794"`
	}

	// 从文件中读取JSON数据
	jsonData, err := os.ReadFile("person.json")
	if err != nil {
		global.GVA_LOG.Error("读取文件失败")
		return
	}

	// 将JSON数据反序列化为Person实例
	var p Data
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		global.GVA_LOG.Error("反序列化失败")
		return
	}

	global.GVA_LOG.Error("读取成功")

	logrus.Infof("Received data: %+v\n", p)

	global.GVA_LOG.Info("TestStruct NAME " + p.Input12931)

	response.OkWithData(p, c)

}

// GetRedisKey 获取RedisKey
// @Tags OtherApi
// @Summary 获取RedisKey
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/GetRedisKey [GET]
func (OtherApiStApi *OtherApiApi) GetRedisKey(c *gin.Context) {

	KEY := c.Query("KEY")
	redis_key, err := global.GVA_REDIS.Get(context.Background(), KEY).Result()

	if err != nil {
		global.GVA_LOG.Error("RedisStoreClearError!", zap.Error(err))
		return
	}

	response.OkWithData(redis_key, c)
}

// SetRedisKey 设置RedisKey
// @Tags OtherApi
// @Summary 设置RedisKey
// @accept application/json
// @Produce application/json
// @Param data query showEmsReq.OtherApiSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /OtherApiSt/SetRedisKey [PUT]
func (OtherApiStApi *OtherApiApi) SetRedisKey(c *gin.Context) {
	// 请添加自己的业务逻辑

	type Data struct {
		Key   string `json:"Key"`
		Value string `json:"Value"`
	}

	var ConsulKvStruct Data

	err := c.ShouldBindJSON(&ConsulKvStruct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//	logrus.Infof("Received data: %+v", ConsulKvStruct)

	err = global.GVA_REDIS.Set(context.Background(), ConsulKvStruct.Key, ConsulKvStruct.Value, 0).Err()

	if err != nil {
		global.GVA_LOG.Error("RedisStoreError!", zap.Error(err))
		return
	}

	//response.OkWithData("234", c)

	response.OkWithMessage("更新成功", c)

}
