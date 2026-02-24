import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Manufacturer, ManufacturerQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import { manufacturerClient } from './client/manufacturer'

export const useManufacturersQuery = (options: Partial<ManufacturerQueryOptions>) => {
  const { data, error, isPending } = useQuery<Manufacturer[], Error>({
    queryKey: [API_ENDPOINTS.MANUFACTURERS, options],
    queryFn: () => manufacturerClient.all(options as ManufacturerQueryOptions),
  })
  // @ts-ignore
  const manufacturers = computed<Manufacturer[]>(() => data.value ?? []) // todo -> fix
  return {
    manufacturers,
    error,
    loading: isPending,
  }
}

export const useCreateManufacturerMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: manufacturerClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.MANUFACTURERS],
      })
    },
    onError: (error: Error) => {
      console.error('Create Manufacturer failed:', error)
    },
  })
}

export const useUpdateManufacturerMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: manufacturerClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.MANUFACTURERS],
      })
    },
  })
}

export const useDeleteManufacturerMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: manufacturerClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.MANUFACTURERS],
      })
    },
  })
}
