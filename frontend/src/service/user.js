// 用户相关服务模块
import { loginApi } from '@/api/login'
import { setToken } from '@/utils/auth'
import { setByKey } from '@/utils/cookie'
import router from '@/route'
import { removeByKey } from '@/utils/cookie'
// import store from '@/store'

const login = (data) => {
    loginApi(data).then((result) => {
        setToken(result['token'])
        // TODO: 更新用户权限值
        setByKey('userAuth', 1)
        router.push('/')
    }).catch((err) => {
        console.log(err)  
    })
}

// 登出操作
const logout = () => {
    removeByKey('user')
    removeByKey('userAuth')
    // 更新页面，防止懒加载
    router.replace('/login')
}

const userService = {
    login,
    logout,
}

export default userService