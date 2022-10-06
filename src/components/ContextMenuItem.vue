<template>
  <span class="px-5 py-2 rounded-lg h-full" :title="tooltip" :class="type"><font-awesome-icon class="mr-2" :icon="icon" />{{text}}</span>
</template>

<script>
export default {
  name: "ContextMenuItem",
  props: {
    icon: String,
    text: String,
    tooltip: String,
    onClick: Function,
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