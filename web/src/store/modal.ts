import { defineStore } from 'pinia'
import { shallowRef } from 'vue'

export const useModalStore = defineStore('modal', {
  state: () => ({
    visible: false,
    title: '',
    component: shallowRef<any>(null),
    props: {} as Record<string, any>,
  }),

  actions: {
    open(component: any, options?: { title?: string; props?: any }) {
      this.component = component
      this.title = options?.title ?? ''
      this.props = options?.props ?? {}
      this.visible = true
    },

    close() {
      this.visible = false
      this.component = null
      this.props = {}
    },
  },
})
