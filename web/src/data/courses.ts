import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { courseClient } from './client/course'
import { API_ENDPOINTS } from './client/api-endpoints'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import type { Course, CourseListResponse, CourseQueryOptions } from '@/types'

export const useCoursesQuery = (options: Partial<CourseQueryOptions>) => {
  const { data, error, isPending } = useQuery<CourseListResponse, Error>({
    queryKey: [API_ENDPOINTS.COURSES, options],
    queryFn: () => courseClient.paginated(options as CourseQueryOptions),
  })
  // @ts-ignore
  const courses = computed<Course[]>(() => data.value?.data.list ?? []) // todo -> fix
  return {
    courses,
    error,
    loading: isPending,
  }
}

export const useCreateCourseMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: courseClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.COURSES],
      })
    },
    onError: (error: Error) => {
      console.error('Create course failed:', error)
    },
  })
}

export const useUpdateCourseMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: courseClient.update,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.COURSES],
      })
    },
  })
}

export const useDeleteCourseMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: courseClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.COURSES],
      })
    },
  })
}
