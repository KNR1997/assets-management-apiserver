import type {
  Category,
  CategoryCreateInput,
  CategoryPaginator,
  CategoryQueryOptions,
  QueryOptions,
} from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { crudFactory } from './curd-factory'
import { HttpClient } from './http-client'

export const categoryClient = {
  ...crudFactory<Category, QueryOptions, CategoryCreateInput>(API_ENDPOINTS.CATEGORIES),
  paginated: ({ name, ...params }: Partial<CategoryQueryOptions>) => {
    return HttpClient.get<CategoryPaginator>(API_ENDPOINTS.CATEGORIES, {
      searchJoin: 'and',
      self,
      ...params,
      // search: HttpClient.formatSearchParams({ name }),
    })
  },
}
