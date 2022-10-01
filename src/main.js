import { createApp } from 'vue'
import vueCookies from 'vue-cookies'
import { createRouter, createWebHistory } from 'vue-router'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faBars, faListUl } from '@fortawesome/free-solid-svg-icons'

library.add(faBars, faListUl)

import css from './index.css'
import Index from './pages/Index.vue'
import Welcome from "./pages/Welcome.vue";
import App from './pages/App.vue'

let app = createApp(Index)

// Use vue-router
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { name:"login", path: '/', component: Welcome},
        { name: "app", path: '/app/:mode?/:listId?/', component: App},
        { name: "404",path: '/:pathMatch(.*)*', redirect: '/' }
    ]
})
app.use(router)

//use vue-cookies to check the presence of some cookies
app.use(vueCookies)

app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')