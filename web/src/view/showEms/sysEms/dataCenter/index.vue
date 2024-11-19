<template>
  <div>
    <div class="data-center-box">
      <CenterCard title="在线情况" class="full-width-on-narrow">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ChainRatio />
        </template>
      </CenterCard>
      <CenterCard title="运行功率" class="full-width-on-narrow">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <ReclaimMileage />
        </template>
      </CenterCard>
      <CenterCard title="设置功率" class="full-width-on-narrow">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <RecoveryRate />
        </template>
      </CenterCard>
      <CenterCard title="电网电压" class="full-width-on-narrow">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <Order />
        </template>
      </CenterCard>
      <CenterCard title="功率曲线" class="full-width-on-narrow" style="grid-column-start: span 3;">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <lineCharts />
        </template>
      </CenterCard>
      <CenterCard title="源接入台数占比" class="full-width-on-narrow" style="grid-auto-columns: 1.5fr">
        <template #action>
          <span class="gvaIcon-prompt" style="color: #999" />
        </template>
        <template #body>
          <part />
        </template>
      </CenterCard>
    </div>
  </div>
</template>

<script setup>
import CenterCard from './dataCenterComponents/centerCard.vue'
import ChainRatio from './dataCenterComponents/chainRatio.vue'
import ReclaimMileage from './dataCenterComponents/ReclaimMileage.vue'
import RecoveryRate from './dataCenterComponents/RecoveryRate.vue'
import Order from './dataCenterComponents/order.vue'
import lineCharts from './dataCenterComponents/lineCharts.vue'
import part from './dataCenterComponents/part.vue'
</script>

<script>
export default {
  name: 'DataCenter',
}
</script>

<style lang="scss" scoped>
.data-center-box {
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1.5fr 1.5fr 1.5fr; // 默认布局
  column-gap: 10px;
  row-gap: 10px; // 增加行间距以保持美观

  @media (max-width: 768px) { // 当屏幕宽度小于768px时应用以下样式
    grid-template-columns: 1fr; // 将所有列合并为单列
  }

  .full-width-on-narrow {
    @media (max-width: 768px) { // 当屏幕宽度小于768px时应用以下样式
      grid-column: 1 / -1; // 占据整个行
    }
  }

  .center-card {
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    background-color: #fff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
}
</style>