import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import GamePlaying from "@/views/GamePlaying.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Index,
  },
  {
    path: "/game-playing/:id",
    component: GamePlaying,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
