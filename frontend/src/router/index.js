import { createRouter, createWebHashHistory } from 'vue-router';
import Login from '../views/Login.vue';
import ManagerDashboard from '../components/ManagerDashboard.vue';
import CourierDashboard from '../components/CourierDashboard.vue';

const routes = [
    { path: '/', name: 'Login', component: Login },
    { path: '/manager', name: 'ManagerDashboard', component: ManagerDashboard },
    { path: '/courier', name: 'CourierDashboard', component: CourierDashboard },
];

const router = createRouter({
    history: createWebHashHistory(), // Menggunakan hash history
    routes,
});

export default router;
