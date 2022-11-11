import { createApp } from 'vue'
import { createStore } from 'vuex'
import { createRouter, createWebHistory } from 'vue-router'

import api from './plugins/api.js'

import css from './index.css'
import App from './App.vue'

let app = createApp(App)

// Use vue-router
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        { name: "app", path: '/:mode?/:listId?/', component: App},
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
        updateList(state, list) {
            //replace the list in the lists array
            state.lists = state.lists.map(i => i.id === list.id ? list : i);
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

//use fontawesome
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {fas} from '@fortawesome/free-solid-svg-icons'
library.add(fas)

//use api plugin
app.use(api)

app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')

export default app