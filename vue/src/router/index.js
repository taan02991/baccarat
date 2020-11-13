import Vue from "vue";
import VueRouter from "vue-router";

const Index = () => import("../views/Index.vue");
const GamePlaying = () => import("@/views/GamePlaying.vue");

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Index
  },
  {
    path: "/game/:id",
    name: GamePlaying,
    component: GamePlaying
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
