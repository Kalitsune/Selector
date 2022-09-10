<script>
    import { decodeCredential, googleOneTap } from 'vue3-google-login';
    import GoogleSignInButton from './GoogleSignInButton.vue';

    export default {
        async mounted() {
            googleOneTap({ autoLogin: true })
            .then((response) => {
                
                const userData = decodeCredential(response)
                console.log("data", userData)

            }).catch(err => {
                if (err === "Prompt was suppressed by user'. Refer https://developers.google.com/identity/gsi/web/guides/features#exponential_cooldown for more info"){ return; }
            });
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