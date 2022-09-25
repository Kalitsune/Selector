<template>
  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :collapsed="sidebarCollapsed" />
</template>

<script>
import Sidebar from "../components/sidebar.vue";
import Topbar from "../components/topbar.vue";

export default {
  name: "App",
  methods: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed;
    },
    sidebarDefaultValue() {
      //check if the device has medium height
      return window.innerWidth > 768;
    },
  },
  data() {
    return {
      sidebarCollapsed: this.sidebarDefaultValue()
    };
  },
  components: {Topbar, Sidebar }
}
</script>

<style scoped>

</style>