import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/extraTools',
    name: 'extraTools',
    component: () => import(/* webpackChunkName: "extraTools" */ '../views/ExtraToolsView.vue')
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
    component: () => import(/* webpackChunkName: "pdfMerge" */ '../views/PDFMergeView.vue')
  },
  {
    path: '/pdfSplit',
    name: 'pdfSplit',
    component: () => import(/* webpackChunkName: "pdfSplit" */ '../views/PDFSplitView.vue')
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
