<template>
  <n-icon mr-20 size="18" style="cursor: pointer" @click="handleLinkClick">
    <Icon icon="mdi:github" />
  </n-icon>
</template>

<script setup>
import { Icon } from '@iconify/vue'

function handleLinkClick() {
  window.open('https://github.com/mizhexiaoxiao/vue-fastapi-admin')
}
</script>
