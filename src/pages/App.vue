<template>
  <context-menu ref="menu">
    <context-menu-item icon="fa-solid fa-pencil" text="Rename" type="classic" tooltip="Renomez ou changez l'icone de votre liste"/>
    <context-menu-item icon="fa-solid fa-clone" text="Duplicate" type="classic" tooltip="Maintenant vous en avez deux!"/>
    <context-menu-item icon="fa-solid fa-share-nodes" text="Share" type="disabled" tooltip="Coming soon"/>
    <context-menu-item icon="fa-solid fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre liste de manière définitive"/>
  </context-menu>

  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :fullscreen="fullscreen " :collapsed="sidebarCollapsed" :lists="lists" @openContextMenu="evtData => this.$refs.menu.open(evtData)"/>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";
import Topbar from "../components/Topbar.vue";
import api from "../api.js";
import ContextMenu from "../components/ContextMenu.vue";
import ContextMenuItem from "../components/ContextMenuItem.vue";

export default {
  name: "App",
  methods: {
    toggleSidebar() {
      this.sidebarCollapsed = !this.sidebarCollapsed;
    },
    sidebarDefaultValue() {
      //check if the device has medium height
      return window.innerWidth < 1024;
    },
    isFullscreen() {
      return window.innerWidth < 	1024;
    },
    onResize() {
      //handle resize operation
      this.fullscreen = this.isFullscreen();
    },
  },
  mounted() {
    //get the lists from the api
    api.getLists().then((lists) => {
      //complete the values of the params and check if they're valid
      const listId = lists.find(i => i.id === this.$route.params.listId) ? this.$route.params.listId : lists[0].id;
      const mode = ["view", "edit"].includes(this.$route.params.mode) ? this.$route.params.mode : "view";

      //redirect to the app page
      this.$router.push({ name: "app", params: { listId, mode } });

      //populate the list
      this.lists = lists;
    }).catch(statusCode => {
      //check if the status code is 401
      if (statusCode === 401) {
        //redirect to the login page
        this.$router.push({name: "login", query: {listId: this.$route.params.listId, mode: this.$route.params.mode}});
      }
    });
  },
  created() {
    //handle screen resizing
    window.addEventListener('resize', this.onResize);
  },
  data() {
    return {
      sidebarCollapsed: this.sidebarDefaultValue(),
      lists: [],
      fullscreen: this.isFullscreen(),
    };
  },
  components: {ContextMenu, Topbar, Sidebar, ContextMenuItem}
}
</script>

<style scoped>

</style>