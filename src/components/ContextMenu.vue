<template>
<div class="context-menu" :class="{'isMobile': isMobile}" v-show="visible" :style="style" ref="context" tabindex="0" @blur="close">
  <p v-show="isMobile" class="text-neutral-400 w-full">{{list.name}}</p>
  <slot :list="list">
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
    open(evtData) {
      this.visible = true;
      this.x = evtData.x;
      this.y = evtData.y;
      this.list = evtData.list;
      this.$nextTick(() => this.$el.focus());
    }
  },
  computed: {
    // get position of context menu
    style() {
      return {
        left: this.$store.state.isMobile ? '0px' : this.x + 'px',
        top: this.$store.state.isMobile ? 'auto' : this.y + 'px',
      };
    },
    isMobile() {
      //check if the sidebar is isMobile
      return this.$store.state.isMobile;
    }
  },
  data() {
    return {
      visible: false,
      x: Number,
      y: Number,
      list: Object
    }
  }
}
</script>

<style scoped>
.context-menu {
  @apply p-1 fixed z-50 shadow-lg bg-neutral-200 dark:bg-neutral-700 overflow-hidden flex flex-col;

  @apply bg-neutral-100 dark:bg-neutral-800 w-48 rounded-lg;
}

.isMobile.context-menu {
  @apply w-full left-0 bottom-0 overflow-y-scroll rounded-none;

  max-height: 50%;
}
</style>