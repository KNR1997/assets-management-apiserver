import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { useUserStore } from '@/store'
import { userClient } from './client/user'
import type { User, UserQueryOptions } from '@/types'
import { API_ENDPOINTS } from './client/api-endpoints'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'

export const useUpdateProfile = () => {
  const queryClient = useQueryClient()
  const message = useMessage()
  const userStore = useUserStore()

  return useMutation({
    mutationFn: userClient.update,
    onSuccess: async () => {
      message.success('Updated successfully')
      // fetch user detail (username, email)
      const me = await fetchMe()
      // store details global state
      userStore.setUserInfo(me.data)
      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.PROFILE],
      })
    },
  })
}

export const fetchMe = () => {
  return userClient.me()
}

export const useUsersQuery = (options: Partial<UserQueryOptions>) => {
  const { data, error, isPending } = useQuery<User[], Error>({
    queryKey: [API_ENDPOINTS.USERS, options],
    queryFn: () => userClient.all(options as UserQueryOptions),
  })
  // @ts-ignore
  const users = computed<User[]>(() => data.value ?? []) // todo -> fix
  return {
    users,
    error,
    loading: isPending,
  }
}

export const useCreateUserMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: userClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.USERS],
      })
    },
    onError: (error: Error) => {
      console.error('Create user failed:', error)
    },
  })
}

export const useUpdateUserMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: userClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.USERS],
      })
    },
  })
}
