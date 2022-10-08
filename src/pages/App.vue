<template>
  <popup-login href="popupLogin"/>

  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :collapsed="sidebarCollapsed"/>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";
import Topbar from "../components/Topbar.vue";
import PopupLogin from "../components/PopupLogin.vue";

export default {
  name: "App",
  methods: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed;
    },
  },
  mounted() {
    this.$api.getLists();
  },
  data() {
    return {
      sidebarCollapsed: !this.$store.state.isMobile,
    };
  },
  components: {PopupLogin, Topbar, Sidebar}
}
</script>

<style scoped>

</style>