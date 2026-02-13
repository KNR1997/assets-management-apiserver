<script setup>
import AppProvider from './components/common/AppProvider.vue'
import GlobalModalHost from './components/common/GlobalModalHost.vue'
</script>

<template>
  <AppProvider>
    <router-view v-slot="{ Component }">
      <component :is="Component" />
    </router-view>

    <!-- GLOBAL MODAL -->
    <GlobalModalHost />
  </AppProvider>
</template>
