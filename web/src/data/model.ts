import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { modelClient } from './client/model'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Model, ModelQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'

export const useModelsQuery = (options: Partial<ModelQueryOptions>) => {
  const { data, error, isPending } = useQuery<Model[], Error>({
    queryKey: [API_ENDPOINTS.MODELS, options],
    queryFn: () => modelClient.all(options as ModelQueryOptions),
  })
  // @ts-ignore
  const models = computed<Model[]>(() => data.value ?? []) // todo -> fix
  return {
    models,
    error,
    loading: isPending,
  }
}

export const useCreateModelMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: modelClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.MODELS],
      })
    },
    onError: (error: Error) => {
      console.error('Create Model failed:', error)
    },
  })
}

export const useUpdateModelMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: modelClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.MODELS],
      })
    },
  })
}

export const useDeleteModelMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: modelClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.MODELS],
      })
    },
  })
}
