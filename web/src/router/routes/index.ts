const Layout = () => import('@/layout/index.vue')

export const basicRoutes = [
  {
    path: '/',
    redirect: '/workbench', // 默认跳转到首页
    meta: { order: 0 },
  },
  {
    name: 'Workbench-tsdsd',
    path: '/workbench',
    component: Layout,
    children: [
      {
        path: '',
        component: () => import('@/views/workbench/index.vue'),
        name: 'Workbench',
        meta: {
          title: 'Workbench',
          icon: 'icon-park-outline:workbench',
          affix: true,
        },
      },
    ],
    meta: { order: 1 },
  },
  {
    name: 'Login-tsdsd',
    path: '/login',
    component: () => import('@/views/login/index.vue'),
    isHidden: true,
  },
  {
    name: 'Assets-ddd',
    path: '/assets',
    component: Layout,
    children: [
      {
        name: 'Assets',
        path: '',
        component: () => import('@/views/asset/index.vue'),
        meta: {
          title: 'Assets',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
      {
        name: 'Asset Create',
        path: 'create',
        component: () => import('@/views/asset/create.vue'),
        meta: {
          title: 'Asset Create',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
      {
        name: 'Asset Edit',
        path: 'edit/:id',
        component: () => import('@/views/asset/edit.vue'),
        meta: {
          title: 'Asset Edit',
        },
      },
    ],
    meta: { order: 2 },
  },
  {
    name: 'System-dev',
    path: '/system',
    component: Layout,
    children: [
      {
        path: '',
        component: () => import('@/views/system/index.vue'),
        name: 'System',
        meta: {
          title: 'System Default',
          icon: 'icon-park-outline:workbench',
          affix: true,
        },
      },
    ],
    meta: { order: 5 },
  },
  {
    name: 'Profile-dev',
    path: '/profile',
    component: Layout,
    isHidden: true,
    children: [
      {
        path: '',
        component: () => import('@/views/profile/index.vue'),
        name: 'Profile',
        meta: {
          title: 'Profile',
          icon: 'user',
          affix: true,
        },
      },
    ],
    meta: { order: 99 },
  },
  {
    name: 'People-ddd',
    path: '/user',
    component: Layout,
    children: [
      {
        path: '',
        component: () => import('@/views/user/index.vue'),
        name: 'People',
        meta: {
          title: 'People',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
    ],
    meta: { order: 2 },
  },
  {
    name: 'Settings-tdd',
    path: '/settings',
    component: Layout,
    children: [
      {
        path: 'categories',
        component: () => import('@/views/settings/categories.vue'),
        name: 'Categories',
        meta: {
          title: 'Categories',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
      {
        path: 'manufacturers',
        component: () => import('@/views/settings/manufacturers.vue'),
        name: 'Manufacturers',
        meta: {
          title: 'Manufacturers',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
      {
        path: 'suppliers',
        component: () => import('@/views/settings/suppliers.vue'),
        name: 'Suppliers',
        meta: {
          title: 'Suppliers',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
      {
        path: 'departments',
        component: () => import('@/views/settings/departments.vue'),
        name: 'Departments',
        meta: {
          title: 'Departments',
          icon: 'fluent:quiz-20-regular',
          affix: true,
        },
      },
    ],
    meta: { order: 2 },
  },
]
