import { computed } from 'vue'
import { useMessage } from 'naive-ui'
import { quizClient } from './client/quiz'
import { API_ENDPOINTS } from './client/api-endpoints'
import type { Quiz, QuizListResponse, QuizQueryOptions } from '@/types'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'

export const useQuizzesQuery = (options: Partial<QuizQueryOptions>) => {
  const { data, error, isPending } = useQuery<QuizListResponse, Error>({
    queryKey: [API_ENDPOINTS.QUIZZES, options],
    queryFn: () => quizClient.paginated(options as QuizQueryOptions),
  })
  // @ts-ignore
  const quizzes = computed<Quiz[]>(() => data.value?.data.list ?? []) // todo -> fix
  return {
    quizzes,
    error,
    loading: isPending,
  }
}

export const useCreateQuizMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: quizClient.create,

    onSuccess: () => {
      message.success('Created successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.QUIZZES],
      })
    },
    onError: (error: Error) => {
      console.error('Create course failed:', error)
    },
  })
}

export const useUpdateQuizMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: quizClient.update,
    onSuccess: () => {
      message.success('Updated successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.QUIZZES],
      })
    },
  })
}

export const useDeleteQuizMutation = () => {
  const queryClient = useQueryClient()
  const message = useMessage()

  return useMutation({
    mutationFn: quizClient.delete,
    onSuccess: () => {
      message.success('Deleted successfully')

      queryClient.invalidateQueries({
        queryKey: [API_ENDPOINTS.QUIZZES],
      })
    },
  })
}
