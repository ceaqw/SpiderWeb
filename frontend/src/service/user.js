// 用户相关服务模块
import { loginApi } from '@/api/login'
import { setToken } from '@/utils/auth'
import router from '@/route'
import store from '@/store'

export function login(data) {
    loginApi(data).then((result) => {
        setToken(result['token'])
        // TODO: 更新用户权限值
        store.commit('set', ['userAuth', 1])
        router.push('/')
    }).catch((err) => {
        console.log(err)  
    })
}