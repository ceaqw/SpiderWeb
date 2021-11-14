import { get } from '@/utils/request'
import baseConf from '@/conf/baseConf'

const preUrl = baseConf.apiVersion

// const basePost = (router, data) => {
//     return post(preUrl + router, data)
// }

const baseGet = (router, data) => {
    return get(preUrl + router, data)
}

const utilsApi = {
    projectList: (data) => {
        return baseGet('/base/projectList', data)
    },
}

export default utilsApi