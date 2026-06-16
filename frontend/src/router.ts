import { createRouter, createWebHistory } from "vue-router";
import Benchmark from "./pages/Benchmark.vue";
import Ranking from "./pages/Ranking.vue";
import Submissions from "./pages/Submissions.vue";
import Login from "./pages/Login.vue";
import Loading from "./pages/Loading.vue";
import Contests from "./pages/Contests.vue";
import Admin from "./pages/Admin.vue";
import EditContests from "./pages/admin/EditContests.vue";
import EditGroups from "./pages/admin/EditGroups.vue";
import { useStateStore } from "./stores/state";
import { Role } from "proto-gen-web/services/backend/resources";

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
    name: "admin",
    component: Admin,
    meta: { requiresAdmin: true },
    children: [
      {
        path: "contests",
        name: "admin-contests",
        component: EditContests,
        meta: { requiresAdmin: true },
      },
      {
        path: "groups",
        name: "admin-groups",
        component: EditGroups,
        meta: { requiresAdmin: true },
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to) => {
  if (!to.matched.some((record) => record.meta.requiresAdmin)) {
    return true;
  }

  const state = useStateStore();
  if (!state.token) {
    return "/login";
  }
  if (state.role !== Role.ADMIN) {
    return "/contests";
  }
  return true;
});

export default router;
