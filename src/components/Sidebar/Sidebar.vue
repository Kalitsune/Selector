<template>
  <!-- sidebar -->
  <div :class="{'collapsed': collapsed, 'isMobile': isMobile}" class="sidebar" @contextmenu="openContextMenu" >
    <!-- sidebar items -->
    <ul>
      <!-- array of lists -->
      <li v-if="lists.length > 0" v-for="list in lists">
        <SidebarItem :list="list" :disabled="isDisabled(list)" :isSelected="isSelected(list)" @closeContextMenu="this.$root.$refs.ContextMenu.close()" @openContextMenu="evtData => this.$root.$refs.ContextMenu.open(evtData)"/>
      </li>

      <!-- if there's no list -->
      <li v-else>
        <SidebarItem :list="{ name: 'There\'s no list to show!', id: 0}" disabled/>
      </li>
    </ul>
  </div>

</template>

<script>

import SidebarItem from "./SidebarItem.vue";
import ContextMenu from "../contextMenu/ContextMenu.vue";
import ContextMenuItem from "../contextMenu/ContextMenuItem.vue";

export default {
  name: "Sidebar",
  components: {SidebarItem, ContextMenu, ContextMenuItem},
  methods: {
    isSelected(list) {
      //check if the provided element is currently selected
      return this.$route.params.listId === list.id;
    },
    isDisabled(list) {
      //check if the provided element is currently disabled
      return list.id === 0;
    },
    openContextMenu(event) {
      //get the coordinates of the click
      let x = event.pageX || event.clientX;
      let y = event.pageY || event.clientY;

      //prevent the default context menu and prevent triggering other events
      event.preventDefault();
      event.stopPropagation();

      //open the context menu
      this.$root.$refs.ContextMenu.open({x, y, menu: "sidebar"});
    }
  },
  computed: {
      lists() {
        //make the lists available to the template
        return this.$store.state.lists;
      },
      isMobile() {
        //check if the sidebar is isMobile
        return this.$store.state.isMobile;
      }
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    },
  },
};
</script>

<style scoped>
.sidebar {
  @apply left-0 top-0 h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll pb-6
}

.sidebar:not(.isMobile) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.isMobile) {
  @apply -translate-x-full;
}

.sidebar.isMobile {
  @apply w-screen bottom-0 bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.isMobile {
  @apply translate-y-full opacity-0;
}
</style>