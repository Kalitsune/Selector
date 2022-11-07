<template>
  <context-menu-item :handler="openCreateListDialogue" icon="fas fa-plus" text="Create" type="primary" tooltip="Créez une nouvelle liste."/>
  <context-menu-item :handler="$api.refreshLists" icon="fas fa-arrows-rotate" text="Refresh" type="classic" tooltip="Vos listes ne sont pas à jour? actualisez les !"/>
</template>

<script>
import ContextMenuItem from "../ContextMenuItem.vue";

export default {
  name: "ContextMenuSidebar",
  methods: {
    openCreateListDialogue() {
      this.$refs.ListCreateForm.open();
    }
  },
  components: {ContextMenuItem},
}
</script>

<style scoped>

</style>