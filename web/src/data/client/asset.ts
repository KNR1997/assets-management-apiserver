import type {
  Asset,
  AssetCheckinInput,
  AssetCheckoutInput,
  AssetCreateInput,
  AssetPaginator,
  AssetQueryOptions,
  QueryOptions,
} from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { crudFactory } from './curd-factory'
import { HttpClient } from './http-client'

export const assetClient = {
  ...crudFactory<Asset, QueryOptions, AssetCreateInput>(API_ENDPOINTS.ASSETS),
  paginated: ({ name, ...params }: Partial<AssetQueryOptions>) => {
    return HttpClient.get<AssetPaginator>(API_ENDPOINTS.ASSETS, {
      searchJoin: 'and',
      self,
      ...params,
      // search: HttpClient.formatSearchParams({ name }),
    })
  },
  checkout: (variables: AssetCheckoutInput) => {
    return HttpClient.post(`${API_ENDPOINTS.ASSETS}/${variables.assetId}/checkout`, variables)
  },
  checkin: (variables: AssetCheckinInput) => {
    return HttpClient.post(`${API_ENDPOINTS.ASSETS}/${variables.assetId}/checkin`, variables)
  },
}
