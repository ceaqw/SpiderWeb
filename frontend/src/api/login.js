import { post } from '@/utils/request'

// 登录方法
export function loginApi(param) {
    return post('/api/v1/user/login', param)
}