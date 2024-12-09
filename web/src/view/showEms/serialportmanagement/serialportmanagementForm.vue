
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
       <el-form-item :label="t('showEms.SerialPortManagement.SerialPortName')"  prop="serialPortName" >
           <el-select v-model="formData.serialPortName" :placeholder="t('showEms.SerialPortManagement.SerialPortName')" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in SerialPortManagementOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
       <el-form-item :label="t('showEms.SerialPortManagement.Usage')"  prop="usage" >
          <el-input v-model="formData.usage" :clearable="false"  :placeholder="t('showEms.SerialPortManagement.Usage')" />
       </el-form-item>
       <el-form-item :label="t('showEms.SerialPortManagement.BaudRate')"  prop="baudRate" >
          <el-input v-model.number="formData.baudRate" :clearable="false" :placeholder="t('general.pleaseEnter')" />
       </el-form-item>
       <el-form-item :label="t('showEms.SerialPortManagement.DataBits')"  prop="dataBits" >
          <el-input v-model.number="formData.dataBits" :clearable="false" :placeholder="t('general.pleaseEnter')" />
       </el-form-item>
       <el-form-item :label="t('showEms.SerialPortManagement.Parity')"  prop="parity" >
          <el-input v-model="formData.parity" :clearable="false"  :placeholder="t('showEms.SerialPortManagement.Parity')" />
       </el-form-item>
       <el-form-item :label="t('showEms.SerialPortManagement.StopBits')"  prop="stopBits" >
          <el-input v-model.number="formData.stopBits" :clearable="false" :placeholder="t('general.pleaseEnter')" />
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
  createSerialPortManagement,
  updateSerialPortManagement,
  findSerialPortManagement
} from '@/api/showEms/serialportmanagement'

defineOptions({
    name: 'SerialPortManagementForm'
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
const SerialPortManagementOptions = ref([])
const formData = ref({
            serialPortName: '',
            usage: '',
            baudRate: 0,
            dataBits: 0,
            parity: '',
            stopBits: 0,
        })
// 验证规则
const rule = reactive({
               serialPortName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               usage : [{
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
      const res = await findSerialPortManagement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    SerialPortManagementOptions.value = await getDictFunc('SerialPortManagement')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSerialPortManagement(formData.value)
               break
             case 'update':
               res = await updateSerialPortManagement(formData.value)
               break
             default:
               res = await createSerialPortManagement(formData.value)
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