<script>
    import GoogleSignInButton from '../components/GoogleSignInButton.vue';

    export default {
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