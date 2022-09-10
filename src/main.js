import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

let app = createApp(App)

let light_theme = {
    "primary": "#ff6700",
    "secondary": "#c4448c",
    "background": "#000000"
}

app.config.globalProperties.$theme = light_theme


app.mount('#app')