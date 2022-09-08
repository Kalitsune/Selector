import { createApp } from 'vue'
import './style.css'
import App from './App.vue'


let app = createApp(App)

let colors = {
    "primary": "#ff6700",
    "secondary": "#c4448c",
    "background": "#2E2E2E"
}
app.config.globalProperties.colors = colors

app.mount('#app')
