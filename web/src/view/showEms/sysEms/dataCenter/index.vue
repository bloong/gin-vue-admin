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
import { useRosClient } from '@/utils/useRosClient';
import { ref , computed, watch, nextTick, onMounted ,onUnmounted, defineComponent, toRefs, reactive, getCurrentInstance } from "vue";


</script>

<script>
export default {
  name: 'DataCenter',
}


async function initFormData() {
  try {

   const websocket = ref(null);
 
   const { isConnected, messages, ready, subscribeToTopic , publishToTopic }  = useRosClient('ws://172.21.27.24:9090');

    function publishMessage() {
      const topicName = '/chatter';
      const messageType = 'std_msgs/String';
      const messageData = { data: 'Hello ROS 2!' };
      publishToTopic(topicName, messageType, messageData);
    }

    // 确保 messages 始终是一个数组
    if (!Array.isArray(messages.value)) {
      console.warn('messages is not an array, initializing as empty array');
      messages.value = [];
    }

    watch(ready, (newReady) => {
       console.log('Ready state changed:', newReady);
       if (newReady) {
         // 订阅一个话题
         const topicName = '/chatter';
         const messageType = 'std_msgs/String';
         const topic = subscribeToTopic(topicName, messageType);
         console.log('Subscribed to topic:', topicName);

         topic.subscribe((message) => {           
          //  state.formData.input12931 = String(message.data);
              // ElMessage({
              //     type: "info",
              //     message: String(message.data),
              //   });
           });
       }
     });

        // 其他字段也可以类似处理
      } catch (error) {
        // ElMessage({
        //   type: "error",
        //   message: "获取consul失败",
        // });
      }
    }

// // Function to initialize the WebSocket connection
const initWebSocket = () => {
  websocket.value = new WebSocket('ws://172.21.27.24:7681') // Replace with your WebSocket URL

  websocket.value.onopen = () => {
            //  ElMessage({
            //       type: "success",
            //       message: 'WebSocket connection opened',
            //     });   
  }
  websocket.value.onerror = (error) => {
                //  ElMessage({
            //       type: "success",
            //       message: 'WebSocket error',
            //     });  
  }

  websocket.value.onclose = () => {
                //  ElMessage({
            //       type: "success",
            //       message: 'WebSocket connection closed',
            //     });  
  }

  websocket.value.onmessage = async (event) => {
    const message = JSON.parse(event.data)
      //  ElMessage({
      //             type: "success",
      //             message: message.type
      //           });  
    switch (message.type){
      case "pong":
        break;
      case "message":
          // ElMessage({
          //         type: "success",
          //         message: message.data.content,
          //       });    

         state.formData.input12931 = message.data.content;
        break;
      case "offline":
        //更新用户在线状态
            ElMessage({
             type: "success",
             message: "offline !",
           });        
           break;

      case "online":
        //更新用户在线状态
             ElMessage({
             type: "success",
             message: "online !",
           });        
           break;
    }
  }
};

const sendPing = () => {
 
    setInterval(function () {
    if(websocket.value && websocket.value.readyState === WebSocket.OPEN){
       websocket.value.send(JSON.stringify({
          type: 'message',data: { content: Date.now(),}
        }))
     }
     publishMessage();  //发布ROS2消息
      
  }, 1000)
};


    // 在组件挂载时初始化数据
  onMounted(() => {
    initFormData();
    initWebSocket();
    sendPing();
  });

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