<template>
  <!-- menu -->
  <div v-show="visible" class="fixed w-3/4 left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 shadow-lg
            bg-neutral-50 dark:bg-neutral-700 overflow-hidden rounded-xl p-5 z-50">

    <!-- title & close button -->
    <span class="w-full flex flex-row">
      <!-- title -->
      <span class="mr-auto">{{ name }}</span>

      <!-- close button -->
      <button class="ml-auto border-neutral-300 border-2 rounded-[50%] h-7 w-7" @click="close">
        <font-awesome-icon icon="fa-solid fa-times" class="text-neutral-500 dark:text-neutral-600" />
      </button>
    </span>

    <!-- menu content -->
    <slot></slot>
  </div>

  <!-- back overlay -->
  <div v-show="visible" class="w-screen h-screen block fixed top-0 left-0 bg-neutral-500 z-40 opacity-50" @click="close()"/>

</template>

<script>
export default {
  name: "basicPopupMenu",
  props: {
    name: {
      type: String,
      default: "Unknown Menu"
    }
  },
  data() {
    return {
      visible: false
    }
  },
  methods: {
    close() {
      this.visible = false;
    },
    open() {
      this.visible = true;
    }
  }
}
</script>

<style scoped>

</style>