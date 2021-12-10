/*
 * @Date: 2021-12-10 16:10:16
 * @LastEditTime: 2021-12-10 17:12:00
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

const projectApi = {
    projectList: (data) => {
        return baseGet('/project/projectList', data)
    },
    getProjectInfos: (data) => {
        return baseGet('/project/projectInfos', data)
    },
    getAllPlatformAndProjectMap: (data) => {
        return baseGet('/project/allPlatformAndProjectMap', data)
    }
}

export default projectApi