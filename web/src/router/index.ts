import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import { basicRoutes } from './routes'

// import Home from "../pages/Home.vue";
// import About from "../pages/About.vue";

// const routes = [
//   {
//     path: "/",
//     component: Home,
//     meta: { layout: "Default" },
//   },
//   {
//     path: "/about",
//     component: About,
//     meta: { layout: "Default" },
//   },
// ];

// const router = createRouter({
//   history: createWebHashHistory(), // Tauri friendly
//   routes,
// });

// export default router;

const isHash = import.meta.env.VITE_USE_HASH === 'true'
export const router = createRouter({
  history: isHash ? createWebHashHistory('/') : createWebHistory('/'),
  routes: basicRoutes,
  scrollBehavior: () => ({ left: 0, top: 0 }),
})

export async function setupRouter(app: any) {
  // await addDynamicRoutes()
  // setupRouterGuard(router)
  app.use(router)
}
