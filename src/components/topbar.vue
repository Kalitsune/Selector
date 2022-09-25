<template>
  <!-- topbar -->
  <span class="left-0 top-0 py-5 w-screen bg-neutral-50 dark:bg-neutral-700 flex flex-row">
    <!-- logo area -->
    <span class="w-25 md:w-72 left-0 top-0 flex flex-row justify-evenly pl-5 md:pl-0">
      <!-- Trigger for the sidebar -->
      <button @click="$emit('toggleSidebar')">
        <font-awesome-icon icon="fa-solid fa-bars" class="align-middle h-7 text-accent-500 dark:text-neutral-200" />
      </button>

      <!-- logo, dark & light -->
      <span class="pr-5 pl-5 md:pl-2">
        <!--big logos-->
        <span class="hidden md:block">
          <img class="dark:block hidden h-9 self-center" src="../assets/selector-dark.svg" alt="@Selector" />
          <img class="dark:hidden block h-9 self-center" src="../assets/selector.svg" alt="@Selector" />
        </span>

        <!--short logos-->
        <span class="block md:hidden">
          <img class="dark:block hidden h-9 self-center" src="../assets/selector-short-dark.svg" alt="@Selector" />
          <img class="dark:hidden block h-9 self-center" src="../assets/selector-short.svg" alt="@Selector" />
        </span>
      </span>
    </span>
  </span>
</template>

<script>

export default {
  name: "topbar",
}
</script>

<style scoped>

</style>