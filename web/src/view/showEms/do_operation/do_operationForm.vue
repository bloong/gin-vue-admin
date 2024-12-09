
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
       <el-form-item :label="t('showEms.OperationSetting.Interface')"  prop="interface" >
           <el-select v-model="formData.interface" :placeholder="t('showEms.OperationSetting.Interface')" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in DOOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
       <el-form-item :label="t('showEms.OperationSetting.DOInitialStatus')"  prop="doInitialStatus" >
          <el-switch v-model="formData.doInitialStatus" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('base.yes')" :inactive-text="t('base.no')" clearable ></el-switch>
       </el-form-item>
       <el-form-item :label="t('showEms.OperationSetting.Function')"  prop="function" >
          <el-input v-model="formData.function" :clearable="false"  :placeholder="t('showEms.OperationSetting.Function')" />
       </el-form-item>
       <el-form-item :label="t('showEms.OperationSetting.ActionMode')"  prop="actionMode" >
          <el-input v-model="formData.actionMode" :clearable="false"  :placeholder="t('showEms.OperationSetting.ActionMode')" />
       </el-form-item>
       <el-form-item :label="t('showEms.OperationSetting.PulseDuration')"  prop="pulseDuration" >
          <el-input v-model.number="formData.pulseDuration" :clearable="false" :placeholder="t('general.pleaseEnter')" />
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
  createOperationSetting,
  updateOperationSetting,
  findOperationSetting
} from '@/api/showEms/do_operation'

defineOptions({
    name: 'OperationSettingForm'
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
const DOOptions = ref([])
const formData = ref({
            interface: '',
            doInitialStatus: false,
            function: '',
            actionMode: '',
            pulseDuration: undefined,
        })
// 验证规则
const rule = reactive({
               interface : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               doInitialStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               function : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               actionMode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               pulseDuration : [{
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
      const res = await findOperationSetting({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    DOOptions.value = await getDictFunc('DO')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOperationSetting(formData.value)
               break
             case 'update':
               res = await updateOperationSetting(formData.value)
               break
             default:
               res = await createOperationSetting(formData.value)
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