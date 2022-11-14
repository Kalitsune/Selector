<template>
<span :class="visible ? 'bottom-5' : '-bottom-full'" class="fixed left-1/2 -translate-x-1/2 h-16 bg-primary-light-50 dark:bg-primary-dark-900 rounded-xl
               p-2 px-3 grid md:grid-cols-2 place-content-center transition-all ease-in-out delay-100">

  <!-- text area -->
  <span class="relative flex flex-row h-auto top-1/4 pl-2">
    <!-- save icon -->
    <font-awesome-icon class="hidden md:block h-6 text-primary-light-500 dark:text-primary-dark-600" icon="fas fa-floppy-disk"/>
    <!-- text -->
    <span class="pl-3 font-extrabold text-primary-light-500 dark:text-primary-dark-600">
      <a class="hidden md:block xl:hidden">Save?</a>
      <a class="hidden xl:block ">Careful! You have unsaved changes.</a>
    </span>
  </span>

  <!-- buttons -->
  <span class="justify-self-end flex flex-row">
    <!-- cancel button -->
    <button @click="cancel()" class="text-primary-light-300 dark:text-primary-dark-700 font-semibold hover:underline pl-3 pr-8 md:pr-5">cancel</button>

    <!-- save button -->
    <span class="bg-primary-light-500 hover:bg-primary-light-600 text-primary-light-50 dark:bg-primary-dark-600 dark:hover:bg-primary-dark-700 dark:text-primary-dark-200 p-2 rounded font-semibold">
      <button @click="save()" class="hidden lg:block  px-4">Save changes</button>
      <button @click="save()" class="block lg:hidden  px-6">Save</button>
    </span>
  </span>
</span>
</template>

<script>
export default {
  name: "saveButton",
  props: {
    visible: Boolean,
    cancel: Function,
    save: Function,
  }
}
</script>

<style scoped>

</style>