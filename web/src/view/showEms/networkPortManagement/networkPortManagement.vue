
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item :label="t('general.createDate')" prop="createdAt">
      <template #label>
        <span>
          {{t('general.createDate')}}
          <el-tooltip :content="t('general.searchDesc')">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" :placeholder="t('general.startData')" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" :placeholder="t('general.endData')" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">{{t('general.search')}}</el-button>
          <el-button icon="refresh" @click="onReset">{{t('general.reset')}}</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">{{t('general.expand')}}</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>{{t('general.collapse')}}</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">{{t('general.add')}}</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">{{t('general.delete')}}</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" :label="t('general.createdAt')" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.Port')" prop="port" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.port,NetPortOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.Enabled')" prop="enabled" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.enabled) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.DHCP')" prop="dhcp" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.dhcp) }}</template>
        </el-table-column>
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.IP')" prop="ip" width="120" />
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.Netmask')" prop="netmask" width="120" />
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.Gateway')" prop="gateway" width="120" />
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.DNS1')" prop="dns1" width="120" />
        <el-table-column align="left" :label="t('showEms.NetworkPortManagement.DNS2')" prop="dns2" width="120" />

        <el-table-column align="left" :label="t('general.operations')" fixed="right" min-width="240">
          <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>{{t('general.desc')}}</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateNetworkPortManagementFunc(scope.row)">{{t('general.change')}}</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">{{t('general.delete')}}</el-button>
          </template>

        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?t('general.add'):t('general.edit')}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">{{t('general.confirm')}}</el-button>
                  <el-button @click="closeDialog">{{t('general.close')}}</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item :label="t('showEms.NetworkPortManagement.Port')"  prop="port" >
              <el-select v-model="formData.port" :placeholder="t('showEms.NetworkPortManagement.Port')" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in NetPortOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.Enabled')"  prop="enabled" >
              <el-switch v-model="formData.enabled" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('general.yes')" :inactive-text="t('general.no')" clearable ></el-switch>
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.DHCP')"  prop="dhcp" >
              <el-switch v-model="formData.dhcp" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('general.yes')" :inactive-text="t('general.no')" clearable ></el-switch>
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.IP')"  prop="ip" >
              <el-input v-model="formData.ip" :clearable="false"  :placeholder="t('showEms.NetworkPortManagement.IP')" />
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.Netmask')"  prop="netmask" >
              <el-input v-model="formData.netmask" :clearable="false"  :placeholder="t('showEms.NetworkPortManagement.Netmask')" />
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.Gateway')"  prop="gateway" >
              <el-input v-model="formData.gateway" :clearable="false"  :placeholder="t('showEms.NetworkPortManagement.Gateway')" />
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.DNS1')"  prop="dns1" >
              <el-input v-model="formData.dns1" :clearable="false"  :placeholder="t('showEms.NetworkPortManagement.DNS1')" />
            </el-form-item>
            <el-form-item :label="t('showEms.NetworkPortManagement.DNS2')"  prop="dns2" >
              <el-input v-model="formData.dns2" :clearable="false"  :placeholder="t('showEms.NetworkPortManagement.DNS2')" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createNetworkPortManagement,
  deleteNetworkPortManagement,
  deleteNetworkPortManagementByIds,
  updateNetworkPortManagement,
  findNetworkPortManagement,
  getNetworkPortManagementList
} from '@/api/showEms/networkPortManagement'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()



defineOptions({
    name: 'NetworkPortManagement'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const NetPortOptions = ref([])
const formData = ref({
            port: '',
            enabled: false,
            dhcp: false,
            ip: '',
            netmask: '',
            gateway: '',
            dns1: '',
            dns2: '',
        })



// 验证规则
const rule = reactive({
               port : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               enabled : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               dhcp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               netmask : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               gateway : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               dns1 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
               dns2 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: t('general.noOnlySpace'),
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error(t('general.placeInputEndData')))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error(t('general.placeInputStartData')))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error(t('general.startDataMustBeforeEndData')))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.enabled === ""){
        searchInfo.value.enabled=null
    }
    if (searchInfo.value.dhcp === ""){
        searchInfo.value.dhcp=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getNetworkPortManagementList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    NetPortOptions.value = await getDictFunc('NetPort')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
        type: 'warning'
    }).then(() => {
            deleteNetworkPortManagementFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm(t('general.deleteConfirm'), t('general.hint'), {
    confirmButtonText: t('general.confirm'),
    cancelButtonText: t('general.cancel'),
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: t('general.selectDataToDelete')
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteNetworkPortManagementByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: t('general.deleteSuccess')
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateNetworkPortManagementFunc = async(row) => {
    const res = await findNetworkPortManagement({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteNetworkPortManagementFunc = async (row) => {
    const res = await deleteNetworkPortManagement({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: t('general.deleteSuccess')
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        port: '',
        enabled: false,
        dhcp: false,
        ip: '',
        netmask: '',
        gateway: '',
        dns1: '',
        dns2: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createNetworkPortManagement(formData.value)
                  break
                case 'update':
                  res = await updateNetworkPortManagement(formData.value)
                  break
                default:
                  res = await createNetworkPortManagement(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: t('general.createUpdateSuccess')
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findNetworkPortManagement({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
 