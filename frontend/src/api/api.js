
import axios from "axios"

const axiosApi = axios.create({
    baseURL: "http://localhost:8080/query",
})
axiosApi.interceptors.response.use(
    response => response,
    error => Promise.reject(error)
)

export async function post(url, data, config = {}) {
    return axiosApi
        .post(url, { ...data }, { ...config })
        .then(response => response)
}