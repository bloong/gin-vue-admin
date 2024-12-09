//useRosClient.js
import { ref, onMounted, onUnmounted } from 'vue';
import * as ROSLIB from 'roslib';

export function useRosClient(url) {
  const ros = ref(null);
  const isConnected = ref(false);
  const messages = ref([]);
  const ready = ref(false);

  function initRosClient(callback) {
    ros.value = new ROSLIB.Ros({
      url: url,
      reconnect: true,
      retryInterval: 5000
    });

    ros.value.on('connection', () => {
      isConnected.value = true;
      console.log('Connected to ROS 2 server');
      ready.value = true;
      callback?.();
    });

    ros.value.on('error', (error) => {
      console.error('ROS 2 connection error:', error);
    });

    ros.value.on('close', () => {
      isConnected.value = false;
      console.log('Disconnected from ROS 2 server');
    });
  }

  function subscribeToTopic(topicName, messageType) {
    if (!ros.value) {
      console.error('ROS client is not initialized');
      return;
    }

    const topic = new ROSLIB.Topic({
      ros: ros.value,
      name: topicName,
      messageType: messageType
    });

    topic.subscribe((message) => {
       console.log('Received message on', topicName, ':', message);
    });

    return topic;
  }

  function publishToTopic(topicName, messageType, messageData) {
    if (!ros.value) {
      console.error('ROS client is not initialized');
      return;
    }

    const topic = new ROSLIB.Topic({
      ros: ros.value,
      name: topicName,
      messageType: messageType
    });

    const message = new ROSLIB.Message(messageData);
    topic.publish(message);
  }

  onMounted(() => {
    initRosClient(() => {
      console.log('ROS client initialized and connected');
    });
  });

  onUnmounted(() => {
    if (ros.value) {
      ros.value.close();
    }
  });

  return { isConnected, messages, ready, subscribeToTopic, publishToTopic };
}