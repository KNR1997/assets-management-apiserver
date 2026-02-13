import '@/styles/reset.css'
import 'uno.css'
import '@/styles/global.scss'
import { setupRouter } from '@/router'
import { createApp } from 'vue'
import { setupStore } from '@/store'
// @ts-ignore
import i18n from '~/i18n'
// @ts-ignore
import { useResize } from '@/utils'
import App from './App.vue'
import naive from 'naive-ui' // This line is crucial
import { VueQueryPlugin } from '@tanstack/vue-query'

async function setupApp() {
  const app = createApp(App)

  setupStore(app)

  await setupRouter(app)
  app.use(useResize)
  app.use(i18n)
  app.use(VueQueryPlugin)
  app.use(naive)

  app.mount('#app')
}

setupApp()
