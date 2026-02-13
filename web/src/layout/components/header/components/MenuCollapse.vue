<template>
  <n-icon size="20" cursor-pointer @click="appStore.switchCollapsed">
    <Icon icon="mdi:format-indent-increase" v-if="appStore.collapsed" />
    <Icon icon="mdi:format-indent-decrease" v-else />
  </n-icon>
</template>

<script setup>
import { useAppStore } from '~/src/store'
import { Icon } from '@iconify/vue'
const appStore = useAppStore()
</script>
