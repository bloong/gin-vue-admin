
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
       <el-form-item :label="t('showEms.NetworkPortManagement.Port')"  prop="port" >
           <el-select v-model="formData.port" :placeholder="t('showEms.NetworkPortManagement.Port')" style="width:100%" :clearable="false" >
              <el-option v-for="(item,key) in NetPortOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
       <el-form-item :label="t('showEms.NetworkPortManagement.Enabled')"  prop="enabled" >
          <el-switch v-model="formData.enabled" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('base.yes')" :inactive-text="t('base.no')" clearable ></el-switch>
       </el-form-item>
       <el-form-item :label="t('showEms.NetworkPortManagement.DHCP')"  prop="dhcp" >
          <el-switch v-model="formData.dhcp" active-color="#13ce66" inactive-color="#ff4949" :active-text="t('base.yes')" :inactive-text="t('base.no')" clearable ></el-switch>
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
  createNetworkPortManagement,
  updateNetworkPortManagement,
  findNetworkPortManagement
} from '@/api/showEms/networkPortManagement'

defineOptions({
    name: 'NetworkPortManagementForm'
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
               }],
               enabled : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dhcp : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               netmask : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gateway : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dns1 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dns2 : [{
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
      const res = await findNetworkPortManagement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    NetPortOptions.value = await getDictFunc('NetPort')
}

init()
// 保存按钮
const save = async() => {
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