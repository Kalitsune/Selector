<template>
  <!-- sidebar items context menu -->
  <context-menu ref="SidebarItemsContextMenu" v-slot="slotProps">
    <context-menu-item icon="fas fa-pencil" text="Rename" type="disabled" tooltip="Renomez ou changez l'icone de votre liste" :list="slotProps.list"/>
    <context-menu-item icon="fas fa-clone" text="Duplicate" type="disabled" tooltip="Maintenant vous en avez deux !" :list="slotProps.list"/>
    <context-menu-item icon="fas fa-share-nodes" text="Share" type="disabled" tooltip="Obtenez un lien partageable pour vôtre liste !" :list="slotProps.list"/>
    <context-menu-item :handler="deleteList" icon="fas fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre liste de manière définitive." :list="slotProps.list"/>
  </context-menu>

  <!-- sidebar context menu -->
  <context-menu ref="SidebarContextMenu" v-slot="slotProps">
    <context-menu-item icon="fas fa-plus" text="Create" type="disabled" tooltip="Créez une nouvelle liste." :list="slotProps.list"/>
    <context-menu-item icon="fas fa-arrows-rotate" text="Refresh" type="disabled" tooltip="Vos listes ne sont pas à jour? actualisez les!" :list="slotProps.list"/>
  </context-menu>

  <!-- sidebar -->
  <div :class="{'collapsed': collapsed, 'isMobile': isMobile}" class="sidebar" @contextmenu="openSidebarContextMenu" >
    <!-- sidebar items -->
    <ul>
      <!-- array of lists -->
      <li v-if="lists.length > 0" v-for="list in lists">
        <SidebarItem :list="list" :disabled="isDisabled(list)" :isSelected="isSelected(list)" @closeContextMenu="this.$refs.SidebarItemsContextMenu.close()" @openContextMenu="evtData => this.$refs.SidebarItemsContextMenu.open(evtData)"/>
      </li>

      <!-- if there's no list -->
      <li v-else>
        <SidebarItem :list="{ name: 'There\'s no list to show!', id: 0}" disabled/>
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
    openSidebarContextMenu(event) {
      //get the coordinates of the click
      let x = event.pageX || event.clientX;
      let y = event.pageY || event.clientY;

      //prevent the default context menu and prevent triggering other events
      event.preventDefault();
      event.stopPropagation();

      //open the context menu
      this.$refs.SidebarContextMenu.open({x, y, list: {name: "sidebar", id: 0}});
    },
    deleteList(list) {
      //delete the list using the api easy handler
      this.$api.deleteList(list);
    }
  },
  computed: {
      lists() {
        //make the lists available to the template
        return this.$store.state.lists;
      },
      isMobile() {
        //check if the sidebar is isMobile
        return this.$store.state.isMobile;
      }
  },
  props: {
    collapsed: {
      type: Boolean,
      default: false,
    },
  },
};
</script>

<style scoped>
.sidebar {
  @apply left-0 top-0 h-full transition-all ease-in-out duration-500 flex flex-col overflow-y-scroll pb-6
}

.sidebar:not(.isMobile) {
  @apply w-72 bg-neutral-50 dark:bg-neutral-700;
}
.collapsed.sidebar:not(.isMobile) {
  @apply -translate-x-full;
}

.sidebar.isMobile {
  @apply w-screen bottom-0 bg-neutral-200 dark:bg-neutral-800 opacity-100;
}
.collapsed.sidebar.isMobile {
  @apply translate-y-full opacity-0;
}
</style>