import { createApp } from 'vue'
import vueCookies from 'vue-cookies'
import index from './pages/index.vue'
import IndexCss from './index.css'

let app = createApp(index)

app.use(vueCookies)


let light_theme = {
    "primary": "#ff6700",
    "secondary": "#c4448c",
    "background": "#000000"
}

app.config.globalProperties.$theme = light_theme

app.mount('#app')