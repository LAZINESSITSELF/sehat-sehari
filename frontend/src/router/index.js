import { createRouter, createWebHistory } from "vue-router";
import Login from "../pages/Login.vue";
import Register from "../pages/Register.vue";
import ProfileForm from "../pages/ProfileForm.vue";
import Dashboard from "../pages/Dashboard.vue";
import SehatMakan from "../pages/SehatMakan.vue";
import SehatAktivitas from "../pages/SehatAktivitas.vue";
import SehatHidup from "../pages/SehatHidup.vue";

const routes = [
    { path: "/", redirect: "/login" },
    { path: "/login", component: Login },
    { path: "/register", component: Register },
    {
        path: "/profile",
        component: ProfileForm,
        meta: { requiresAuth: true },
    },
    {
        path: "/dashboard",
        component: Dashboard,
        meta: { requiresAuth: true },
        children: [
            { path: "makan", component: SehatMakan },
            { path: "aktivitas", component: SehatAktivitas },
            { path: "hidup", component: SehatHidup },
            { path: "", redirect: "makan" },
        ],
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// Guard: jika route membutuhkan auth tapi token tidak ada → redirect ke login
router.beforeEach((to, from, next) => {
    const token = localStorage.getItem("token");
    if (to.meta.requiresAuth && !token) {
        return next("/login");
    }
    next();
});

export default router;
