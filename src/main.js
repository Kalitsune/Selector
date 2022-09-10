import { createApp } from 'vue'
import App from './App.vue'

import vue3GoogleLogin from 'vue3-google-login'
let app = createApp(App)

app.use(vue3GoogleLogin, {
    clientId: import.meta.env.VITE_GOOGLE_CLIENT_ID
})

let light_theme = {
    "primary": "#ff6700",
    "secondary": "#c4448c",
    "background": "#000000"
}

app.config.globalProperties.$theme = light_theme


app.mount('#app')