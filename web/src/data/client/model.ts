import type {
  Model,
  ModelCreateInput,
  ModelPaginator,
  ModelQueryOptions,
  QueryOptions,
} from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { crudFactory } from './curd-factory'
import { HttpClient } from './http-client'

export const modelClient = {
  ...crudFactory<Model, QueryOptions, ModelCreateInput>(API_ENDPOINTS.MODELS),
  paginated: ({ name, ...params }: Partial<ModelQueryOptions>) => {
    return HttpClient.get<ModelPaginator>(API_ENDPOINTS.MODELS, {
      searchJoin: 'and',
      self,
      ...params,
      // search: HttpClient.formatSearchParams({ name }),
    })
  },
}
