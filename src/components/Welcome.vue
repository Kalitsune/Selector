<script>
    import { googleOneTap } from 'vue3-google-login';
    import GoogleSignInButton from './GoogleSignInButton.vue';

    export default {
        methods: {
            callback(response){

            }
        },
        mounted() {
            googleOneTap({ autoLogin: true })
            .then((response) => {
                
                // This promise is resolved when user selects an account from the the One Tap prompt
                console.log("Handle the response", response)
            }).catch((error) => {
                console.log("Handle the error", error)
            })
        },
        components: { GoogleSignInButton }
    }
</script>

<template>
    <div id="container">
        <img id="logo" src="../assets/selector.svg" alt="Selector logo"/>
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

    #logo  {
        padding-bottom: 50px;
        width: 50vw;
    }

</style>