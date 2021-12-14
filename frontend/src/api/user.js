/*
 * @Date: 2021-11-16 09:08:27
 * @LastEditTime: 2021-12-14 14:03:27
 * @Author: ceaqw
 */
import { post, get } from '@/utils/request'
import baseConf from '@/conf/baseConf'

const preUrl = baseConf.apiVersion + '/user'

const basePost = (router, data) => {
    return post(preUrl + router, data)
}

const baseGet = (router, data) => {
    return get(preUrl + router, data)
}

const userApi = {
    login: (data) => {
        return basePost('/login', data)
    },
    logout: (data) => {
        return basePost('/logout', data)
    },
    userList: (data) => {
        return basePost('/userList', data)
    },
    getUserInfo: (data) => {
        return baseGet('/user', data)
    },
    option: (data) => {
        return basePost('/option', data)
    },
    getRoles: (data) => {
        return basePost('/getRoles', data)
    },
    register: (data) => {
        return basePost('/register', data)
    },
    update: (data) => {
        return basePost('/update', data)
    },
    validator: (data) => {
        return baseGet('/validator/' + data.type, data)
    }
}

export default userApi