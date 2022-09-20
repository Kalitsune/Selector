<template>
  <div :class="{collapsed: collapsed}" class="sidebar flex flex-col">
    Hello
  </div>
</template>

<script>
export default {
  name: "Sidebar",
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
  @apply w-72 bg-neutral-50 dark:bg-neutral-700 left-0 bottom-0 h-full transition-transform ease-in-out duration-500;
}
.collapsed.sidebar {
  @apply -translate-x-full;
;
}
</style>