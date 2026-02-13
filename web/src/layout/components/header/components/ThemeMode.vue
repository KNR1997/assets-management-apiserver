<template>
  <n-icon mr-20 cursor-pointer size="18" @click="toggleDark">
    <Icon icon="mdi-moon-waning-crescent" v-if="isDark" />
    <Icon icon="mdi-white-balance-sunny" v-else />
  </n-icon>
</template>

<script setup>
import { useAppStore } from '~/src/store'
import { useDark, useToggle } from '@vueuse/core'
import { Icon } from '@iconify/vue'

const appStore = useAppStore()
const isDark = useDark()
const toggleDark = () => {
  appStore.toggleDark()
  useToggle(isDark)()
}
</script>
