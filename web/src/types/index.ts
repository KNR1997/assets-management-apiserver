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
}

export interface GetParams {
  id: string
}

export interface QueryOptions {
  language: string
  limit?: number
  page?: number
  orderBy?: string
  sortedBy?: SortOrder
  per_page?: number
}

export interface PaginatorInfo<T> {
  list: T[]
  total: number
  page: number
  pageSize: number
}

export interface LoginInput {
  email: string
  password: string
}

export interface AuthResponse {
  token: string
  tokens: { access: string; refresh: string }
  permissions: string[]
  role: string
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

export interface QuizQueryOptions extends QueryOptions {
  name: string
}

export interface CoursePaginator extends PaginatorInfo<Course> {}

export interface QuizPaginator extends PaginatorInfo<Quiz> {}
