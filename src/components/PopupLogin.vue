<template>
<div v-if="visible" class="w-1/2 h-1/2 rounded-xl z-50 fixed align-middle mx-auto ">
  <Welcome></Welcome>
</div>
</template>

<script>
import Welcome from "../pages/Welcome.vue";
export default {
  name: "PopupLogin",
  components: {Welcome},
  methods: {
    close() {
      this.visible = false;
    },
    open() {
      this.visible = true;
    }
  },
  data() {
    return {
      visible: false,
    };
  },
}
</script>

<style scoped>

</style>