<template>
  <div @click="clickHandler" class="context-menu-item" :title="tooltip" :class="type"><font-awesome-icon class="m-2" :icon="icon" />{{text}}</div>
</template>

<script>
export default {
  name: "ContextMenuItem",
  methods: {
    clickHandler(evt) {

      if (this.type === "disabled" || this.handler === undefined) {
        evt.preventDefault();
        evt.stopPropagation();
        return;
      }

      //close the context menu
      this.$parent.close();

      //call the handler
      this.handler(this.list);
    },
  },
  props: {
    list: {
      type: Object,
      required: false,
    },
    icon: String,
    text: String,
    tooltip: String,
    handler: Function,
    type: {
      type: String,
      default: "classic",
      validator(value) {
        return ['classic', 'primary', 'destructive', 'disabled'].includes(value)
      }
    }
  }
}
</script>

<style scoped>
.context-menu-item {
  @apply flex flex-row items-center justify-start rounded-lg cursor-pointer
          p-4 lg:p-2 h-full w-full mr-14
          text-xl lg:text-base
}

.disabled {
  @apply text-neutral-400 dark:text-neutral-600 cursor-not-allowed !important;
}

.classic {
  @apply text-neutral-700 dark:text-neutral-300;
}

.classic:hover {
  @apply bg-primary-light-300 dark:bg-primary-dark-400 text-neutral-50;
}

.destructive {
  @apply text-red-500;
}

.destructive:hover {
  @apply bg-red-500 text-neutral-50;
}
</style>