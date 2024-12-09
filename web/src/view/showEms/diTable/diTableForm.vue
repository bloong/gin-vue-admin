
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
       <el-form-item :label="t('showEms.DITable.InterfaceName')"  prop="interfaceName" >
           <el-select v-model="formData.interfaceName" :placeholder="t('showEms.DITable.InterfaceName')" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in DIOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
       <el-form-item :label="t('showEms.DITable.RealStatus')"  prop="realStatus" >
          <el-switch v-model="formData.realStatus" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('base.yes')" :inactive-text="t('base.no')" clearable ></el-switch>
       </el-form-item>
       <el-form-item :label="t('showEms.DITable.UseFor')"  prop="useFor" >
          <el-input v-model="formData.useFor" :clearable="false"  :placeholder="t('showEms.DITable.UseFor')" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">{{ t('general.save') }}</el-button>
          <el-button type="primary" @click="back">{{ t('general.back') }}</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDITable,
  updateDITable,
  findDITable
} from '@/api/showEms/diTable'

defineOptions({
    name: 'DITableForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()


const route = useRoute()
const router = useRouter()

const type = ref('')
const DIOptions = ref([])
const formData = ref({
            interfaceName: '',
            realStatus: false,
            useFor: '',
        })
// 验证规则
const rule = reactive({
               interfaceName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               realStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               useFor : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDITable({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    DIOptions.value = await getDictFunc('DI')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDITable(formData.value)
               break
             case 'update':
               res = await updateDITable(formData.value)
               break
             default:
               res = await createDITable(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: t('general.createUpdateSuccess')
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>