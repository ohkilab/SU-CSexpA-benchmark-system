import Benchmark from './pages/Benchmark.vue'
import Ranking from './pages/Ranking.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/ranking'
  },
  {
    path: '/ranking',
    component: Ranking
  },
  {
    path: '/benchmark',
    component: Benchmark
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
