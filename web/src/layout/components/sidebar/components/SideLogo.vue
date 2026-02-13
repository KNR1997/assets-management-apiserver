<template>
  <router-link h-60 f-c-c to="/">
    <!-- <icon-custom-logo text-36 color-primary></icon-custom-logo> -->
    <Icon text-24 icon="logos:vue" />
    <h2
      v-show="!appStore.collapsed"
      ml-2
      mr-8
      max-w-150
      flex-shrink-0
      text-16
      font-bold
      color-primary
    >
      {{ title }}
    </h2>
  </router-link>
</template>

<script setup>
import { useAppStore } from '@/store'
const title = import.meta.env.VITE_TITLE
import { Icon } from '@iconify/vue'

const appStore = useAppStore()
</script>
