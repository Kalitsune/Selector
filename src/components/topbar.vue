<template xmlns="http://www.w3.org/1999/html">
  <!-- topbar -->
  <div class="left-0 top-0 py-5 w-screen bg-neutral-50 dark:bg-neutral-800">
    <!-- logo area -->
    <span class="w-72 left-0 top-0 flex flex-row justify-evenly">
      <!-- Trigger for the sidebar -->
      <button @click="$emit('toggleSidebar')">
        <font-awesome-icon icon="fa-solid fa-bars" class="align-middle h-7" />
      </button>

      <!-- logo, dark & light -->
      <span class="pr-4 pl-2">
        <img class="dark:block hidden h-9 self-center" src="../assets/selector-dark.svg" alt="@Selector" />
        <img class="dark:hidden block h-9 self-center" src="../assets/selector.svg" alt="@Selector" />
      </span>
    </span>
  </div>
</template>

<script>

export default {
  name: "topbar",
}
</script>

<style scoped>

</style>