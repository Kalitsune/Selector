<template>
  <context-menu-item :handler="rename_list" icon="fas fa-pencil" text="Rename" type="disabled" tooltip="Renomez ou changez l'icone de votre liste" :list="list"/>
  <context-menu-item :handler="duplicate_list" icon="fas fa-clone" text="Duplicate" type="classic" tooltip="Maintenant vous en avez deux !" :list="list"/>
  <context-menu-item :handler="share_list" icon="fas fa-share-nodes" text="Share" type="disabled" tooltip="Obtenez un lien partageable pour vôtre liste !" :list="list"/>
  <context-menu-item :handler="delete_list" icon="fas fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre liste de manière définitive." :list="list"/>
</template>

<script>
import ContextMenuItem from "../ContextMenuItem.vue";

export default {
  name: "CtxMenuSidebarItems",
  props: {
    list: Object,
  },
  methods: {
    rename_list() {
      //not implemented yet
    },
    duplicate_list() {
      //create a new list
      this.$api.createList(this.list)
    },
    delete_list() {
      this.$api.deleteList(this.list)
    },
    share_list() {
      //not implemented yet
    },
  },
  components: {ContextMenuItem},
}
</script>

<style scoped>

</style>