import { getToken } from '@/utils/auth/token';
import axios, { type AxiosError } from 'axios'

// create an Axios instance
const baseURL = import.meta.env.VITE_APP_BASE_API?.trim() || '/api'

const Axios = axios.create({
  baseURL,
  timeout: 50000,
  headers: { 'Content-Type': 'application/json;charset=utf-8' },
})

Axios.interceptors.request.use(
  (config) => {
    // 1️⃣ Dynamic base API (Electron setup)
    // const baseApi = baseURL();

    // 2️⃣ Inject auth token (Zustand-safe)
    // const { userToken } = useUserStore.getState();
    // const accessToken = userToken?.accessToken;

    const accessToken = getToken();

    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }

    return config;
  },
  (error) => Promise.reject(error)
);

// Change response data/error here
// Axios.interceptors.response.use(
//   (response) => response,
//   (error: AxiosError<Result>) => {
//     const { response, message } = error || {};

//     const errMsg =
//       response?.data?.message || message || t("sys.api.errorMessage");
//     toast.error(errMsg, {
//       position: "top-center",
//     });

//     const status = response?.status;
//     if (status === 401) {
//       useUserStore.getState().actions.clearUserInfoAndToken(); // ✅ clear on 401
//       useAppStore.getState().openSetup();
//     }

//     return Promise.reject(error);
//   }
// );

export class HttpClient {
  static async get<T>(url: string, params?: unknown) {
    const response = await Axios.get<T>(url, { params })
    return response.data
  }

  static async post<T>(url: string, data: unknown, options?: any) {
    const response = await Axios.post<T>(url, data, options)
    return response.data
  }

  static async put<T>(url: string, data: unknown) {
    const response = await Axios.put<T>(url, data)
    return response.data
  }

  static async patch<T>(url: string, data: unknown) {
    const response = await Axios.patch<T>(url, data)
    return response.data
  }

  static async delete<T>(url: string) {
    const response = await Axios.delete<T>(url)
    return response.data
  }
}
