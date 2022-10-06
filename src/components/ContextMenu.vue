<template>
<div class="context-menu" v-show="visible" :style="style" ref="context" tabindex="0" @blur="close">
  <slot>
    <context-menu-item icon="fa-solid fa-ban" text="There's nothing here" type="disabled"/>
  </slot>
</div>
</template>

<script>
import ContextMenuItem from "./ContextMenuItem.vue";
export default {
  name: "ContextMenu",
  components: {ContextMenuItem},
  methods: {
    close() {
      this.visible = false;
    },
    open(coords) {
      this.visible = true;
      this.x = coords.x;
      this.y = coords.y;
      this.$nextTick(() => this.$el.focus());
    }
  },
  computed: {
    // get position of context menu
    style() {
      return {
        left: this.x + 'px',
        top: this.y + 'px',
      };
    },
  },
  data() {
    return {
      visible: false,
      x: Number,
      y: Number
    }
  }
}
</script>

<style scoped>
.context-menu {
  @apply p-1 absolute z-50 shadow-lg bg-neutral-200 dark:bg-neutral-700 overflow-hidden flex flex-col;

  @apply bg-neutral-100 dark:bg-neutral-800 w-48 rounded-lg;
}
</style>