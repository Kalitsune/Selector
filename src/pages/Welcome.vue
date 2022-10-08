<script>
    import Login from "../components/Login.vue";
    export default {
      name: "Welcome",
      components: {Login },
      mounted() {
        //check if the user is authenticated
        this.$api.getLists().then(lists => {
          if (lists.constructor === Array) {
            //redirect to the app page
            this.toApp();
          }
        }).catch(() => {});
      },
      methods: {
        toApp() {
          //redirect to the app page
          this.$router.push({name: "app"});
        },
      },
    }
</script>

<template>
  <login :handler="toApp"/>
</template>