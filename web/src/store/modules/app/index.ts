import { defineStore } from 'pinia'
import { useDark } from '@vueuse/core'
import { lStorage } from '@/utils'
// import i18n from '~/i18n'

// const { locale } = i18n.global

const isDark = useDark()
export const useAppStore = defineStore('app', {
  state() {
    return {
      reloadFlag: true,
      collapsed: false,
      fullScreen: true,
      /** keepAlive路由的key，重新赋值可重置keepAlive */
      aliveKeys: {},
      isDark,
      locale: 'en',
    }
  },
  actions: {
    async reloadPage() {
      // @ts-ignore
      $loadingBar.start()
      this.reloadFlag = false
      // @ts-ignore
      await nextTick()
      this.reloadFlag = true

      setTimeout(() => {
        document.documentElement.scrollTo({ left: 0, top: 0 })
        // @ts-ignore
        $loadingBar.finish()
      }, 100)
    },
    switchCollapsed() {
      this.collapsed = !this.collapsed
    },
    setCollapsed(collapsed: any) {
      this.collapsed = collapsed
    },
    setFullScreen(fullScreen: any) {
      this.fullScreen = fullScreen
    },
    setAliveKeys(key: any, val: any) {
      // @ts-ignore
      this.aliveKeys[key] = val
    },
    /** 设置暗黑模式 */
    setDark(isDark: any) {
      this.isDark = isDark
    },
    /** 切换/关闭 暗黑模式 */
    toggleDark() {
      this.isDark = !this.isDark
    },
    setLocale(newLocale: any) {
      this.locale = newLocale
      // @ts-ignore
      locale.value = newLocale
      // @ts-ignore
      lStorage.set('locale', newLocale)
    },
  },
})
