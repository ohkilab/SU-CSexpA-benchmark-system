import Benchmark from './pages/Benchmark.vue'
import Ranking from './pages/Ranking.vue'
import Submissions from './pages/Submissions.vue'
import Login from './pages/Login.vue'
import Loading from './pages/Loading.vue'
import Contests from './pages/Contests.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'index',
    component: Loading
  },
  {
    path: '/login',
    component: Login
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
  },
  {
    path: '/contests',
    name: 'contests',
    component: Contests
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
