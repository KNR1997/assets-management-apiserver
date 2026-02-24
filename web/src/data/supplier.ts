import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Supplier, SupplierQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import { supplierClient } from './client/supplier'

export const useSuppliersQuery = (options: Partial<SupplierQueryOptions>) => {
  const { data, error, isPending } = useQuery<Supplier[], Error>({
    queryKey: [API_ENDPOINTS.SUPPLIERS, options],
    queryFn: () => supplierClient.all(options as SupplierQueryOptions),
  })
  // @ts-ignore
  const suppliers = computed<Supplier[]>(() => data.value ?? []) // todo -> fix
  return {
    suppliers,
    error,
    loading: isPending,
  }
}

export const useCreateSupplierMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: supplierClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.SUPPLIERS],
      })
    },
    onError: (error: Error) => {
      console.error('Create Supplier failed:', error)
    },
  })
}

export const useUpdateSupplierMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: supplierClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.SUPPLIERS],
      })
    },
  })
}

export const useDeleteSupplierMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: supplierClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.SUPPLIERS],
      })
    },
  })
}
