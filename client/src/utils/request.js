/*
 * @Description: 网络请求工具
 * @Author: pixel-revolve
 * @Date: 2022-04-30 14:39:53
 */
import axios from 'axios'
import {
    MessageBox,
    Message
} from 'element-ui'

// 创建Axios全局实例
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API + '/api',
    withCredentials: true, // 设置请求时带Cookies
    timeout: 25000 // 请求延迟
})

// response interceptor
service.interceptors.response.use(
    response => {
        const res = response.data;
        if (res.code !== 200) {
            Message({
                message: res.message,
                type: 'error'
            })
            return Promise.reject(res)
        } else {
            return response;
        }
    },
    error => {
        return Promise.reject(error)
    }
)

export default service