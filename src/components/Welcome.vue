<template>
    <div id="rainbow"></div>
    <div>
        <img src="../assets/selector.svg" alt="Selector logo"/>
    </div>

</template>

<script>
</script>

<style>
    body {
        background-color: v-bind(this.colors.background);
    }

    #rainbow {
        background: linear-gradient(v-bind(this.colors.primary), v-bind(this.colors.secondary));
        z-index: -999;

        position: absolute;
        width: 100vw;
        height: 100vh;
        left:0;
        top: 0;
    }

</style>