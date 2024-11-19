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
                label="动态调压"
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
              <el-form-item prop="input_qa">
                <div class="flex gap-4">
                  <span class="span">Q<sub>A</sub></span>
                  <el-input
                    v-model="formData.input_qa"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="最大输出无功功率" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_qb" class="required">
                <div class="flex gap-4">
                  <span class="span">Q<sub>B</sub></span>
                  <el-input
                    v-model="formData.input_qb"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="最大吸收无功功率" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_uN" class="required">
                <div class="flex gap-4">
                  <span class="span">U<sub>N</sub></span>
                  <el-input
                    v-model="formData.input_uN"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="并网点额定电压" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_ud" class="required">
                <div class="flex gap-4">
                  <span class="span">U<sub>d</sub></span>
                  <el-input
                    v-model="formData.input_ud"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="调压死区" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_ud" class="required">
                <div class="flex gap-4">
                  <span class="span">K</span>
                  <el-input
                    v-model="formData.input_ud"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip
                    content="电压调差系数(K=-△Q(%)/△U(%))"
                    placement="top"
                  >
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_uL" class="required">
                <div class="flex gap-4">
                  <span class="span">β<sub>1</sub></span>
                  <el-input
                    v-model="formData.input_uL"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="无功功率限幅系数下限" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>

              <el-form-item prop="input_uH" class="required">
                <div class="flex gap-4">
                  <span class="span">β<sub>2</sub></span>
                  <el-input
                    v-model="formData.input_uH"
                    type="text"
                    :minlength="1"
                    :maxlength="7"
                    :disabled="formData.select_onoff !== 1"
                  />
                  <el-tooltip content="无功功率限幅系数上限" placement="top">
                    <el-icon><InfoFilled /></el-icon>
                  </el-tooltip>
                </div>
              </el-form-item>
            </div>
            <div v-else-if="card.content === 'image'">
              <div class="svg-container">
                <img
                  src="@/assets/ems/qfc.svg"
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
  defineComponent,
  toRefs,
  reactive,
  computed,
  onMounted,
  onUnmounted,
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
} from "element-plus";
import { InfoFilled } from "@element-plus/icons-vue";

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
      formData: {
        select_freqresp: "",
        select_onoff: "",
        input_qa: "",
        input_qb: "",
        input_q0: "",
        input_uN: "",
        input_ud: "",
        input_ul: "",
        input_uh: "",
      },
      rules: {
        select_freqresp: [{ required: true, message: "字段值不可为空" }],
        select_onoff: [{ required: true, message: "字段值不可为空" }],
        input_qa: [
          { required: true, message: "字段值不可为空" },
          {
            pattern: /^\d+(\.\d+)?$/,
            trigger: ["blur", "change"],
            message: "请输入有效的数字",
          },
        ],
        input_qb: [{ required: true, message: "字段值不可为空" }],
        input_q0: [{ required: true, message: "字段值不可为空" }],
        input_fN: [{ required: true, message: "字段值不可为空" }],
        input_fd: [{ required: true, message: "字段值不可为空" }],
        input_fL: [{ required: true, message: "字段值不可为空" }],
        input_fH: [{ required: true, message: "字段值不可为空" }],
      },
      select_freqrespOptions: [{ label: "动态调压", value: 1 }],
      select_onoffOptions: [
        { label: "投入", value: 1 },
        { label: "退出", value: 2 },
      ],
    });

    const isNarrow = computed(() => window.innerWidth < 768);

    const submitForm = () => {
      this.$refs["vForm"].validate((valid) => {
        if (valid) {
          // TODO: 提交表单
        }
      });
    };

    const resetForm = () => {
      this.$refs["vForm"].resetFields();
    };

    const handleResize = () => {
      state.isNarrow = window.innerWidth < 768;
    };

    onMounted(() => {
      window.addEventListener("resize", handleResize);
      handleResize(); // 初始化时检查窗口宽度
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
      resetForm,
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
  width: 30px; /* 设置宽度 */
}

@media (max-width: 768px) {
  .form-card {
    flex: 1 1 100%;
    border: none;
    box-shadow: none;
  }
}
</style>
