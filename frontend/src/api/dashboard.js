import { post, get } from '@/utils/request'
import baseConf from '@/conf/baseConf'

const preUrl = baseConf.apiVersion + '/dashboard'

const basePost = (router, data) => {
    return post(preUrl + router, data)
}

const baseGet = (router, data) => {
    return get(preUrl + router, data)
}

const dashboardApi = {
    login: (data) => {
        return basePost('/login', data)
    },
}

export default dashboardApi