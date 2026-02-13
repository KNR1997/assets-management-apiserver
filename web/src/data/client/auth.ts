import type { AuthResponse, LoginInput } from '@/types'
import { HttpClient } from './http-client'
import { API_ENDPOINTS } from './api-endpoints'

export const authClient = {
  login: (variables: LoginInput) => {
    return HttpClient.post<AuthResponse>(API_ENDPOINTS.TOKEN, variables)
  },
}
