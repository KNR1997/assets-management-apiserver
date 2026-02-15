import { useMessage } from 'naive-ui'
import { userClient } from './client/user'
import { API_ENDPOINTS } from './client/api-endpoints'
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { useUserStore } from '@/store'

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
