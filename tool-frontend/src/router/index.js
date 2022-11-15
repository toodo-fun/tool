import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/windows',
    name: 'windows',
    component: () => import(/* webpackChunkName: "windows" */ '../views/WindowsView.vue')
  },
  {
    path: '/downloadSpeedUp',
    name: 'downloadSpeedUp',
    component: () => import(/* webpackChunkName: "downloadSpeedUp" */ '../views/DownloadSpeedUpView.vue')
  },
  {
    path: '/pdfMerge',
    name: 'pdfMerge',
    component: () => import(/* webpackChunkName: "pdf" */ '../views/PDFMergeView.vue')
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
