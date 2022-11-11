<template>
  <!-- sidebar -->
  <div :class="{'collapsed': collapsed, 'isMobile': isMobile}" class="sidebar" @contextmenu="openContextMenu" >
    <!-- sidebar items -->
    <ul>
      <!-- array of lists -->
      <li v-if="lists.length > 0" v-for="list in lists">
        <SidebarItem :list="list" :disabled="isDisabled(list)" :isSelected="isSelected(list)" @openContextMenu="evtData => this.$root.$refs.ContextMenu.open(evtData)"/>
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
      let left = event.pageX || event.clientX;
      let top = event.pageY || event.clientY;

      //prevent the default context menu and prevent triggering other events
      event.preventDefault();
      event.stopPropagation();

      //open the context menu
      this.$root.$refs.ContextMenu.open({left, top, menu: "sidebar"});
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
  @apply h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll pb-6;
}

.sidebar:not(.isMobile) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.isMobile) {
  @apply -translate-x-full;
}

.sidebar.isMobile {
  @apply w-screen left-0 bottom-0 bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.isMobile {
  @apply translate-y-full opacity-0;
}
</style>