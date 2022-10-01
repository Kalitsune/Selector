<template>
  <Topbar @toggleSidebar="toggleSidebar"/>
  <Sidebar :collapsed="sidebarCollapsed" :lists="lists" :selected="selected" @switchList="switchList"/>
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
    switchList(list) {
      // change the selected list
      this.selected = list;
    },
  },
  mounted() {
    //get the lists from the api
    api.getLists().then((lists) => {
      this.selected = lists[0]
      this.lists = lists;
    });
  },
  data() {
    return {
      sidebarCollapsed: this.sidebarDefaultValue(),
      lists: [],
      selected: {},
    };
  },
  components: {Topbar, Sidebar }
}
</script>

<style scoped>

</style>