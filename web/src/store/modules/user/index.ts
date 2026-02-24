import { defineStore } from 'pinia'
import { resetRouter } from '@/router'
import { toLogin } from '@/utils/auth/auth'
import { removeToken } from '@/utils/auth/token'

export const useUserStore = defineStore('user', {
  state: () => ({
    userInfo: null as any,
  }),
  getters: {
    userId: (state) => state.userInfo?.id ?? null,
    username: (state) => state.userInfo?.username ?? '',
    email: (state) => state.userInfo?.email ?? '',
    avatar: (state) => state.userInfo?.avatar ?? '',
    role: (state) => state.userInfo?.role ?? 'user',
    // isSuperUser() {
    //   return this.userInfo?.is_superuser
    // },
    // isActive() {
    //   return this.userInfo?.is_active
    // },
  },
  actions: {
    async logout() {
      removeToken()
      resetRouter()
      this.$reset()
      toLogin()
      //   const { resetTags } = useTagsStore()
      //   const { resetPermission } = usePermissionStore()
      //   resetTags()
      //   resetPermission()
    },
    setUserInfo(userInfo = {}) {
      this.userInfo = { ...this.userInfo, ...userInfo }
    },
  },
})
