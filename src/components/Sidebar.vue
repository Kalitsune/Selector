<template>
  <div :class="{'collapsed': collapsed, 'fullscreen': fullscreen}" class="sidebar">
    <ul>
      <li v-if="lists.length > 0" v-for="list in lists">
        <SidebarItem :list="list" :disabled="isDisabled(list)" :isSelected="isSelected(list)" @openContextMenu="evtData => this.$emit('openContextMenu', evtData)"/>
      </li>
      <li v-else>
        <SidebarItem :list="{ name: 'There\'s no list to show!', id: 0}" disabled/>
      </li>
    </ul>
  </div>

</template>

<script>

import SidebarItem from "./SidebarItem.vue";
export default {
  name: "Sidebar",
  components: {SidebarItem},
  methods: {
    isSelected(list) {
      //check if the provided element is currently selected
      return this.$route.params.listId === list.id;
    },
    isDisabled(list) {
      //check if the provided element is currently disabled
      return list.id === 0;
    },
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    },
    fullscreen: {
      type: Boolean,
      default: false,
    },
    lists: {
      type: Array,
      default: [],
    },
  },
};
</script>

<style scoped>
.sidebar {
  @apply left-0 top-0 h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll pb-6
}

.sidebar:not(.fullscreen) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.fullscreen) {
  @apply -translate-x-full;
}

.sidebar.fullscreen {
  @apply w-screen bottom-0 bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.fullscreen {
  @apply translate-y-full opacity-0;
}
</style>