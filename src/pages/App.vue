<template>
  <login v-if="needLogin" :handler="loggedIn" expired class="backdrop-blur-lg" />

  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :collapsed="sidebarCollapsed"/>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";
import Topbar from "../components/Topbar.vue";
import Login from "../components/Login.vue";

export default {
  name: "App",
  methods: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed;
    },
    loggedIn() {
      this.$store.commit('needLogin', false)
    }
  },
  mounted() {
    this.$api.refreshLists();
  },
  data() {
    return {
      sidebarCollapsed: !this.$store.state.isMobile,
    };
  },
  computed: {
    needLogin() {
      return this.$store.state.needLogin;
    },
  },
  components: {Login, Topbar, Sidebar}
}
</script>

<style scoped>

</style>