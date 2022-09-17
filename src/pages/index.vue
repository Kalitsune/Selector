<template>
  <App v-if="this.$cookies.isKey('token')"/>
  <Welcome v-else/>
</template>

<script>
import App from './App.vue'
import Welcome from './Welcome.vue'

export default {
  name: "index",
  components: { App, Welcome },
}
</script>
