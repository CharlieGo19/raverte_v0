import { createRouter, createMemoryHistory, RouteRecordRaw } from 'vue-router'
import Welcome from '../login/Welcome.vue'

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "welcome",
        component:  Welcome
    },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router
