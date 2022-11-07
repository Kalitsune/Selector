<template>
<div class="context-menu" :class="{'isMobile': isMobile}" v-show="visible" :style="style" ref="context" tabindex="0" @blur="close">
  <p v-if="isMobile && list !== undefined" class="text-neutral-400 w-full">{{list.name}}</p>
  <slot :list="list">
    <template v-if="menu === 'sidebar_item'">
      <context-menu-item icon="fas fa-pencil" text="Rename" type="disabled" tooltip="Renomez ou changez l'icone de votre liste" :list="list"/>
      <context-menu-item :handler="$api.createList" icon="fas fa-clone" text="Duplicate" type="classic" tooltip="Maintenant vous en avez deux !" :list="list"/>
      <context-menu-item icon="fas fa-share-nodes" text="Share" type="disabled" tooltip="Obtenez un lien partageable pour vôtre liste !" :list="list"/>
      <context-menu-item :handler="$api.deleteList" icon="fas fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre liste de manière définitive." :list="list"/>
    </template>

    <template v-else-if="menu === 'sidebar'">
      <context-menu-item icon="fas fa-plus" text="Create" type="disabled" tooltip="Créez une nouvelle liste."/>
      <context-menu-item :handler="$api.refreshLists" icon="fas fa-arrows-rotate" text="Refresh" type="classic" tooltip="Vos listes ne sont pas à jour? actualisez les !"/>
    </template>

    <template v-else>
      <context-menu-item icon="fa-solid fa-ban" text="There's nothing here" type="disabled"/>
    </template>

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
      this.menu = evtData.menu;
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
      list: Object,
      menu: "blank",
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