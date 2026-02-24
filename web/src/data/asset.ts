import { computed } from 'vue'
import { router } from '@/router'
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

export const useAssetQuery = ({ id }: { id: string }) => {
  const { data, error, isPending } = useQuery<Asset, Error>({
    queryKey: [API_ENDPOINTS.ASSETS, id],
    queryFn: () => assetClient.get({ id }),
  })
  // @ts-ignore
  const asset = computed<Asset>(() => data.value?.data ?? {}) // todo -> fix
  return {
    asset,
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
      router.push('/assets')

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
      router.push('/assets')

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

export const useCheckoutAssetMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: assetClient.checkout,
    onSuccess: () => {
      message.success('Checkout successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.ASSETS],
      })
    },
  })
}

export const useCheckinAssetMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: assetClient.checkin,
    onSuccess: () => {
      message.success('Checkin successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.ASSETS],
      })
    },
  })
}
