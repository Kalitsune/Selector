<template>
  <context-menu-item :handler="$api.createList" icon="fas fa-plus" text="Create" type="classic" tooltip="Créez une nouvelle liste."/>
  <context-menu-item :handler="$api.refreshLists" icon="fas fa-arrows-rotate" text="Refresh" type="classic" tooltip="Vos listes ne sont pas à jour? actualisez les !"/>
</template>

<script>
import ContextMenuItem from "../ContextMenuItem.vue";

export default {
  name: "CtxMenuSidebar",
  components: {ContextMenuItem},
}
</script>

<style scoped>

</style>