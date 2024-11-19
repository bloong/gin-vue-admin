<template>
  <div class="dashboard">
    <el-card class="box-card" v-for="(item, index) in chartsData" :key="index">
      <div slot="header" class="clearfix">
        <span>{{ item.title }}</span>
      </div>
      <div :ref="item.ref" :id="item.id" style="width: 100%; height: 400px;"></div>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>今日充电量与放电量</span>
      </div>
      <p>今日充电量: {{ todayCharge }} kWh</p>
      <p>今日放电量: {{ todayDischarge }} kWh</p>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>累计充电量与放电量</span>
      </div>
      <p>累计充电量: {{ totalCharge }} kWh</p>
      <p>累计放电量: {{ totalDischarge }} kWh</p>
    </el-card>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>实时有功与无功</span>
      </div>
      <p>实时有功: {{ realPower }} kW</p>
      <p>实时无功: {{ reactivePower }} kVar</p>
    </el-card>
  </div>
</template>

<script>
import * as echarts from 'echarts';
import { ref, onMounted } from 'vue';

export default {
  name: 'Dashboard',
  setup() {
    const chartsData = [
      { title: '近7日充放电电量', id: 'chart1', ref: 'chart1' },
      { title: '今日SOC曲线', id: 'chart2', ref: 'chart2' },
      { title: '今日有功曲线', id: 'chart3', ref: 'chart3' },
      { title: '今日无功曲线', id: 'chart4', ref: 'chart4' }
    ];

    const todayCharge = ref(10); // 示例数据
    const todayDischarge = ref(5); // 示例数据
    const totalCharge = ref(1000); // 示例数据
    const totalDischarge = ref(800); // 示例数据
    const realPower = ref(20); // 示例数据
    const reactivePower = ref(10); // 示例数据

    onMounted(() => {
      chartsData.forEach(item => {
        let chart = echarts.init(document.getElementById(item.id));
        // 每个图表的配置项可以根据需要调整
        let option = {
          xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
          },
          yAxis: {
            type: 'value'
          },
          series: [{
            data: [120, 200, 150, 80, 70, 110, 130],
            type: 'bar'
          }]
        };
        chart.setOption(option);
      });
    });

    return {
      chartsData,
      todayCharge,
      todayDischarge,
      totalCharge,
      totalDischarge,
      realPower,
      reactivePower
    };
  }
};
</script>

<style scoped>
.dashboard .box-card {
  margin-bottom: 20px;
}
</style>