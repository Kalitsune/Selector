<script>
    import { googleTokenLogin } from "vue3-google-login"

    export default {
        name: "GoogleSignInButton",
        methods: {
            async profileLoad() {
                const element = document.getElementById("google-signin")

                if (!element.classList.contains("waiting")) {
                    element.classList.add("waiting")
                    
                    //reset the animation
                    setTimeout(_ => {element.classList.remove("waiting")}, 5000)

                    // connect to google
                    googleTokenLogin().then(response => {
                        console.log(response);
                    });

                }
            }
        }
    }
    
</script>


<template>
    <button class="firebaseui-idp-google flex-center--all noselect" id="google-signin" v-on:click="profileLoad" data-provider-id="google.com">

        <span class="firebaseui-idp-icon-wrapper">
            <img class="firebaseui-idp-icon" alt="" src="https://www.gstatic.com/firebasejs/ui/2.0.0/images/auth/google.svg">
        </span>

        <span class="firebaseui-idp-text firebaseui-idp-text-long">
            Se connecter avec Google
        </span>  
    </button>
</template>

<style>
    @keyframes waiting {
    0% {
      opacity: 50%;
    }
    50% {
        opacity: 100%;
    }
    100% {
        opacity: 50%;
    }
}

* {
    -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
    box-sizing: border-box;
}

.noselect {
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
}

:focus {
    outline: 0;
}

.flex-center--all {
    display: flex;
    align-items: center;
    justify-content: center;
}

button.firebaseui-idp-google {
    padding: 12px 20px;
    background: white;
    border: 0;
    border-radius: 99rem;
    box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
    outline: 0;
    overflow: hidden;
    will-change: box-shadow;
    transform: scale(1);
    transition: 0.2s all;
    cursor: pointer;
    font-family: Roboto, Helvetica, Arial, sans-serif;
    width: 25%;
}

@media not all and (pointer: coarse) {
    button.firebaseui-idp-google:hover {
        box-shadow: 0 0 8px rgba(0, 0, 0, 0.18), 0 8px 16px rgba(0, 0, 0, 0.36);
    }
}

button.firebaseui-idp-google.waiting {
    padding: 12px;
    background: #f6f6f6;
    box-shadow: 0 0 8px rgba(0, 0, 0, 0.18), 0 8px 16px rgba(0, 0, 0, 0.36);
    filter: grayscale(100%);
    cursor: default;
}
button.firebaseui-idp-google.waiting .firebaseui-idp-text {
    height: 0;
    width: 0;
    margin: 0;
    opacity: 0;
}

.waiting img.firebaseui-idp-icon {
    animation: waiting 2s linear infinite;
}


button.firebaseui-idp-google .firebaseui-idp-icon-wrapper img {
    height: 100%;
    width: auto;
    transition: 0.2s all;
}
button.firebaseui-idp-google .firebaseui-idp-text {
    margin-left: 0.5rem;
    font-weight: 500;
    transition: 0.3s height width;
}
button.firebaseui-idp-google .firebaseui-idp-icon-wrapper {
    height: 25px;
    width: 25px;
}

@media (max-width: 600px) {
    button.firebaseui-idp-google .firebaseui-idp-icon-wrapper {
        height: 50px;
        width: 50px;
    }
    button.firebaseui-idp-google .firebaseui-idp-text-long {
        display: none;
    }
}
</style>