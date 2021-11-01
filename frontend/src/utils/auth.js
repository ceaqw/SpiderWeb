import Cookies from 'js-cookie'
import router from '@/route/index'
import {sideItem} from '@/store/routerAuth'

const TokenKey = 'Admin-Token'

export function getToken() {
    return Cookies.get(TokenKey)
}

export function setToken(token) {
    return Cookies.set(TokenKey, token)
}

export function removeToken() {
    return Cookies.remove(TokenKey)
}

export function routerAuth(to, from, next) {
    const currentAuth = 1
    const pathList = to.path.split('/')
    let routerMap = sideItem.routers
    for (let index = 0; index < pathList.length; index++) {
        const path = pathList[index]
        if (path != '') {
            routerMap = routerMap[path]
            console.log(routerMap)
            if (routerMap.auth < currentAuth) {
                router.push('/401')
                break
            }
            routerMap = routerMap['routes'] ? routerMap['routes'] : routerMap
        }
    }
    next()
    // router.push('/401')
}