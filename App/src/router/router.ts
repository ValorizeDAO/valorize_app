import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "../views/Dashboard.vue";
import Login from "../views/Login.vue";
import store from "../vuex/vuex";

const routes = [
  {
    path: "/",
    name: "Dashboard",
    component: Dashboard,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
];
const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const publicRoutes = ["Login", "Register"];
  let { name } = to;
  name = name?.toString() || "";
  const isAuthenticated = store.state.authenticated;
  if (!publicRoutes.includes(name) && !isAuthenticated) next({ name: "Login" });
  else {
    console.log("going to next");
    next();
  }
});

export default router;
