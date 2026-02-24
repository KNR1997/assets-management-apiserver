<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import { renderIcon } from '@/utils'
import { useRouter } from 'vue-router'
import { useDialog, useMessage } from 'naive-ui'
import { useUserStore } from '@/store/modules/user'

const router = useRouter()
const dialog = useDialog()
const message = useMessage()
const { t } = useI18n({ useScope: 'global' })
const userStore = useUserStore()

const options = [
  {
    label: 'Profile',
    key: 'profile',
    icon: renderIcon('mdi-account-arrow-right-outline', { size: '14px' }),
  },
  {
    label: 'Logout',
    key: 'logout',
    icon: renderIcon('mdi:exit-to-app', { size: '14px' }),
  },
]

function handleSelect(key) {
  if (key === 'profile') {
    router.push('/profile')
  } else if (key === 'logout') {
    dialog.warning({
      title: 'Confirm',
      content: 'Are you sure?',
      positiveText: 'Sure',
      negativeText: 'Not Sure',
      draggable: true,
      onPositiveClick: () => {
        userStore.logout()
        message.success('Logout!')
      },
      onNegativeClick: () => {
        message.error('Not Sure')
      },
    })
  }
}
</script>

<template>
  <n-dropdown :options="options" @select="handleSelect">
    <div flex cursor-pointer items-center>
      <img src="/user.svg" class="mr10 h-35 w-35 rounded-full" />
      <span>admin</span>
    </div>
  </n-dropdown>
</template>
