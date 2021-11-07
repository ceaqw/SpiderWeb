import axios from 'axios'
import { ElNotification, ElMessageBox, ElMessage  } from 'element-plus'
import { getToken } from '@/utils/auth'
import errorCode from '@/utils/errorCode'
// import qs from 'qs'

axios.defaults.headers['Content-Type'] = 'application/jsoncharset=utf-8'
axios.defaults.timeout = 10000

// request拦截器
axios.interceptors.request.use(
    config => {
        // 是否需要设置 token
        const isToken = (config.headers || {}).isToken === false
        if (getToken() && !isToken) {
            config.headers['Authorization'] = 'Basic ' + getToken() // 让每个请求携带自定义token 请根据实际情况自行修改
        }
        // get请求映射params参数
        if (config.method === 'get' && config.params) {
            let url = config.url + '?'
            for (const propName of Object.keys(config.params)) {
                const value = config.params[propName]
                var part = encodeURIComponent(propName) + '='
                if (value !== null && typeof(value) !== 'undefined') {
                    if (typeof value === 'object') {
                        for (const key of Object.keys(value)) {
                            let params = propName + '[' + key + ']'
                            var subPart = encodeURIComponent(params) + '='
                            url += subPart + encodeURIComponent(value[key]) + '&'
                        }
                    } else {
                        url += part + encodeURIComponent(value) + '&'
                    }
                }
            }
            url = url.slice(0, -1)
            config.params = {}
            config.url = url
        }
        return config    
    },
    error => {
        console.log(error)
        return error
    }
)

// 响应拦截器
axios.interceptors.response.use(
    res => {
        // 未设置状态码则默认成功状态
        const code = res.data.status || 200
        // 获取错误信息
        const msg = errorCode[code] || res.data.msg || errorCode['default']
        if (code === 401) {
            ElMessageBox.confirm('登录状态已过期，您可以继续留在该页面，或者重新登录', '系统提示', {
                confirmButtonText: '重新登录',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                // TODO: 退出
            }).catch(() => {})
        } else if (code === 500) {
            ElMessage({
                message: msg,
                type: 'error'
            })
            return msg
        } else if (code !== 200) {
            ElNotification.error({
                title: msg
            })
            return 'error'
        } else {
            return res
        }  
    }, 
    error => {
        console.log('err' + error)
        let message = error
        if (message == 'Network Error') {
            message = '后端接口连接异常'
        } else if (message.includes('timeout')) {
            message = '系统接口请求超时'
        } else if (message.includes('Request failed with status code')) {
            message = '系统接口' + message.substr(message.length - 3) + '异常'
        }
        ElMessage({
            message: message,
            type: 'error',
            duration: 5 * 1000
        })
        return error 
    }
)

export function post(url, data) {
    return new Promise((resolve, reject) => {    
        axios({
                method: 'post',
                url,
                data: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            })
        }
    )
}
export function get(url, data) {
    return new Promise((resolve, reject) => {
        axios({
                method: 'get',
                url,
                params: data,
            }).then(res => {
                resolve(res.data)
            }).catch(err => {
                reject(err)
            })
        }
    )
}
