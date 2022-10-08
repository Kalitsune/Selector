<template>
  <div class="h-screen w-screen bg-neutral-200 dark:bg-neutral-800 overflow-hidden">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "index",
  methods: {
    isMobile() {
      return window.innerWidth < 	1024;
    },
    updateIsMobile() {
      //handle resize operation
      this.$store.commit("setIsMobile", this.isMobile());
    }
  },
  created() {
    //handle screen resizing
    window.addEventListener('resize', this.updateIsMobile);

    //set the default value for the isMobile property
    this.updateIsMobile();
  },
}
</script>

<style scoped>

</style>