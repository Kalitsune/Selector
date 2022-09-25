<template>
  <div :class="{'collapsed': collapsed, 'fullscreen': fullscreen}" class="sidebar">
    Hello
  </div>
</template>

<script>

export default {
  name: "Sidebar",
  methods: {
    isFullscreen() {
      return window.innerWidth < 768;
    },
    onResize() {
      //handle resize operation
      this.fullscreen = this.isFullscreen();
    }
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    },
  },
  mounted() {
    window.addEventListener('resize', this.onResize);
  },
  unmounted() {
    window.removeEventListener('resize', this.onResize);
  },
  data() {
    return {
      fullscreen: this.isFullscreen()
    };
  },
};
</script>

<style scoped>
.sidebar {
  @apply left-0 bottom-0 h-full transition-all ease-in-out duration-500 flex flex-col
}

.sidebar:not(.fullscreen) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.fullscreen) {
  @apply -translate-x-full;
;
}

.sidebar.fullscreen {
  @apply w-screen bg-neutral-100 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.fullscreen {
  @apply translate-y-full opacity-0;
}
</style>