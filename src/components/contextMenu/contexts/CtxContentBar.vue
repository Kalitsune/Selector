<template>
  <context-menu-item :handler="createElement" icon="fas fa-plus" text="Create" type="classic" tooltip="Créer un nouvel élément."/>
  <context-menu-item :handler="() => $api.getListById(list.id)" icon="fas fa-arrows-rotate" text="Refresh" type="classic" tooltip="Actualiser un élément!"/>
</template>

<script>
import ContextMenuItem from "../ContextMenuItem.vue";

export default {
  name: "CtxMenuSidebar",
  components: {ContextMenuItem},
  props: {
    list: {
      type: Object,
      required: true,
    }
  },
  methods: {
    createElement() {
      //create a new element
      this.list.elements.push({
        name: "New element",
        probability: 0,
        icon: "fas fa-file-alt",
      });

      //add the list to the buffer changes
      this.$store.commit("addUpdatedList", this.list);
    },
  }
}
</script>

<style scoped>

</style>