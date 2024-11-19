<template>
  <div id="app">
    <WebEmbed />
  </div>
</template>

<script>
import WebEmbed from './WebEmbed.vue'

export default {
  name: 'App',
  components: {
    WebEmbed
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin: 0; /* Remove default margin to ensure full screen */
  padding: 0; /* Remove default padding to ensure full screen */
  height: 100vh; /* Ensure app takes up the full height of the viewport */
}
</style>