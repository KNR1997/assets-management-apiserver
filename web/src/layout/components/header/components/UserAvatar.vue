<template>
  <n-dropdown :options="options" @select="handleSelect">
    <div flex cursor-pointer items-center>
      <img src="/user.svg" class="mr10 h-35 w-35 rounded-full" />
      <span>admin</span>
    </div>
  </n-dropdown>
</template>

<script setup>
import { renderIcon } from '@/utils'
import { useRouter } from 'vue-router'
const router = useRouter()

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
    // $dialog.confirm({
    //   title: t('header.label_logout_dialog_title'),
    //   type: 'warning',
    //   content: t('header.text_logout_confirm'),
    //   confirm() {
    //     userStore.logout()
    //     $message.success(t('header.text_logout_success'))
    //   },
    // })
  }
}
</script>
