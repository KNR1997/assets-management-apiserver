import type {
  Course,
  CourseCreateInput,
  CoursePaginator,
  CourseQueryOptions,
  QueryOptions,
} from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { crudFactory } from './curd-factory'
import { HttpClient } from './http-client'

export const courseClient = {
  ...crudFactory<Course, QueryOptions, CourseCreateInput>(API_ENDPOINTS.COURSES),
  paginated: ({ name, ...params }: Partial<CourseQueryOptions>) => {
    return HttpClient.get<CoursePaginator>(API_ENDPOINTS.COURSES, {
      searchJoin: 'and',
      self,
      ...params,
      // search: HttpClient.formatSearchParams({ name }),
    })
  },
}
