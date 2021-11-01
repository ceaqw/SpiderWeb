import { createWebHistory, createRouter } from 'vue-router'
import Layout from '@/layout'
import { routerAuth } from '@/utils/auth'

const history = createWebHistory()

const router = createRouter({
    history,
    routes: [{
            path: '/',
            redirect: '/home'
        },
        {
            path: '/home',
            component: Layout,
            redirect: '/home/dashboard',
            children: [
                {
                    path: 'dashboard',
                    component: () => import ('@/views/home'),
                },
                {
                    path: 'project',
                    component: () => import ('@/views/project'),
                },
            ]
        },
        {
            path: '/trend',
            component: Layout,
            children: [
                {
                    path: 'kpi',
                    component: () => import ('@/views/trend/kpi'),
                },
                {
                    path: 'fail',
                    component: () => import ('@/views/trend/fail'),
                },
            ]
        },
        {
            path: '/member',
            component: Layout,
            children: [
                {
                    path: 'list',
                    beforeEnter: routerAuth,
                    component: () => import ('@/views/member/list'),
                },
                {
                    path: 'create',
                    component: () => import ('@/views/member/create'),
                },
                {
                    path: 'update',
                    component: () => import ('@/views/member/update'),
                },
                {
                    path: 'project',
                    component: () => import ('@/views/member/project'),
                },
            ]
        },
        {
            path: '/login',
            name: 'login',
            component: () => import ('@/views/login'),
            meta: { title: '登录' }
        },
        {
            path: '/register',
            name: 'register',
            component: () => import ('@/views/register'),
            meta: { title: '注册' }
        },
        {
            path: '/401',
            component: () => import ('@/views/error/401'),
            meta: { title: '权限不足' }
        },
        {
            path: '/403',
            name: '403',
            component: () => import ('@/views/error/403'),
            meta: { title: '权限不足' }
        },
        {
            path: '/500',
            name: '500',
            component: () => import ('@/views/error/500'),
            meta: { title: '服务器错误' }
        },
        {
            path: '/404',
            name: '404',
            component: () => import ('@/views/error/404'),
            meta: { title: '页面访问错误' }
        }
    ]
})

export default router