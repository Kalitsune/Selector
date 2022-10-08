<template>
  <div class="fixed left-0 top-0 grid w-screen h-screen place-items-center">
    <div class="md:w-1/4 w-1/2 flex flex-col justify-center space-y-10 text-center">
      <div class="px-1/4">
        <img class="hidden dark:block" src="../assets/selector-dark.svg" alt="Selector logo"/>
        <img class="block dark:hidden" id="logo-light" src="../assets/selector.svg" alt="Selector logo"/>
        <p v-if="expired" class="text-base md:text-lg text-neutral-500 dark:text-neutral-200">Connexion requise</p>
      </div>
      <GoogleSignInButton :handler="handler" class="w-full"/>
    </div>
  </div>
</template>

<script>
import GoogleSignInButton from '../components/GoogleSignInButton.vue';
export default {
  name: "Login",
  components: {GoogleSignInButton},
  props: {
    expired: {
      type: Boolean,
      default: false,
    },
    handler: {
      type: Function,
      default: () => {},
    },
  },
}
</script>

<style>
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