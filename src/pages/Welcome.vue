<script>
    import Login from "../components/Login.vue";
    export default {
      name: "Welcome",
      components: {Login },
      mounted() {
        //check if the user is authenticated
        this.$api._getLists().then(lists => {
          if (lists.constructor === Array) {
            //redirect to the app page
            this.$router.push({name: "app", params: {listId: this.$route.query.listId, mode: this.$route.query.mode}});
          }
        }).catch(() => {});
      },
      methods: {
        toApp() {
          //redirect to the app page
          this.$router.push({name: "app", params: {listId: this.$route.query.listId, mode: this.$route.query.mode}});
        },
      },
    }
</script>

<template>
  <login :handler="toApp"/>
</template>