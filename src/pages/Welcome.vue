<script>
    import GoogleSignInButton from '../components/GoogleSignInButton.vue';
    import api from "../api.js";

    export default {
      name: "Welcome",
      components: { GoogleSignInButton },
      mounted() {
        //check if the user is authenticated
        api.getLists().then(lists => {
          if (lists.constructor === Array) {
            //redirect to the app page
            this.$router.push({name: "app", params: {listId: this.$route.query.listId, mode: this.$route.query.mode}});
          }
        }).catch(() => {});
      }
    }
</script>

<template>
    <div id="container">
        <img class="logo" id="logo-dark" src="../assets/selector-dark.svg" alt="Selector logo"/>
        <img class="logo" id="logo-light" src="../assets/selector.svg" alt="Selector logo"/>
      <GoogleSignInButton/>
    </div>
</template>

<style>
    #container {
        display: flex;
        flex-direction: column;
        align-items: center;
        position: relative;
        top: 50vh;
        transform: translateY(-50%);
    }

    .logo  {
      padding-bottom: 50px;
      padding-left: 25vw;
      padding-right: 25vw;
    }

    .dark #logo-light {
        display: none;
    }

    .dark #logo-dark {
      display: block;
    }

    #logo-dark {
      display: none;
    }


</style>