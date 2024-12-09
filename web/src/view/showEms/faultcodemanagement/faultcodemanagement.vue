
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
      
        <el-form-item :label="t('showEms.FaultCode.FaultID')" prop="faultID">
         <el-input v-model="searchInfo.faultID" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
           <el-form-item :label="t('showEms.FaultCode.FaultLevel')" prop="faultLevel">
            <el-select v-model="searchInfo.faultLevel" clearable :placeholder="t('general.pleaseSelect')" @clear="()=>{searchInfo.faultLevel=undefined}">
              <el-option v-for="(item,key) in FaultLevelOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item :label="t('showEms.FaultCode.FaultDescription')" prop="faultDescription">
         <el-input v-model="searchInfo.faultDescription" :placeholder="t('general.searchCriteria')" />
        </el-form-item>
        <el-form-item :label="t('showEms.FaultCode.Solution')" prop="solution">
         <el-input v-model="searchInfo.solution" :placeholder="t('general.searchCriteria')" />
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
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog">{{t('general.add')}}</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">{{t('general.delete')}}</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="showEms_FaultCode" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="showEms_FaultCode" />
            <ImportExcel v-auth="btnAuth.importExcel" template-id="showEms_FaultCode" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" :label="t('general.createdAt')" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
        <el-table-column sortable align="left" :label="t('showEms.FaultCode.FaultID')" prop="faultID" width="120" />
        <el-table-column sortable align="left" :label="t('showEms.FaultCode.FaultLevel')" prop="faultLevel" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.faultLevel,FaultLevelOptions) }}
            </template>
        </el-table-column>
        <el-table-column sortable align="left" :label="t('showEms.FaultCode.FaultDescription')" prop="faultDescription" width="120" />
        <el-table-column sortable align="left" :label="t('showEms.FaultCode.Solution')" prop="solution" width="120" />

        <el-table-column align="left" :label="t('general.operations')" fixed="right" min-width="240">
          <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>{{t('general.desc')}}</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateFaultCodeFunc(scope.row)">{{t('general.change')}}</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">{{t('general.delete')}}</el-button>
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
            <el-form-item :label="t('showEms.FaultCode.FaultID')"  prop="faultID" >
              <el-input v-model="formData.faultID" :clearable="false"  :placeholder="t('showEms.FaultCode.FaultID')" />
            </el-form-item>
            <el-form-item :label="t('showEms.FaultCode.FaultLevel')"  prop="faultLevel" >
              <el-select v-model="formData.faultLevel" :placeholder="t('showEms.FaultCode.FaultLevel')" style="width:100%" :clearable="false" >
                <el-option v-for="(item,key) in FaultLevelOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item :label="t('showEms.FaultCode.FaultDescription')"  prop="faultDescription" >
              <el-input v-model="formData.faultDescription" :clearable="false"  :placeholder="t('showEms.FaultCode.FaultDescription')" />
            </el-form-item>
            <el-form-item :label="t('showEms.FaultCode.Solution')"  prop="solution" >
              <el-input v-model="formData.solution" :clearable="false"  :placeholder="t('showEms.FaultCode.Solution')" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item :label="t('showEms.FaultCode.FaultID')">
                        {{ detailFrom.faultID }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('showEms.FaultCode.FaultLevel')">
                        {{ detailFrom.faultLevel }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('showEms.FaultCode.FaultDescription')">
                        {{ detailFrom.faultDescription }}
                    </el-descriptions-item>
                    <el-descriptions-item :label="t('showEms.FaultCode.Solution')">
                        {{ detailFrom.solution }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createFaultCode,
  deleteFaultCode,
  deleteFaultCodeByIds,
  updateFaultCode,
  findFaultCode,
  getFaultCodeList
} from '@/api/showEms/faultcodemanagement'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'

defineOptions({
    name: 'FaultCode'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const FaultLevelOptions = ref([])
const formData = ref({
            faultID: '',
            faultLevel: '',
            faultDescription: '',
            solution: '',
        })



// 验证规则
const rule = reactive({
               faultID : [{
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
               faultLevel : [{
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
               faultDescription : [{
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
               solution : [{
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
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            faultID: 'fault_id',
            faultLevel: 'fault_level',
            faultDescription: 'fault_description',
            solution: 'solution',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

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
  const table = await getFaultCodeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    FaultLevelOptions.value = await getDictFunc('FaultLevel')
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
            deleteFaultCodeFunc(row)
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
      const res = await deleteFaultCodeByIds({ IDs })
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
const updateFaultCodeFunc = async(row) => {
    const res = await findFaultCode({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteFaultCodeFunc = async (row) => {
    const res = await deleteFaultCode({ ID: row.ID })
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
        faultID: '',
        faultLevel: '',
        faultDescription: '',
        solution: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createFaultCode(formData.value)
                  break
                case 'update':
                  res = await updateFaultCode(formData.value)
                  break
                default:
                  res = await createFaultCode(formData.value)
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
  const res = await findFaultCode({ ID: row.ID })
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
 