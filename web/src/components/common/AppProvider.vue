<template>
  <n-config-provider
    wh-full
    :locale="enUS"
    :date-locale="enUSDate"
    :theme="appStore.isDark ? darkTheme : undefined"
    :theme-overrides="naiveThemeOverrides"
  >
    <n-loading-bar-provider>
      <n-dialog-provider>
        <n-notification-provider>
          <n-message-provider>
            <slot></slot>
            <NaiveProviderContent />
          </n-message-provider>
        </n-notification-provider>
      </n-dialog-provider>
    </n-loading-bar-provider>
  </n-config-provider>
</template>

<script setup>
import { defineComponent, h, onMounted } from 'vue'
import {
  zhCN,
  dateZhCN,
  darkTheme,
  NLoadingBarProvider,
  useLoadingBar,
  useDialog,
  useMessage,
  useNotification,
} from 'naive-ui'
import { useCssVar } from '@vueuse/core'
import { kebabCase } from 'lodash-es'
import { setupMessage, setupDialog } from '@/utils'
import { naiveThemeOverrides } from '~/settings'
import { useAppStore } from '@/store'

const appStore = useAppStore()

function setupCssVar() {
  const common = naiveThemeOverrides.common
  for (const key in common) {
    useCssVar(`--${kebabCase(key)}`, document.documentElement).value = common[key] || ''
    if (key === 'primaryColor') window.localStorage.setItem('__THEME_COLOR__', common[key] || '')
  }
}

// Use NLoadingBarProvider's methods instead of useLoadingBar()
function setupNaiveTools() {
  // Get references to the providers' API
  const messageApi = useMessage()
  const dialogApi = useDialog()
  const notificationApi = useNotification()

  // Setup message and dialog with proper error handling
  window.$message = setupMessage(messageApi)
  window.$dialog = setupDialog(dialogApi)
  window.$notification = notificationApi

  // Loading bar will be set up separately in the component
}

const NaiveProviderContent = defineComponent({
  setup() {
    // Setup CSS variables
    setupCssVar()
    setupNaiveTools()

    // Setup loading bar separately
    const loadingBar = useLoadingBar()

    // Expose loading bar to window on mounted
    onMounted(() => {
      window.$loadingBar = loadingBar
    })

    // Return the loading bar for template if needed
    return {
      loadingBar,
    }
  },
  render() {
    return h('div')
  },
})
</script>
