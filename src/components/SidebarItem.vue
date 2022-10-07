<template>
  <button @contextmenu="openContextMenu" @click="switchList" :class="{selected: isSelected, disabled: disabled}" id="button-bg" class="w-11/12 rounded-r-2xl py-4 pl-5 flex flex-row space-x-2">
    <font-awesome-icon class="list-element h-6" icon="fa-solid fa-list-ul"/>
    <span class="list-element h-7">{{list.name}}</span>
  </button>
</template>

<script>
export default {
  name: "SidebarItem",
  props: {
    list: Object,
    isSelected: {
      type: Boolean,
      default: false
    },
    disabled: {
      type: Boolean,
      default: false
    },
    fullscreen: Boolean
  },
  methods: {
    switchList() {
      //redirect to the new url by replacing the list parameter
      this.$router.push({name: "app", params: {listId: this.list.id, mode: this.$route.params.mode}});
    },
    openContextMenu(evt) {
      //get the coordinates of the click
      let x = evt.pageX || evt.clientX;
      let y = evt.pageY || evt.clientY;

      //prevent the default context menu
      evt.preventDefault();

      //open the context menu
      this.$emit('openContextMenu', {x, y, list: this.list})
    }
  },
  unmounted() {
    //emit an event to the parent component to close the context menu
    this.$emit('closeContextMenu');
  }
}
</script>

<style scoped>

.disabled#button-bg, .disabled#button-bg:hover,
.disabled .list-element, .disabled .list-element:hover {
  @apply text-neutral-400 dark:text-neutral-600 bg-neutral-50 dark:bg-neutral-700 cursor-not-allowed !important;
}

.selected#button-bg, .selected#button-bg:hover {
  @apply bg-primary-light-100 dark:bg-neutral-900 dark:bg-opacity-30;
}

#button-bg:hover {
  @apply bg-primary-light-100 bg-opacity-30  dark:bg-neutral-800;
}

.selected .list-element {
  @apply text-primary-light-500 dark:text-primary-dark-700;
}
.list-element {
  @apply align-middle text-accent-400 dark:text-accent-500;
}
</style>