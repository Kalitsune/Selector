<template>
  <context-menu-item icon="fas fa-pencil" text="Rename" type="disabled" tooltip="Renomez ou changez l'icone de votre liste" :list="list"/>
  <context-menu-item :handler="$api.createList" icon="fas fa-clone" text="Duplicate" type="classic" tooltip="Maintenant vous en avez deux !" :list="list"/>
  <context-menu-item icon="fas fa-share-nodes" text="Share" type="disabled" tooltip="Obtenez un lien partageable pour vôtre liste !" :list="list"/>
  <context-menu-item :handler="$api.deleteList" icon="fas fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre liste de manière définitive." :list="list"/>
</template>

<script>
import ContextMenuItem from "../ContextMenuItem.vue";

export default {
  name: "CtxMenuSidebarItems",
  props: {
    list: Object,
  },
  components: {ContextMenuItem},
}
</script>

<style scoped>

</style>