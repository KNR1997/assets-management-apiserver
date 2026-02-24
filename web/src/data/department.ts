import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { departmentClient } from './client/department'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Department, DepartmentQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'

export const useDepartmentsQuery = (options: Partial<DepartmentQueryOptions>) => {
  const { data, error, isPending } = useQuery<Department[], Error>({
    queryKey: [API_ENDPOINTS.DEPARTMENTS, options],
    queryFn: () => departmentClient.all(options as DepartmentQueryOptions),
  })
  // @ts-ignore
  const departments = computed<Department[]>(() => data.value ?? []) // todo -> fix
  return {
    departments,
    error,
    loading: isPending,
  }
}

export const useCreateDepartmentMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: departmentClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.DEPARTMENTS],
      })
    },
    onError: (error: Error) => {
      console.error('Create Department failed:', error)
    },
  })
}

export const useUpdateDepartmentMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: departmentClient.patch,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.DEPARTMENTS],
      })
    },
  })
}

export const useDeleteDepartmentMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: departmentClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.DEPARTMENTS],
      })
    },
  })
}
