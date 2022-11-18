<template>
  <button @contextmenu="openContextMenu" :class="{disabled: disabled}" class="background">
    <font-awesome-icon class="element h-6" icon="fa-solid fa-cube"/>
    <span class="element h-7">{{element.name}}</span>
  </button>
</template>

<script>
export default {
  name: "ContentItem",
  props: {
    element: Object,
    list: Object,
    disabled: Boolean,
  },
  methods: {
    openContextMenu(event) {
      //get the coordinates of the click
      let left = event.pageX || event.clientX;
      let top = event.pageY || event.clientY;

      //prevent the default context menu and prevent triggering other events
      event.preventDefault();
      event.stopPropagation();

      //open the context menu
      this.$root.$refs.ContextMenu.open({left, top, context: {element: this.element, list: this.list}, menu: "contentBarItem"});
    }
  }
}
</script>

<style scoped>
.background.disabled, .background.disabled:hover,
.disabled .element, .disabled .element:hover {
  @apply text-neutral-400 dark:text-neutral-600 bg-transparent cursor-not-allowed !important;
}


.background.selected, .background.selected:hover {
  @apply bg-primary-light-100 dark:bg-neutral-900 dark:bg-opacity-30;
}

.background:hover {
  @apply bg-primary-light-100 bg-opacity-30  dark:bg-neutral-800;
}

.background {
  @apply w-11/12 rounded-2xl py-4 px-5 flex flex-row space-x-2
}

.selected .element {
  @apply text-primary-light-500 dark:text-primary-dark-700;
}
.element {
  @apply align-middle text-accent-400 dark:text-accent-500;
}
</style>