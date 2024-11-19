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
    <div class="card-container">
      <el-card class="form-card" shadow="never">
        <template #header>
          <div class="card-header">
            <span>有功控制</span>
          </div>
        </template>

        <el-form-item label="功能投退" prop="select_onoff">
          <el-select
            v-model="formData.select_onoff"
            class="full-width-input"
            clearable
          >
            <el-option
              v-for="(item, index) in select_onoff_Options"
              :key="index"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="本地/远程" prop="select_local_ctrl">
          <el-select
            v-model="formData.select_local_ctrl"
            class="full-width-input"
            clearable
            @change="handleLocalCtrlChange"
          >
            <el-option
              v-for="(item, index) in select_local_ctrl_Options"
              :key="index"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item prop="input_active_power_r">
          <div class="flex gap-4">
            <span class="span">P<sub>R</sub></span>
            <el-input
              v-model="formData.input_active_power_r"
              type="text"
              :minlength="1"
              :maxlength="7"
              :disabled="formData.select_local_ctrl !== 2"
            />
            <el-tooltip content="远程设置有功功率(MW)" placement="top">
              <el-icon><InfoFilled /></el-icon>
            </el-tooltip>
          </div>
        </el-form-item>

        <el-form-item prop="input_active_power_L">
          <div class="flex gap-4">
            <span class="span">P<sub>L</sub></span>
            <el-input
              v-model="formData.input_active_power_L"
              type="text"
              :minlength="1"
              :maxlength="7"
              :disabled="formData.select_local_ctrl !== 1"
              s
            />
            <el-tooltip content="本地设置有功功率(MW)" placement="top">
              <el-icon><InfoFilled /></el-icon>
            </el-tooltip>
          </div>
        </el-form-item>

        <div class="save-button-container">
          <el-button type="primary" @click="submitForm">保存</el-button>
        </div>
      </el-card>
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
  components: {},
  props: {},
  setup() {
    const state = reactive({
      consul: {
        key: "",
        value: "",
      },
      formData: {
        select_inertia: "",
        input_fd: "",
        input_derta: "",
        input_a: "",
        input_t: "",
      },
      rules: {},

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

      input_a: [
        { required: true, message: "字段值不可为空" },
        {
          pattern: /^\d+(\.\d+)?$/,
          trigger: ["blur", "change"],
          message: "请输入有效的数字",
        },
      ],

      input_t: [
        { required: true, message: "字段值不可为空" },
        {
          pattern: /^\d+(\.\d+)?$/,
          trigger: ["blur", "change"],
          message: "请输入有效的数字",
        },
      ],

      select_onoff_Options: [
        {
          label: "投入",
          value: 1,
        },
        {
          label: "退出",
          value: 2,
        },
      ],

      select_local_ctrl_Options: [
        { label: "本地", value: 1 },
        { label: "远程", value: 2 },
      ],
    });

    const isNarrow = computed(() => window.innerWidth < 768);

    const instance = getCurrentInstance();

    const submitForm = () => {
      instance.ctx.$refs["vForm"].validate(async (valid) => {
        if (valid) {
          // TODO: 提交表单

          state.consul.key = "select_inertia";
          state.consul.value = state.formData.select_inertia.toString();
          const response_inertia = await SetConsulKey(
            JSON.stringify(state.consul)
          );
          if (response_inertia.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_derta";
          state.consul.value = state.formData.input_derta;
          const response_input_derta = await SetConsulKey(
            JSON.stringify(state.consul)
          );
          if (response_input_derta.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_fd";
          state.consul.value = state.formData.input_fd;
          const response_input_fd = await SetConsulKey(
            JSON.stringify(state.consul)
          );
          if (response_input_fd.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_a";
          state.consul.value = state.formData.input_a;
          const response_a = await SetConsulKey(JSON.stringify(state.consul));
          if (response_a.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "input_t";
          state.consul.value = state.formData.input_t;
          const response_t = await SetConsulKey(JSON.stringify(state.consul));
          if (response_t.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          ElMessage({
            type: "success",
            message: "保存成功！",
          });

          initFormData();
        }
      });
    };

    async function initFormData() {
      try {
        // 从consul服务器获取数据

        const response_inertia = await GetConsulKey({ KEY: "select_inertia" });

        if (response_inertia.code === 0) {
          state.formData.select_inertia = parseInt(response_inertia.data, 10);
        } else {
          state.formData.select_inertia = "";
          return;
        }

        const response_input_fd = await GetConsulKey({ KEY: "input_fd" });

        if (response_input_fd.code === 0) {
          state.formData.input_fd = response_input_fd.data;
        } else {
          state.formData.input_fd = "";
        }

        const response_derta = await GetConsulKey({ KEY: "input_derta" });

        if (response_derta.code === 0) {
          state.formData.input_derta = response_derta.data;
        } else {
          state.formData.input_derta = "";
        }

        const response_a = await GetConsulKey({ KEY: "input_a" });
        if (response_a.code === 0) {
          state.formData.input_a = response_a.data;
        } else {
          state.formData.input_a = "";
        }

        const response_t = await GetConsulKey({ KEY: "input_t" });
        if (response_t.code === 0) {
          state.formData.input_t = response_t.data;
        } else {
          state.formData.input_t = "";
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

    return {
      ...toRefs(state),
      isNarrow,
      submitForm,
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

.form-card {
  margin-bottom: 20px;
  transition: all 0.3s ease;

  &.no-border {
    border: none;
    box-shadow: none;
  }

  .card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
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
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 8px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 28px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}

div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;

    td.table-cell {
      display: table-cell;
      height: 36px;
      border: 1px solid #e1e2e3;
    }
  }
}

div.tab-container {
}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.custom-label {
}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}

.save-button-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.span {
  display: inline-block; /* 将 span 转换为行内块级元素 */
  width: 20px; /* 设置宽度 */
  text-align: left;
}

.custom-input.is-readonly {
  background-color: #f5f7fa; /* 更改背景颜色为浅灰色 */
}

@media (max-width: 768px) {
  .form-card {
    flex: 1 1 100%;
    border: none;
    box-shadow: none;
  }
}
</style>
