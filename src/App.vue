<template>
  <div class="h-screen w-screen bg-neutral-200 dark:bg-neutral-800 overflow-hidden" @contextmenu="openContextMenu">

    <context-menu ref="ContextMenu"/>

    <login v-if="needLogin" :handler="loggedIn" class="backdrop-blur" />

    <Topbar @toggleSidebar="toggleSidebar"/>
    <Sidebar :collapsed="sidebarCollapsed"/>
  </div>
</template>

<script>
import Sidebar from "./components/Sidebar/Sidebar.vue";
import Topbar from "./components/Topbar.vue";
import Login from "./components/Login/Login.vue";
import ContextMenu from "./components/contextMenu/ContextMenu.vue";

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
    },
    openContextMenu(event) {
      //get the coordinates of the click
      let x = event.pageX || event.clientX;
      let y = event.pageY || event.clientY;

      //prevent the default context menu and prevent triggering other events
      event.preventDefault();
      event.stopPropagation();

      //open the context menu
      this.$refs.ContextMenu.open({x, y, menu: "blank"});
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
  components: {Login, Topbar, Sidebar, ContextMenu}
}
</script>

<style scoped>

</style>