import { isNullOrUndef } from '@/utils'

class Storage {
  constructor(option: any) {
    // @ts-ignore
    this.storage = option.storage
    // @ts-ignore
    this.prefixKey = option.prefixKey
  }

  getKey(key: any) {
    // @ts-ignore
    return `${this.prefixKey}${key}`.toUpperCase()
  }

  set(key: any, value: any, expire: any) {
    const stringData = JSON.stringify({
      value,
      time: Date.now(),
      expire: !isNullOrUndef(expire) ? new Date().getTime() + expire * 1000 : null,
    })
    // @ts-ignore
    this.storage.setItem(this.getKey(key), stringData)
  }

  get(key: any) {
    // @ts-ignore
    const { value } = this.getItem(key, {})
    return value
  }

  getItem(key: any, def = null) {
    // @ts-ignore
    const val = this.storage.getItem(this.getKey(key))
    if (!val) return def
    try {
      const data = JSON.parse(val)
      const { value, time, expire } = data
      if (isNullOrUndef(expire) || expire > new Date().getTime()) {
        return { value, time }
      }
      this.remove(key)
      return def
    } catch (error) {
      this.remove(key)
      return def
    }
  }

  remove(key: any) {
    // @ts-ignore
    this.storage.removeItem(this.getKey(key))
  }

  clear() {
    // @ts-ignore
    this.storage.clear()
  }
}

export function createStorage({ prefixKey = '', storage = sessionStorage }) {
  return new Storage({ prefixKey, storage })
}
