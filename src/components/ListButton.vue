<template>
  <div :class="{selected: selected, disabled: disabled}" id="button-bg" class="w-full py-4 pl-5 flex flex-row space-x-2">
    <font-awesome-icon class="list-element h-6" icon="fa-solid fa-list-ul"/>
    <p class="list-element h-7">{{list.name}}</p>
  </div>
</template>

<script>
export default {
  name: "ListButton",
  props: {
    list: Object,
    selected: {
      type: Boolean,
      default: false
    },
    disabled: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style scoped>

.disabled#button-bg, .disabled#button-bg:hover,
.disabled .list-element, .disabled .list-element:hover {
  @apply text-neutral-400 dark:text-neutral-600 bg-neutral-50 dark:bg-neutral-700 cursor-not-allowed !important;
}

.selected#button-bg, .selected#button-bg:hover {
  @apply bg-primary-light-100 dark:bg-primary-dark-800;
}

#button-bg:hover {
  @apply bg-primary-light-50 dark:bg-primary-dark-900;
}

.selected .list-element {
  @apply text-primary-light-500 dark:text-primary-dark-700;
}
.list-element {
  @apply align-middle text-accent-400 dark:text-accent-600;
}
</style>