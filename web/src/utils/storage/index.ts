import { createStorage } from './storage'

const prefixKey = ''

export const createLocalStorage = function (option = {}) {
  return createStorage({
    // @ts-ignore
    prefixKey: option.prefixKey || '',
    storage: localStorage,
  })
}

export const createSessionStorage = function (option = {}) {
  return createStorage({
    // @ts-ignore
    prefixKey: option.prefixKey || '',
    storage: sessionStorage,
  })
}

export const lStorage = createLocalStorage({ prefixKey })

export const sStorage = createSessionStorage({ prefixKey })
