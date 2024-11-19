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
            <span>备电功能</span>
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

        <el-form-item prop="input_soc_H">
          <div class="flex gap-8">
            <span class="span">SOC<sub>H</sub></span>
            <el-input
              v-model="formData.input_soc_H"
              type="text"
              :minlength="1"
              :maxlength="7"
              :disabled="formData.select_onoff !== 1"
            />
            <el-tooltip content="充电SOC上限(%)" placement="top">
              <el-icon><InfoFilled /></el-icon>
            </el-tooltip>
          </div>
        </el-form-item>

        <el-form-item prop="input_soc_L">
          <div class="flex gap-8">
            <span class="span">SOC<sub>L</sub></span>
            <el-input
              v-model="formData.input_soc_L"
              type="text"
              :minlength="1"
              :maxlength="7"
              :disabled="formData.select_onoff !== 1"
              s
            />
            <el-tooltip content="放电SOC下限(%)" placement="top">
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
        select_onoff: "",
        input_soc_H: "",
        input_soc_L: "",
      },
      rules: {},

      input_soc_H: [
        { required: true, message: "字段值不可为空" },
        {
          pattern: /^\d+(\.\d+)?$/,
          trigger: ["blur", "change"],
          message: "请输入有效的数字",
        },
      ],

      input_soc_L: [
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
    });

    const websocket = ref(null);

    const isNarrow = computed(() => window.innerWidth < 768);

    const instance = getCurrentInstance();

    const submitForm = () => {
      instance.ctx.$refs["vForm"].validate(async (valid) => {
        if (valid) {
          // TODO: 提交表单

          state.consul.key = "select_onoff";
          state.consul.value = state.formData.select_onoff.toString();
          const response_onoff = await SetConsulKey(
            JSON.stringify(state.consul)
          );
          if (response_onoff.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "common/soc_up_limit";
          state.consul.value = state.formData.input_soc_H;
          const response_input_soc_H = await SetConsulKey(
            JSON.stringify(state.consul)
          );
          if (response_input_soc_H.code === 0) {
          } else {
            initFormData();
            ElMessage({
              type: "error",
              message: "保存失败！(" + state.consul.key + ")",
            });
            return;
          }

          state.consul.key = "common/soc_dn_limit";
          state.consul.value = state.formData.input_soc_L;
          const response_input_soc_L = await SetConsulKey(
            JSON.stringify(state.consul)
          );
          if (response_input_soc_L.code === 0) {
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

          if (
            websocket.value &&
            websocket.value.readyState === WebSocket.OPEN
          ) {
            websocket.value.send(
              JSON.stringify({
                type: "ch_consul",
                data: { content: Date.now() },
              })
            );
          }

          initFormData();
        }
      });
    };

    async function initFormData() {
      try {
        // 从consul服务器获取数据

        // const response_select_onoff = await GetConsulKey({ KEY: "select_onoff" });

        // if (response_select_onoff.code === 0) {
        //   state.formData.select_onoff = parseInt(response_select_onoff.data, 10);
        // } else {
        //   state.formData.select_onoff = "";
        //   return;
        // }

        const response_input_soc_H = await GetConsulKey({ KEY: "common/soc_up_limit" });

        if (response_input_soc_H.code === 0) {
          state.formData.input_soc_H = response_input_soc_H.data;
        } else {
          state.formData.input_soc_H = "";
        }

        const response_input_soc_L = await GetConsulKey({ KEY: "common/soc_dn_limit" });

        if (response_input_soc_L.code === 0) {
          state.formData.input_soc_L = response_input_soc_L.data;
        } else {
          state.formData.input_soc_L = "";
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

    // // Function to initialize the WebSocket connection
    const initWebSocket = () => {
      websocket.value = new WebSocket("ws://192.168.0.216:7681"); // Replace with your WebSocket URL
      //  websocket.value = new WebSocket('ws://172.21.27.24:7681') // Replace with your WebSocket URL

      websocket.value.onopen = () => {
        //  ElMessage({
        //       type: "success",
        //       message: 'WebSocket connection opened',
        //     });
      };
      websocket.value.onerror = (error) => {
        //  ElMessage({
        //       type: "success",
        //       message: 'WebSocket error',
        //     });
      };

      websocket.value.onclose = () => {
        //  ElMessage({
        //       type: "success",
        //       message: 'WebSocket connection closed',
        //     });
      };

      websocket.value.onmessage = async (event) => {
        const message = JSON.parse(event.data);

        switch (message.type) {
          case "pong":
            break;
          case "run_inf":
            // ElMessage({
            //   type: "success",
            //   message: message.data.power_inf.charge_active_max_power,
            // });
            state.formData.input_soc_H =
              message.data.power_inf.charge_active_max_power;
            break;
        }
      };
    };

    const sendPing = () => {
      setInterval(function () {
        if (websocket.value && websocket.value.readyState === WebSocket.OPEN) {
          websocket.value.send(
            JSON.stringify({
              type: "run_inf",
              data: { content: Date.now() },
            })
          );

          //  websocket.value.send("run_inf");
        }
      }, 500);
    };

    onMounted(() => {
      window.addEventListener("resize", handleResize);
      handleResize(); // 初始化时检查窗口宽度
      initFormData();
      initWebSocket();
      sendPing();
    });

    onUnmounted(() => {
      if (websocket.value) {
        websocket.value.close();
      }

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
