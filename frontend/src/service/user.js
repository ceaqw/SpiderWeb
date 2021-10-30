// 用户相关服务模块
import { loginApi } from '@/api/login'
import { setToken } from '@/utils/auth'
import router from '@/route'

export function login(data) {
    loginApi(data).then((result) => {
        setToken(result['token'])
        router.push('/')
    }).catch((err) => {
        console.log(err)  
    })
}