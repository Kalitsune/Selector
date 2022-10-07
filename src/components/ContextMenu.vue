<template>
<div class="context-menu" :class="{'fullscreen': fullscreen}" v-show="visible" :style="style" ref="context" tabindex="0" @blur="close">
  <p v-show="fullscreen" class="text-neutral-400 w-full">{{list.name}}</p>
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
        left: this.fullscreen ? '0px' : this.x + 'px',
        top: this.fullscreen ? 'auto' : this.y + 'px',
      };
    },
  },
  props: {
    fullscreen: Boolean,
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

.fullscreen.context-menu {
  @apply w-full left-0 bottom-0 overflow-y-scroll rounded-none;

  max-height: 50%;
}
</style>