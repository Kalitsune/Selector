<template>
  <context-menu-item :handler="rename_element" icon="fas fa-pencil" text="Rename" type="disabled" tooltip="Renomez de vôtre élément."/>
  <context-menu-item :handler="duplicate_element" icon="fas fa-clone" text="Duplicate" type="classic" tooltip="Maintenant vous en avez deux !"/>
  <context-menu-item :handler="delete_element" icon="fas fa-trash-can" text="Delete" type="destructive" tooltip="Suprimme vôtre élément."/>
</template>

<script>
import ContextMenuItem from "../ContextMenuItem.vue";

export default {
  name: "CtxMenuSidebarItems",
  props: {
    ctx: Object,
  },
  methods: {
    rename_element() {
      //not implemented yet
    },
    duplicate_element() {
      //add the element in the list
      this.ctx.list.elements.push(this.ctx.element);

      //mark the list as edited
      this.$store.commit("addUpdatedList", this.ctx.list.id);
    },
    delete_element() {
      //remove the element from the list
      this.ctx.list.elements.splice(this.ctx.list.elements.indexOf(this.ctx.element), 1);

      //mark the list as edited
      this.$store.commit("addUpdatedList", this.ctx.list.id);
    },
  },
  components: {ContextMenuItem},
}
</script>

<style scoped>

</style>