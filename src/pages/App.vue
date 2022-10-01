<template>
  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :collapsed="sidebarCollapsed" :lists="lists"/>
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
  },
  mounted() {
    //get the lists from the api
    api.getLists().then((lists) => {
      //check if there's a default route selected
      if (!this.$route.params.hasOwnProperty('listId')) {
        console.log(lists[0].id);
        //set the list to the default
        this.$router.push({name: "app", params: {listId: lists[0].id, mode: "view"}});
      }

      //populate the list
      this.lists = lists;
    }).catch(statusCode => {
      //check if the status code is 401
      if (statusCode === 401) {
        //redirect to the login page
        this.$router.push({name: "login"});
      }
    });
  },
  data() {
    return {
      sidebarCollapsed: this.sidebarDefaultValue(),
      lists: [],
    };
  },
  components: {Topbar, Sidebar }
}
</script>

<style scoped>

</style>