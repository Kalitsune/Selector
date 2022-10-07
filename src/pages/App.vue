<template>
  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :fullscreen="fullscreen " :collapsed="sidebarCollapsed" :lists="lists"/>
</template>

<script>
import Sidebar from "../components/Sidebar.vue";
import Topbar from "../components/Topbar.vue";
import api from "../api.js";

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
    getLists() {
      //get the lists from the api
      api.getLists().then((lists) => {
        //complete the values of the params and check if they're valid
        const listId = lists.find(i => i.id === this.$route.params.listId) ? this.$route.params.listId : lists[0].id;
        const mode = ["view", "edit", "delete"].includes(this.$route.params.mode) ? this.$route.params.mode : "view";

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
    deleteList(listId) {
      //remove the 'delete' param to avoid deleting the wrong list by accident
      this.$router.push({ name: "app", params: { mode: "view", listId: this.$route.params.listId } });

      //delete the list
        api.deleteList(listId).then(() => {
          //remove the deleted list from the lists array
          this.lists = this.lists.filter(i => i.id !== listId);

          //redirect to the first list of the array
          this.$router.push({ name: "app", params: { listId: this.lists[0].id, mode: "view" } });
        }).catch(statusCode => {
          //check if the status code is 401
          if (statusCode === 401) {
            //redirect to the login page
            this.$router.push({name: "login", query: {listId: this.$route.params.listId, mode: "delete"}});

          } else if (statusCode === 404) {
            //redirect to the app page
            this.$router.push({ name: "app", params: { listId: this.lists[0].id, mode: "view" } });
          }
      });
    },
  },
  async mounted() {
    this.getLists();
  },
  created() {
    //handle screen resizing
    window.addEventListener('resize', this.onResize);

    this.$watch(
        () => this.$route.params,
        (toParams, previousParams) => {
          if (toParams.mode === "delete") {
            this.deleteList(toParams.listId);
          }
        }
    )
  },
  data() {
    return {
      sidebarCollapsed: this.sidebarDefaultValue(),
      lists: [],
      fullscreen: this.isFullscreen(),
    };
  },
  components: { Topbar, Sidebar}
}
</script>

<style scoped>

</style>