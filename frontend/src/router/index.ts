import { createRouter, createMemoryHistory, RouteRecordRaw } from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'

const routes: Array<RouteRecordRaw> = [

]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router
