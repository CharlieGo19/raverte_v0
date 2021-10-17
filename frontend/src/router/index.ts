import { createRouter, createMemoryHistory, RouteRecordRaw } from 'vue-router'
import Welcome from "../login/Welcome.vue"
import MainStage from "../mainstage/MainStage.vue"



const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "welcome",
        component:  Welcome
        
    },
    {
      path: "/mainstage",
      name: "mainstage",
      component: MainStage
    }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

export default router

