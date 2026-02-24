import { useMessage } from 'naive-ui'
import { categoryClient } from './client/category'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Category, CategoryPaginator, CategoryQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import { computedAsync } from '@vueuse/core'
import { mapPaginatorData } from '@/utils/data-mappers'

export const useCategoriesQuery = (options: Partial<CategoryQueryOptions>) => {
  const { data, error, isPending } = useQuery<CategoryPaginator, Error>({
    queryKey: [API_ENDPOINTS.CATEGORIES, options],
    queryFn: ({ queryKey, pageParam }) =>
      categoryClient.paginated(Object.assign({}, queryKey[1], pageParam)),
  })
  const categories = computedAsync<Category[]>(() => data.value?.rows ?? []) // todo -> fix
  return {
    categories,
    paginatorInfo: mapPaginatorData(data.value),
    error,
    loading: isPending,
  }
}

export const useCreateCategoryMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: categoryClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.CATEGORIES],
      })
    },
    onError: (error: Error) => {
      console.error('Create category failed:', error)
    },
  })
}

export const useUpdateCategoryMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: categoryClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.CATEGORIES],
      })
    },
  })
}

export const useDeleteCategoryMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: categoryClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.CATEGORIES],
      })
    },
  })
}
