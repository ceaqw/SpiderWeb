/*
 * @Date: 2021-11-16 09:08:27
 * @LastEditTime: 2021-12-10 16:22:44
 * @Author: ceaqw
 */
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
        return baseGet('/project/projectList', data)
    },
}

export default utilsApi