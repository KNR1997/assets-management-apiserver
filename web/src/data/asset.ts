import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { assetClient } from './client/asset'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Asset, AssetQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'

export const useAssetsQuery = (options: Partial<AssetQueryOptions>) => {
  const { data, error, isPending } = useQuery<Asset[], Error>({
    queryKey: [API_ENDPOINTS.ASSETS, options],
    queryFn: () => assetClient.all(options as AssetQueryOptions),
  })
  // @ts-ignore
  const assets = computed<Asset[]>(() => data.value ?? []) // todo -> fix
  return {
    assets,
    error,
    loading: isPending,
  }
}

export const useCreateAssetMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: assetClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.ASSETS],
      })
    },
    onError: (error: Error) => {
      console.error('Create asset failed:', error)
    },
  })
}

export const useUpdateAssetMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: assetClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.ASSETS],
      })
    },
  })
}

export const useDeleteAssetMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: assetClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.ASSETS],
      })
    },
  })
}
