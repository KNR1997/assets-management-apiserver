import type {
  AssetQueryOptions,
  Manufacturer,
  ManufacturerCreateInput,
  ManufacturerPaginator,
  QueryOptions,
} from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { crudFactory } from './curd-factory'
import { HttpClient } from './http-client'

export const manufacturerClient = {
  ...crudFactory<Manufacturer, QueryOptions, ManufacturerCreateInput>(API_ENDPOINTS.MANUFACTURERS),
  paginated: ({ name, ...params }: Partial<AssetQueryOptions>) => {
    return HttpClient.get<ManufacturerPaginator>(API_ENDPOINTS.MANUFACTURERS, {
      searchJoin: 'and',
      self,
      ...params,
      // search: HttpClient.formatSearchParams({ name }),
    })
  },
}
