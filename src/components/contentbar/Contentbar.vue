<template>
  <!-- content bar -->
  <div :class="{'collapsed': collapsed, 'isMobile': isMobile}" class="content-bar" @contextmenu="openContextMenu" >
    <!-- content bar title -->
    <div class="font-semibold text-center pt-5 text-accent-400">
      <h1>Contenu de la liste</h1>
    </div>
    <!-- content bar items -->
    <ul>
      <!-- list content -->
      <li v-if="!isDisabled(list)" v-for="element in list.elements">

      </li>

      <!-- if the list is empty -->
      <li v-else>

      </li>
    </ul>
  </div>

</template>

<script>

import ContextMenu from "../contextMenu/ContextMenu.vue";
import ContextMenuItem from "../contextMenu/ContextMenuItem.vue";

export default {
  name: "ContentBar",
  components: {ContextMenu, ContextMenuItem},
  methods: {
    isDisabled(list) {
      //check if the provided element is currently disabled
      return list === undefined || list.id === 0;
    },
    openContextMenu(event) {
      //get the coordinates of the click
      let left = event.pageX || event.clientX;
      let top = event.pageY || event.clientY;

      //prevent the default context menu and prevent triggering other events
      event.preventDefault();
      event.stopPropagation();

      //open the context menu
      this.$root.$refs.ContextMenu.open({left, top, menu: "sidebar"});
    }
  },
  computed: {
    isMobile() {
      //check if the sidebar is isMobile
      return this.$store.state.isMobile;
    },
    list() {
        //make the list available to the template
        return this.$store.state.lists.find(list => list.id === this.$route.params.listId);
      }
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    }
  },
};
</script>

<style scoped>
.content-bar {
  @apply h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll pb-6 float-right;
}

.content-bar:not(.isMobile) {
  @apply w-96 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.content-bar:not(.isMobile) {
  @apply translate-x-full;
}

.content-bar.isMobile {
  @apply w-screen bottom-0 bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.content-bar.isMobile {
  @apply translate-y-full opacity-0;
}
</style>