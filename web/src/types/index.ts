import type { Ref } from "vue"
import type { AssetStatus } from "./enum"

export enum SortOrder {
  Asc = 'asc',
  Desc = 'desc',
}

// -> TODO: Simplify this
export interface MappedPaginatorInfo {
  currentPage: number
  firstPageUrl: string
  from: number
  lastPage: number
  lastPageUrl: string
  links: any[]
  nextPageUrl: string | null
  path: string
  perPage: number
  prevPageUrl: string | null
  to: number
  total: number
  hasMorePages: boolean
  page: number
  limit: number
  totalRows: number
}

export interface GetParams {
  id: string
}

export interface QueryOptions {
  language: string
  limit?: Ref<number, number>
  page?: Ref<number, number>
  orderBy?: string
  sortedBy?: SortOrder
  per_page?: number
}

export interface PaginatorInfo<T> {
  limit: number
  page: number
  sort: string
  total_rows: number
  total_pages: number
  rows: T[]
}

export interface LoginInput {
  email: string
  password: string
}

export interface AuthResponse {
  data: string
  // token: string
  // tokens: { access: string; refresh: string }
  // permissions: string[]
  // role: string
}

export interface Course {
  id: string
  ID: string // todo -> check this
  name: string
  slug: string
}

export interface CourseCreateInput {
  name: string
  slug: string
}

export interface User {
  id: string
  username: string
  email: string
}

export interface UserCreateInput {
  username: string
  email: string | null
}

export interface Manufacturer {
  id: string
  name: string
  email: string
}

export interface ManufacturerCreateInput {
  name: string
  email: string
}

export interface Supplier {
  id: string
  name: string
  email: string
}

export interface SupplierCreateInput {
  name: string
}

export interface Model {
  id: string
  name: string
  categoryId: string
  manufacturerId: string
  modelNumber: string
}

export interface ModelCreateInput {
  name: string
  categoryId: string
  manufacturerId: string
  modelNumber: string
}

export interface Asset {
  id: string
  name: string
  tag: string
  serialNumber: string
  categoryId: string
  status: AssetStatus
  description: string
  model: Model
}

export interface AssetCreateInput {
  name: string
  tag: string
  serialNumber: string
  description: string
  status: string
  modelId: string
}

export interface AssetCheckoutInput {
  assetName: string
  assetId: number
  userId: number
  checkoutDate: string
  expectedCheckinDate: string | null
  notes: string
}

export interface AssetCheckinInput {
  assetName: string
  assetId: number
  checkinDate: string
  status: AssetStatus
  notes: string
}

export interface Category {
  id: string
  name: string
  serialNumber: string
}

export interface CategoryCreateInput {
  name: string
  description: string | null
}

export interface Department {
  id: string
  name: string
  notes: string
}

export interface DepartmentCreateInput {
  name: string
  notes: string
}

export interface ProfileUpdateInput {
  username: string
  email: string
}

export interface UserMeResponse {
  data: {
    id: number
    username: string
    email: string
    role: string
  }
}

export interface Quiz {
  id: number
  ID: number
  courseId: number
  name: string
  code: string
  date: number
}

export interface QuizCreateInput {
  courseId: number
  name: string
  dateTime: string
}

export interface CourseListResponse {
  list: Course[]
  total: number
  page: number
  pageSize: number
}

export interface QuizListResponse {
  list: Quiz[]
  total: number
  page: number
  pageSize: number
}

export interface CourseQueryOptions extends QueryOptions {
  name: string
}

export interface AssetQueryOptions extends QueryOptions {
  name: string
}

export interface SupplierQueryOptions extends QueryOptions {
  name: string
}

export interface ModelQueryOptions extends QueryOptions {
  name: string
}

export interface ManufacturerQueryOptions extends QueryOptions {
  name: string
}
export interface CategoryQueryOptions extends QueryOptions {
  name: string
}

export interface DepartmentQueryOptions extends QueryOptions {
  name: string
}

export interface UserQueryOptions extends QueryOptions {
  name: string
}

export interface QuizQueryOptions extends QueryOptions {
  name: string
}

export interface CoursePaginator extends PaginatorInfo<Course> {}

export interface AssetPaginator extends PaginatorInfo<Asset> {}

export interface ModelPaginator extends PaginatorInfo<Model> {}

export interface ManufacturerPaginator extends PaginatorInfo<Manufacturer> {}

export interface SupplierPaginator extends PaginatorInfo<Supplier> {}

export interface CategoryPaginator extends PaginatorInfo<Category> {}

export interface DepartmentPaginator extends PaginatorInfo<Department> {}

export interface QuizPaginator extends PaginatorInfo<Quiz> {}
