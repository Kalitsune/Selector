<template>
<!-- detect touch outside the context menu on every platform (mainly here for mobile devices) -->
<div v-show="visible" class="w-screen h-screen block fixed top-0 left-0 z-40" @click="close()"/>

<div class="context-menu" :class="{'isMobile': isMobile}" v-show="visible" :style="style" ref="context" tabindex="0" @blur="close">
  <p v-if="isMobile && list !== undefined" class="text-neutral-400 w-full">{{list.name}}</p>
  <slot :list="list">
    <!-- ../contextMenu/contexts/CtxSidebarItems.vue -->
    <ContextMenuSidebarItems v-if="menu === 'sidebar_item'" :list="list"/>

    <!-- ../contextMenu/contexts/CtxSidebar.vue -->
    <context-menu-sidebar v-else-if="menu === 'sidebar'"/>

    <!-- ../contextMenu/contexts/CtxSidebar.vue -->
    <context-menu-content-bar v-else-if="menu === 'contentBar'" :list="list"/>

  </slot>
</div>
</template>

<script>
import ContextMenuItem from "./ContextMenuItem.vue";
import ContextMenuSidebarItems from "./contexts/CtxSidebarItems.vue";
import ContextMenuSidebar from "./contexts/CtxSidebar.vue";
import ContextMenuContentBar from "./contexts/CtxContentBar.vue";

export default {
  name: "ContextMenu",
  components: {ContextMenuSidebar, ContextMenuSidebarItems, ContextMenuItem, ContextMenuContentBar},
  methods: {
    close(condition) {
      //check if there are no conditions provided
      if (condition === undefined) {
        this.visible = false;
      // check if the condition provided a list that match with the current one
      } else if (condition.list === this.list) {
        this.visible = false;
      }
    },
    open(evtData) {
      //define the menu config from the event data
      this.visible = true;
      this.left = evtData.left + "px";
      this.top = evtData.top + "px";
      this.right = "auto";
      this.bottom = "auto";
      this.menu = evtData.menu;
      this.list = evtData.list;

      //wait for the element to render
      this.$nextTick(() => {
        //check for X axis overflow
        if (this.$refs.context.getBoundingClientRect().right > window.innerWidth) {
          this.left = "auto";
          this.right = "0px";
        }
        //check for Y axis overflow
        if (this.$refs.context.getBoundingClientRect().bottom > window.innerHeight) {
          this.top = "auto";
          this.bottom = "0px";
        }
        //focus the element
        this.$refs.context.focus()

      });
    }
  },
  computed: {
    // get position of context menu
    style() {
      return {
        left: this.$store.state.isMobile ? '0px' : this.left,
        right: this.$store.state.isMobile ? '0px' : this.right,
        top: this.$store.state.isMobile ? 'auto' : this.top,
        bottom: this.$store.state.isMobile ? '0px' : this.bottom,
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
      left: String,
      right: String,
      top: String,
      bottom: String,
      list: Object,
      menu: "blank",
      reversed: false,
    }
  }
}
</script>

<style scoped>
.context-menu {
  @apply p-1 fixed z-50 shadow-lg bg-neutral-200 dark:bg-neutral-700 overflow-hidden flex flex-col;

  @apply bg-neutral-100 dark:bg-neutral-800 min-w-[12rem] rounded-lg;
}

.isMobile.context-menu {
  @apply w-full left-0 bottom-0 overflow-y-scroll rounded-none;

  max-height: 50%;
}
</style>