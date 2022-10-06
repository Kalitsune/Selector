<template>
  <div :class="{'collapsed': collapsed, 'fullscreen': fullscreen}" class="sidebar">
    <ul>
      <li v-if="lists.length > 0" v-for="list in lists">
        <ListButton :list="list" :disabled="isDisabled(list)" :isSelected="isSelected(list)" @openContextMenu="coords => this.$emit('openContextMenu', coords)"/>
      </li>
      <li v-else>
        <ListButton :list="{ name: 'There\'s no list to show!', id: 0}" disabled/>
      </li>
    </ul>
  </div>

</template>

<script>

import ListButton from "./ListButton.vue";
export default {
  name: "Sidebar",
  components: {ListButton},
  methods: {
    isFullscreen() {
      return window.innerWidth < 	1024;
    },
    onResize() {
      //handle resize operation
      this.fullscreen = this.isFullscreen();
    },
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
    lists: {
      type: Array,
      default: [],
    },
  },
  created() {
    //handle screen resizing
    window.addEventListener('resize', this.onResize);
  },
  data() {
    return {
      fullscreen: this.isFullscreen(),
    };
  },
};
</script>

<style scoped>
.sidebar {
  @apply left-0 top-0 h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll
}

.sidebar:not(.fullscreen) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.fullscreen) {
  @apply -translate-x-full;
}

.sidebar.fullscreen {
  @apply w-screen bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.fullscreen {
  @apply translate-y-full opacity-0;
}
</style>