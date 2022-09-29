import { createApp } from 'vue'
import vueCookies from 'vue-cookies'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faBars, faListUl } from '@fortawesome/free-solid-svg-icons'

library.add(faBars, faListUl)

import index from './pages/index.vue'
import css from './index.css'

let app = createApp(index)

//use vue-cookies to check the presence of some cookies
app.use(vueCookies)

app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')