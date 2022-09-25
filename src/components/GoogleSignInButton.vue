<script>
    export default {
        name: "GoogleSignInButton",
        methods : {
          signIn() {
            // open a new window pointing to the Google login url
            let win = window.open('/auth', '_blank', 'location=yes,height=570,width=520,scrollbars=yes,status=yes');
            // listen for the success event
            window.addEventListener('message', (event) => {
              // check if the event is coming from the correct origin
              if (event.source === win) {
                // check if the event is the success event
                if (event.data === 'success') {
                  // close the window
                  win.close();
                  // reload the page
                  window.location.reload();
                }
              }
            });
          }
        }
    }
    
</script>


<template>
    <button class="firebaseui-idp-google flex-center--all noselect md:w-1/4 w-1/2 bg-white dark:bg-neutral-700" id="google-signin" v-on:click="signIn" data-provider-id="google.com">

        <span class="firebaseui-idp-icon-wrapper bg-white p-1.5 rounded">
            <img class="firebaseui-idp-icon" alt="" src="https://www.gstatic.com/firebasejs/ui/2.0.0/images/auth/google.svg">
        </span>

        <span class="firebaseui-idp-text hidden xl:block dark:text-white">
            Se connecter avec Google
        </span>
        <span class="firebaseui-idp-text block xl:hidden dark:text-white">
             Connexion
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
    height: 40px;
    width: 40px;
}
</style>