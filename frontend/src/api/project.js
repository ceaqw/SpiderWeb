/*
 * @Date: 2021-12-10 16:10:16
 * @LastEditTime: 2021-12-14 11:20:40
 * @Author: ceaqw
 */
import { get, post } from '../utils/request'
import baseConf from '@/conf/baseConf'

const preUrl = baseConf.apiVersion

const basePost = (router, data) => {
    return post(preUrl + router, data)
}

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
    },
    createProject: (data) => {
        return basePost('/project/createProject', data)
    }
}

export default projectApi