<template>
  <!-- topbar -->
  <span class="left-0 top-0 py-6 w-screen bg-neutral-50 dark:bg-neutral-700 flex flex-row">
    <!-- logo area -->
    <span class="w-28 md:w-72 left-0 top-0 flex flex-row justify-evenly pl-0 md:pl-5">
      <!-- Trigger for the sidebar -->
      <button @click="$emit('toggleSidebar')">
        <font-awesome-icon icon="fa-solid fa-bars" class="align-middle h-7 text-primary-light-500 dark:text-primary-dark-700 pl-1 md:pl-0" />
      </button>
      <!-- logo, dark & light -->
      <span class="pl-1 md:pr-5 md:pl-5 ">
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