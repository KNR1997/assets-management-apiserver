export function useResize(el: any, cb: any) {
  const observer = new ResizeObserver((entries) => {
    cb(entries[0]?.contentRect)
  })
  observer.observe(el)
  return observer
}

const install = (app: any) => {
  let observer: any

  app.directive('resize', {
    mounted(el: any, binding: any) {
      observer = useResize(el, binding.value)
    },
    beforeUnmount() {
      observer?.disconnect()
    },
  })
}

useResize.install = install
