import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/task",
      name: "task",
      component: () => import("../views/TaskView.vue"),
    },
    {
      path: "/anime",
      name: "anime",
      component: () => import("../views/AnimeSeries.vue"),
    },
    {
      path: "/anime/edit/:id",
      name: "anime-edit",
      component: () => import("../views/AnimeEdit.vue"),
      props: true,
    },
  ],
});

export default router;
