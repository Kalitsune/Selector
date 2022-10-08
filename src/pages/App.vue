<template>
  <div class="h-screen w-screen bg-neutral-200 dark:bg-neutral-800 overflow-hidden">
    <login v-if="needLogin" :handler="loggedIn" expired class="backdrop-blur" />

    <Topbar @toggleSidebar="toggleSidebar"/>
    <Sidebar :collapsed="sidebarCollapsed"/>
  </div>
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
    },
    isMobile() {
      return window.innerWidth < 	1024;
    },
    updateIsMobile() {
      //handle resize operation
      this.$store.commit("setIsMobile", this.isMobile());
    }
  },
  mounted() {
    this.$api.refreshLists();
  },
  data() {
    return {
      sidebarCollapsed: this.isMobile(),
    };
  },
  created() {
    //handle screen resizing
    window.addEventListener('resize', this.updateIsMobile);

    //set the default value for the isMobile property
    this.updateIsMobile();
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