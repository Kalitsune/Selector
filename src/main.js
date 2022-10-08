import { createApp } from 'vue'
import { createStore } from 'vuex'
import { createRouter, createWebHistory } from 'vue-router'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {fas} from '@fortawesome/free-solid-svg-icons'

import api from './api.js'

library.add(fas)

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

// Use vuex
const store = createStore({
    state () {
        return {
            lists: [],
            isMobile: false,
            needLogin: false,
        }
    },
    mutations: {
        setLists(state, lists) {
            state.lists = lists;
        },
        activateList(state, list) {
            state.lists = state.lists.map(i => {
                if (i.name === list.name && i.id === 0) {
                    i = list;
                }
                return i;
            });
        },
        setIsMobile(state, isMobile) {
            state.isMobile = isMobile;
        },
        needLogin(state, needLogin) {
            state.needLogin = needLogin;
        }
    }
})
app.use(store)

//use api plugin
app.use(api)

app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')

export default app