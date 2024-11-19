<template>
  <el-form
    :model="formData"
    ref="vForm"
    :rules="rules"
    label-position="left"
    label-width="80px"
    size="medium"
    @submit.prevent
  >
    <el-row :gutter="20" class="form-row">
      <!-- 卡片容器 -->
      <el-col :span="24" class="card-container">
        <transition-group name="card-transition" tag="div" class="card-group">
          <el-card
            v-for="(card, index) in cards"
            :key="index"
            class="form-card"
            :class="{ 'full-width': isNarrow, 'no-border': isNarrow }"
          >
            <template #header>
              <div class="card-header">
                <span>{{ card.title }}</span>
              </div>
            </template>
            <div v-if="card.content === 'first'">
              <el-form-item
                label="频率响应"
                prop="select_freqresp"
                class="required"
              >
                <el-select
                  v-model="formData.select_freqresp"
                  class="full-width-input"
                >
                  <el-option
                    v-for="(item, idx) in select_freqrespOptions"
                    :key="idx"
                    :label="item.label"
                    :value="item.value"
                    :disabled="item.disabled"
                  ></el-option>
                </el-select>
              </el-form-item>
              <el-form-item
                label="功能投退"
                prop="select_onoff"
                class="required"
              >
                <el-select
                  v-model="formData.select_onoff"
                  class="full-width-input"
                  clearable
                >
                  <el-option
                    v-for="(item, idx) in select_onoffOptions"
                    :key="idx"
                    :label="item.label"
                    :value="item.value"
                    :disabled="item.disabled"
                  ></el-option>
                </el-select>
              </el-form-item>
              <div class="static-content-item" v-if="!isNarrow">
                <el-button type="primary" @click="submitForm">保存</el-button>
              </div>
            </div>
            <div v-else-if="card.content === 'second'">
              <el-form-item prop="input_pa">
                <div class="flex gap-4">
                  <span class="span">P<sub>A</sub></span>
                  <el-input
                    v-model="formData.input_pa"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="最大允许放电有功功率" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_pb" class="required">
                <div class="flex gap-4">
                  <span class="span">P<sub>B</sub></span>
                  <el-input
                    v-model="formData.input_pb"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="最大允许充电有功功率" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_fn" class="required">
                <div class="flex gap-4">
                  <span class="span">f<sub>N</sub></span>
                  <el-input
                    v-model="formData.input_fn"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="并网点额定频率" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_fd" class="required">
                <div class="flex gap-4">
                  <span class="span">f<sub>d</sub></span>
                  <el-input
                    v-model="formData.input_fd"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="调频死区" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_derta" class="required">
                <div class="flex gap-4">
                  <span class="span">δ%</span>
                  <el-input
                    v-model="formData.input_derta"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip
                    content="频率调差系数(δ%=-△f(%)/△P(%))"
                    placement="top"
                  >
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_a1" class="required">
                <div class="flex gap-4">
                  <span class="span">α<sub>1</sub></span>
                  <el-input
                    v-model="formData.input_a1"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="有功功率限幅系数下限" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_a2" class="required">
                <div class="flex gap-4">
                  <span class="span">α<sub>2</sub></span>
                  <el-input
                    v-model="formData.input_a2"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="有功功率限幅系数上限" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>
            </div>
            <div v-else-if="card.content === 'image'">
              <div class="svg-container">
                <img
                  src="@/assets/ems/pfc.svg"
                  alt="SVG Image"
                  class="responsive-svg"
                />
              </div>
            </div>
          </el-card>
        </transition-group>
      </el-col>
    </el-row>
    <!-- 保存按钮 -->
    <div class="save-button-container" v-if="isNarrow">
      <el-button type="primary" @click="submitForm">保存</el-button>
    </div>
  </el-form>
</template>

<script>
import {
  ref,
  computed,
  watch,
  nextTick,
  onMounted,
  onUnmounted,
  defineComponent,
  toRefs,
  reactive,
  getCurrentInstance,
} from "vue";
import {
  ElForm,
  ElFormItem,
  ElSelect,
  ElOption,
  ElInput,
  ElTooltip,
  ElIcon,
  ElButton,
  ElCard,
  ElRow,
  ElCol,
  ElMessage,
} from "element-plus";
import { InfoFilled } from "@element-plus/icons-vue";

import {
  GetConsulKey,
  SetConsulKey,
  SetJsonfile,
  GetJsonfile,
} from "@/api/showEms/sysOtherApi";

export default defineComponent({
  components: {
    ElForm,
    ElFormItem,
    ElSelect,
    ElOption,
    ElInput,
    ElTooltip,
    ElIcon,
    ElButton,
    ElCard,
    ElRow,
    ElCol,
    InfoFilled,
  },
  props: {},
  setup() {
    const state = reactive({
      consul: {
        key: "",
        value: "",
      },
      formData: {
        select_freqresp: "",
        select_onoff: "",
        input_pa: "",
        input_pb: "",
        input_fn: "",
        input_fd: "",
        input_derta: "",
        input_a1: "",
        input_a2: "",
      },
      rules: {
        select_freqresp: [{ required: true, message: "字段值不可为空" }],
        select_onoff: [{ required: true, message: "字段值不可为空" }],
        input_pa: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_pb: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_fn: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_fd: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_derta: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_a1: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_a2: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
      },
      select_freqrespOptions: [{ label: "一次调频", value: 1 }],
      select_onoffOptions: [
        { label: "投入", value: 1 },
        { label: "退出", value: 2 },
      ],
    });

    const isNarrow = computed(() => window.innerWidth < 768);

    const instance = getCurrentInstance();

    const submitForm = () => {
      instance.ctx.$refs["vForm"].validate(async (valid) => {
        if (valid) {
          // TODO: 提交表单
   
          state.consul.key = "select_freqresp";
          state.consul.value = state.formData.select_freqresp.toString();
          const response_freq = await SetConsulKey(JSON.stringify(state.consul));
          if (response_freq.code === 0) {
 
           } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "select_onoff";
          state.consul.value = state.formData.select_onoff.toString();
          const response_onoff = await SetConsulKey(JSON.stringify(state.consul));
          if (response_onoff.code === 0) {
 
           } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }


          state.consul.key = "input_pa";
          state.consul.value = state.formData.input_pa;
          const response_pa = await SetConsulKey(JSON.stringify(state.consul));
          if (response_pa.code === 0) {
 
           } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_pb";
          state.consul.value = state.formData.input_pb;
          const response_pb = await SetConsulKey(JSON.stringify(state.consul));
          if (response_pb.code === 0) {
             
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_fn"
          state.consul.value = state.formData.input_fn;
          const response_fn = await SetConsulKey(JSON.stringify(state.consul));
          if (response_fn.code === 0) {
             
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_fd"
          state.consul.value = state.formData.input_fd;
          const response_fd = await SetConsulKey(JSON.stringify(state.consul));
          if (response_fd.code === 0) {
             
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_derta"
          state.consul.value = state.formData.input_derta;
          const response_derta = await SetConsulKey(JSON.stringify(state.consul));
          if (response_derta.code === 0) {
             
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_a1"
          state.consul.value = state.formData.input_a1;
          const response_a1 = await SetConsulKey(JSON.stringify(state.consul));
          if (response_a1.code === 0) {
             
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_a2"
          state.consul.value = state.formData.input_a2;
          const response_a2 = await SetConsulKey(JSON.stringify(state.consul));
          if (response_a2.code === 0) {
             
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message:   "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          ElMessage({
                type: "success",
                message:   "保存成功！"
              });

          initFormData();
        }
      });
    };

    async function initFormData() {
      try {
        // 从consul服务器获取数据

        const response_freq = await GetConsulKey({ KEY: "select_freqresp" });

        if (response_freq.code === 0) {
          state.formData.select_freqresp = parseInt(response_freq.data, 10);
        } else {
          state.formData.select_freqresp = "";
          return;
        }

        const response_onoff= await GetConsulKey({ KEY: "select_onoff" });

        if (response_onoff.code === 0) {
          state.formData.select_onoff = parseInt(response_onoff.data, 10);
        } else {
          state.formData.select_onoff = "";
        }

        const response_pa = await GetConsulKey({ KEY: "input_pa" });
        if (response_pa.code === 0) {
          state.formData.input_pa = response_pa.data;
        } else {
          state.formData.input_pa = "";
        }

        const response_pb = await GetConsulKey({ KEY: "input_pb" });
        if (response_pb.code === 0) {
          state.formData.input_pb = response_pb.data;
        } else {
          state.formData.input_pb = "";
        }

       const response_fn = await GetConsulKey({ KEY: "input_fn" });
        if (response_fn.code === 0) {
          state.formData.input_fn = response_fn.data;
        } else {
          state.formData.input_fn = "";
        }

        const response_fd = await GetConsulKey({ KEY: "input_fd" });
        if (response_fd.code === 0) {
          state.formData.input_fd = response_fd.data;
        } else {
          state.formData.input_fd = "";
        }

        const response_derta = await GetConsulKey({ KEY: "input_derta" });
        if (response_derta.code === 0) {
          state.formData.input_derta = response_derta.data;
        } else {
          state.formData.input_derta = "";
        }

        const response_a1 = await GetConsulKey({ KEY: "input_a1" });
        if (response_derta.code === 0) {
          state.formData.input_a1 = response_a1.data;
        } else {
          state.formData.input_a1 = "";
        }

        const response_a2 = await GetConsulKey({ KEY: "input_a2" });
        if (response_a2.code === 0) {
          state.formData.input_a2 = response_a2.data;
        } else {
          state.formData.input_a2 = "";
        }

       } catch (error) {
         // ElMessage({
         //   type: "error",
         //   message: "获取consul失败",
         // });
       }
    }

    const handleResize = () => {
      state.isNarrow = window.innerWidth < 768;
    };

    onMounted(() => {
      window.addEventListener("resize", handleResize);
      handleResize(); // 初始化时检查窗口宽度
      initFormData();
    });

    onUnmounted(() => {
      window.removeEventListener("resize", handleResize);
    });

    const cards = [
      { title: "功能", content: "first" },
      { title: "图像", content: "image" },
      { title: "参数", content: "second" },
    ];

    return {
      ...toRefs(state),
      isNarrow,
      submitForm,
      cards,
    };
  },
});
</script>

<style lang="scss" scoped>
.card-container {
  display: flex;
  flex-wrap: wrap;
  margin-top: 20px;
}

.card-group {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
}

.form-card {
  margin-bottom: 20px;
  flex: 1 1 calc(33.333% - 10px); // 3列布局，每列宽度减去间距
  transition: all 0.3s ease;

  &.full-width {
    flex: 1 1 100%;
  }

  &.no-border {
    border: none;
    box-shadow: none;
  }

  .card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .el-input.full-width-input {
    width: 100%;
  }

  .el-icon {
    cursor: pointer;
    font-size: 18px;
  }

  .static-content-item {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 20px;
  }

  .svg-container {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 20px;
    width: 100%;
    height: 100%;
  }

  .responsive-svg {
    max-width: 100%;
    height: auto;
  }

  .el-button {
    margin-top: 20px;
  }
}

.save-button-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.card-transition-enter-active,
.card-transition-leave-active {
  transition: all 0.3s ease;
}

.card-transition-enter-from,
.card-transition-leave-to {
  opacity: 0;
  transform: translateY(30px);
}

.span {
  display: inline-block; /* 将 span 转换为行内块级元素 */
  width: 20px; /* 设置宽度 */
}

@media (max-width: 768px) {
  .form-card {
    flex: 1 1 100%;
    border: none;
    box-shadow: none;
  }
}
</style>
