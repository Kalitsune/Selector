<script setup>
    import GoogleSignInButton from './GoogleSignInButton.vue';
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
        margin-top: auto;
        margin-bottom: auto;
    }

    #logo  {
        padding-bottom: 50px;
    }

</style>