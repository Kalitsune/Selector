import { createApp } from 'vue'
import vueCookies from 'vue-cookies'
import { Vue3Mq } from 'vue3-mq'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faBars } from '@fortawesome/free-solid-svg-icons'

library.add(faBars)

import index from './pages/index.vue'
import css from './index.css'

let app = createApp(index)

//use vue-cookies to check the presence of some cookies
app.use(vueCookies)

//use vue-mq to enable responsive design and alternative components (see https://www.digitalocean.com/community/tutorials/vuejs-vue-media-queries)
app.use(Vue3Mq, {
    breakpoints: {
        "sm": 640,
        "md": 768,
        "lg": 1024,
        "xl": 1280,
        "xxl": 1536,
    }
})


app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')