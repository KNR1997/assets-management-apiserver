import { API_ENDPOINTS } from "./api-endpoints"
import { HttpClient } from "./http-client"

export const healthClient = {
    check: () => {
        return HttpClient.get(API_ENDPOINTS.HEALTH)
    }
}