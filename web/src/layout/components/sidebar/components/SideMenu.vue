<script setup lang="ts">
import { computed } from 'vue'
import { renderIcon } from '@/utils'
import { useRoute, useRouter } from 'vue-router'

const router = useRouter()
const curRoute = useRoute()

type MenuOption = {
  key: string
  path?: string
  children?: MenuOption[]
}

function findActiveKey(options: MenuOption[], path: string): string | undefined {
  for (const item of options) {
    // 1️⃣ check children first (deepest match wins)
    if (item.children) {
      const childKey = findActiveKey(item.children, path)
      if (childKey) return childKey
    }

    // 2️⃣ check self
    if (item.path && path.startsWith(item.path)) {
      return item.key
    }
  }
}

const activeKey = computed(() => findActiveKey(menuOptions, curRoute.path))

const menuOptions = [
  {
    label: 'Workbench',
    key: 'workbench',
    icon: renderIcon('icon-park-outline:workbench', { size: 18 }),
    path: '/workbench',
  },
  {
    label: 'Asset',
    key: 'asset',
    icon: renderIcon('qlementine-icons:items-tree-16', { size: 18 }),
    path: '/assets',
  },
  {
    label: 'People',
    key: 'people',
    icon: renderIcon('ic:baseline-people', { size: 18 }),
    children: [
      { label: 'All people', key: 'profile', path: '/user' },
      // { label: "Security", key: "security", path: '/security'  },
    ],
  },
  // {
  //   label: 'System',
  //   key: 'system',
  //   icon: renderIcon('mdi-account-off', { size: 18 }),
  //   path: '/system',
  // },
  {
    label: 'Settings',
    key: 'settings',
    icon: renderIcon('ic:baseline-settings', { size: 18 }),
    children: [
      { label: 'Categories', key: 'categories', path: '/settings/categories' },
      { label: 'Manufacturers', key: 'manufacturers', path: '/settings/manufacturers' },
      { label: 'Suppliers', key: 'suppliers', path: '/settings/suppliers' },
      { label: 'Departments', key: 'departments', path: '/settings/departments' },
      // { label: "Security", key: "security", path: '/security'  },
    ],
  },
  // {
  //   label: 'Profile',
  //   key: 'profile',
  //   icon: renderIcon('mdi-account-off', { size: 18 }),
  //   path: '/profile',
  // },
  // {
  //   label: "Settings",
  //   key: "settings",
  //   children: [
  //     { label: "Profile", key: "profile", path: '/profile' },
  //     { label: "Security", key: "security", path: '/security'  },
  //   ],
  // },
]

function handleMenuSelect(key: string, item: any) {
  if (item.path) {
    router.push(item.path)
  }
}
</script>

<style lang="scss">
.side-menu:not(.n-menu--collapsed) {
  .n-menu-item-content {
    &::before {
      left: 5px;
      right: 5px;
    }
    &.n-menu-item-content--selected,
    &:hover {
      &::before {
        border-left: 4px solid var(--primary-color);
      }
    }
  }
}
</style>

<template>
  <n-menu
    ref="menu"
    class="side-menu"
    accordion
    :indent="18"
    :collapsed-icon-size="22"
    :collapsed-width="64"
    :options="menuOptions"
    :value="activeKey"
    @update:value="handleMenuSelect"
  />
</template>
