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
  },
  data() {
    return {
      sidebarCollapsed: false,
    };
  },
  components: {Topbar, Sidebar }
}
</script>

<style scoped>

</style>