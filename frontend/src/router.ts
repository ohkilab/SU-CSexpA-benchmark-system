import Benchmark from './pages/Benchmark.vue'
import Ranking from './pages/Ranking.vue'
import Submissions from './pages/Submissions.vue'
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
    path: '/submissions',
    component: Submissions
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
