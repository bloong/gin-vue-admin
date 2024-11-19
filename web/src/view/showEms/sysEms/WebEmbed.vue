<template>
  <div class="web-embed">
    <iframe
      :src="url"
      class="full-screen-iframe"
      frameborder="0"
      allowfullscreen
    ></iframe>
  </div>
</template>

<script>
export default {
  name: 'WebEmbed',
  data() {
    return {
        url: 'http://172.21.27.24:1881' // 这里填入你想嵌入的网页地址
        // url: 'http://10.30.0.6:3000/d/c98e8e8a-813a-4b34-9ce5-b6a553797e02/ems?from=now%2Fd&to=now'
    }
  }
}
</script>

<style scoped>
.web-embed {
  position: relative;
  width: 100%;
  height: 100vh; /* 100% of the viewport height */
  overflow: hidden;
}

.full-screen-iframe {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border: none; /* Remove iframe border */
  overflow: hidden;
}
</style>