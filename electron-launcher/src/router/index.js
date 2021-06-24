import Vue from "vue";
import VueRouter from "vue-router";
import Setup from "../views/Setup.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Setup",
    component: Setup,
  },  
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
