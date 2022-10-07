<template>
  <context-menu :fullscreen="fullscreen" ref="menu">
    <context-menu-item icon="fa-solid fa-pencil" text="Rename" type="disabled" tooltip="Renomez ou changez l'icone de votre liste"/>
    <context-menu-item icon="fa-solid fa-clone" text="Duplicate" type="disabled" tooltip="Maintenant vous en avez deux !"/>
    <context-menu-item icon="fa-solid fa-share-nodes" text="Share" type="disabled" tooltip="Obtenez un lien partageable pour vôtre liste !"/>
    <context-menu-item icon="fa-solid fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre liste de manière définitive."/>
  </context-menu>

  <div :class="{'collapsed': collapsed, 'fullscreen': fullscreen}" class="sidebar">
    <ul>
      <li v-if="lists.length > 0" v-for="list in lists">
        <SidebarItem :fullscreen="fullscreen" :list="list" :disabled="isDisabled(list)" :isSelected="isSelected(list)" @openContextMenu="evtData => this.$refs.menu.open(evtData)"/>
      </li>
      <li v-else>
        <SidebarItem :fullscreen="fullscreen" :list="{ name: 'There\'s no list to show!', id: 0}" disabled/>
      </li>
    </ul>
  </div>

</template>

<script>

import SidebarItem from "./SidebarItem.vue";
import ContextMenu from "./ContextMenu.vue";
import ContextMenuItem from "./ContextMenuItem.vue";

export default {
  name: "Sidebar",
  components: {SidebarItem, ContextMenu, ContextMenuItem},
  methods: {
    isSelected(list) {
      //check if the provided element is currently selected
      return this.$route.params.listId === list.id;
    },
    isDisabled(list) {
      //check if the provided element is currently disabled
      return list.id === 0;
    },
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    },
    fullscreen: {
      type: Boolean,
      default: false,
    },
    lists: {
      type: Array,
      default: [],
    },
  },
};
</script>

<style scoped>
.sidebar {
  @apply left-0 top-0 h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll pb-6
}

.sidebar:not(.fullscreen) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.fullscreen) {
  @apply -translate-x-full;
}

.sidebar.fullscreen {
  @apply w-screen bottom-0 bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.fullscreen {
  @apply translate-y-full opacity-0;
}
</style>