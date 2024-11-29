import { createRouter, createWebHistory } from '@ionic/vue-router';
import HomePage from '../views/HomePage.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/login',
    component: HomePage,
    children: [
      {
        path: 'login',
        name: 'Login',
        component: () => import('../views/auth/login.vue'),
      },
      {
        path: 'books',
        name: 'BookList',
        component: () => import('../views/book/book_list.vue'),
      },
      {
        path: 'add-book/:id?',
        name: 'BookAdd',
        component: () => import('../views/book/add_book.vue'),
      },
      {
        path: 'staffs',
        name: 'StaffList',
        component: () => import('../views/staff/staff_list.vue'),
      },
      {
        path: 'add-staff/:id?',
        name: 'StaffAdd',
        component: () => import('../views/staff/add_staff.vue'),
      },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
