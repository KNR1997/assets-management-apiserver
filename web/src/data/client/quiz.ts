import type { QueryOptions, Quiz, QuizCreateInput, QuizPaginator, QuizQueryOptions } from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { crudFactory } from './curd-factory'
import { HttpClient } from './http-client'

export const quizClient = {
  ...crudFactory<Quiz, QueryOptions, QuizCreateInput>(API_ENDPOINTS.QUIZZES),
  paginated: ({ name, ...params }: Partial<QuizQueryOptions>) => {
    return HttpClient.get<QuizPaginator>(API_ENDPOINTS.QUIZZES, {
      searchJoin: 'and',
      self,
      ...params,
      // search: HttpClient.formatSearchParams({ name }),
    })
  },
}
