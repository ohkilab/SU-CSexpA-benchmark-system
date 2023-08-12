import { createRouter, createWebHistory } from "vue-router";
import Benchmark from "./pages/Benchmark.vue";
import Ranking from "./pages/Ranking.vue";
import Submissions from "./pages/Submissions.vue";
import Login from "./pages/Login.vue";
import Loading from "./pages/Loading.vue";
import Contests from "./pages/Contests.vue";
import Admin from "./pages/Admin.vue";
import EditContests from "./pages/admin/EditContests.vue";
import EditUsers from "./pages/admin/EditUsers.vue";

const routes = [
  {
    path: "/",
    name: "index",
    component: Loading,
  },
  {
    path: "/login",
    component: Login,
  },
  {
    path: "/ranking",
    component: Ranking,
  },
  {
    path: "/submissions/",
    component: Submissions,
  },
  {
    path: "/submissions/:id",
    component: Submissions,
  },
  {
    path: "/benchmark",
    component: Benchmark,
  },
  {
    path: "/contests",
    name: "contests",
    component: Contests,
  },
  {
    path: "/admin",
    name: 'admin',
    component: Admin,
    children: [
      {
        path: "contests",
        name: 'admin-contests',
        component: EditContests,
      },
      {
        path: "users",
        name: 'admin-users',
        component: EditUsers,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
