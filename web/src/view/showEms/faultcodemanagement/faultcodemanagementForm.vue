
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
  createFaultCode,
  updateFaultCode,
  findFaultCode
} from '@/api/showEms/faultcodemanagement'

defineOptions({
    name: 'FaultCodeForm'
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
               }],
               faultLevel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               faultDescription : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               solution : [{
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
      const res = await findFaultCode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    FaultLevelOptions.value = await getDictFunc('FaultLevel')
}

init()
// 保存按钮
const save = async() => {
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