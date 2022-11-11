<template>
  <!-- topbar -->
  <span class="left-0 top-0 py-6 w-screen bg-neutral-50 dark:bg-neutral-700 grid grid-cols-3">
    <!-- logo area -->
    <span class="justify-self-start w-28 md:w-72 left-0 top-0 flex flex-row justify-evenly pl-0 md:pl-5">
      <!-- Trigger for the sidebar -->
      <button @click="toggleSidebar">
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

    <span class="justify-self-center">Middle</span>

    <!-- right side -->
    <span class="justify-self-end">
      <!-- Trigger for the contentBar -->
      <button @click="toggleContentBar">
        <font-awesome-icon icon="fa-solid fa-layer-group" class="align-middle h-7 text-primary-light-500 dark:text-primary-dark-700 pr-7 md:pl-6" />
      </button>
    </span>
  </span>
</template>

<script>

export default {
  name: "topbar",
  methods: {
    toggleSidebar() {
      //toggle the sidebar
      this.$root.toggleSidebar();
    },
    toggleContentBar() {
      //toggle the contentBar
      this.$root.toggleContentBar();
    }
  }
}
</script>

<style scoped>

</style>