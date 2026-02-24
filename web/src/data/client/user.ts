import type {
  ProfileUpdateInput,
  QueryOptions,
  User,
  UserCreateInput,
  UserMeResponse,
} from '@/types'
import { API_ENDPOINTS } from './api-endpoints'
import { HttpClient } from './http-client'
import { crudFactory } from './curd-factory'

export const userClient = {
  ...crudFactory<User, QueryOptions, UserCreateInput>(API_ENDPOINTS.USERS),
  update: (variables: ProfileUpdateInput) => {
    return HttpClient.patch(API_ENDPOINTS.PROFILE, variables)
  },
  me: () => {
    return HttpClient.get<UserMeResponse>(API_ENDPOINTS.ME)
  },
}
