import { createRouter, createWebHistory } from 'vue-router'
import AssetView from '../views/AssetView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'home',
            component: AssetView
        }
    ]
})

export default router