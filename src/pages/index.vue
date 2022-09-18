<template>
  <div class="-z-50 fixed bg-neutral-100 dark:bg-neutral-800 w-screen h-screen">
    <App v-if="this.$cookies.isKey('token')"/>
    <Welcome v-else/>
<!--    {{ this.$cookies.set('token', "test") }}-->
  </div>
</template>

<script>
import App from './App.vue'
import Welcome from './Welcome.vue'

export default {
  name: "index",
  components: { App, Welcome },
}
</script>
