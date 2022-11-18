<template>
  <!-- content bar -->
  <div :class="{'collapsed': collapsed, 'isMobile': isMobile}" class="content-bar" @contextmenu="openContextMenu" >
    <!-- content bar title -->
    <div class="font-semibold pt-5 text-accent-400">
      <h1>Contenu de la liste</h1>
    </div>

    <!-- content bar items -->
    <ul class="w-11/12">
      <!-- list content -->
      <li v-if="!isDisabled(list) && list?.elements?.length > 0" v-for="element in list.elements">
        <ContentItem :element="element" :list="list" />
      </li>

      <!-- if the list is empty -->
      <li v-else>
        <ContentItem :element="{name:'There\'s nothing here yet...'}" disabled />
      </li>
    </ul>
  </div>

</template>

<script>

import ContextMenu from "../contextMenu/ContextMenu.vue";
import ContextMenuItem from "../contextMenu/ContextMenuItem.vue";
import ContentItem from "./ContentItem.vue";

export default {
  name: "ContentBar",
  components: {ContextMenu, ContextMenuItem, ContentItem},
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
      this.$root.$refs.ContextMenu.open({left, top, context: {list: this.list}, menu: "contentBar"});
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
.content-bar:not(.isMobile) {
  @apply w-96 bg-neutral-50 dark:bg-neutral-700 h-full transition-all ease-in-out duration-500 flex flex-col
  overflow-y-scroll pb-6 float-right items-center;
}
.collapsed.content-bar:not(.isMobile) {
  @apply translate-x-full;
}

/*TODO: implement content bar on the mobile interface*/

.content-bar.isMobile {
  @apply hidden;
}
.collapsed.content-bar.isMobile {
  @apply hidden;
}
</style>