import { useMutation } from '@tanstack/vue-query'
import { authClient } from './client/auth'
import type { LoginInput, AuthResponse } from '@/types'

export function useLogin() {
  return useMutation<AuthResponse, Error, LoginInput>({
    mutationFn: authClient.login,
  })
}
