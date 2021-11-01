import Cookies from 'js-cookie'
import router from '@/route/index'
import {sideItem} from '@/store/routerAuth'
import store from '../store'

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
    // const store = useStore()
    console.log(1111)
    console.log(store.state.userAuth)
    const pathList = to.path.split('/')
    let routerMap = sideItem.routers
    for (let index = 0; index < pathList.length; index++) {
        const path = pathList[index]
        if (path != '') {
            if (sideItem.allowRoutes.indexOf(path) >= 0) continue
            routerMap = routerMap[path]
            if (!routerMap) {
                router.push('/404')
                break
            } else if (routerMap.auth < store.state.userAuth) {
                router.push('/401')
                break
            }
            routerMap = routerMap['routes'] ? routerMap['routes'] : routerMap
        }
    }
    next()
}