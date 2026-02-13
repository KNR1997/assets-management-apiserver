import { lStorage } from '@/utils'

const TOKEN_CODE = 'access_token'

export function getToken() {
  return lStorage.get(TOKEN_CODE)
}

export function setToken(token: string) {
  lStorage.set(TOKEN_CODE, token, null)
}

export function removeToken() {
  lStorage.remove(TOKEN_CODE)
}
