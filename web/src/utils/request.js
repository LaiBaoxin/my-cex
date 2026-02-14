import axios from 'axios'
import { ElMessage } from 'element-plus'

// 创建 axios 实例
const service = axios.create({
    baseURL: '/',
    timeout: 8000
})

// 请求拦截器
service.interceptors.request.use(
    config => {
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

// 响应拦截器
service.interceptors.response.use(
    response => {
        return response.data
    },
    error => {
        const msg = error.response?.data?.message || '网络异常，请检查后端服务'
        ElMessage.error(msg)
        return Promise.reject(error)
    }
)

export default service