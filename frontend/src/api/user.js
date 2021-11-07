import { post } from '@/utils/request'

const preUrl = '/api/v1/user'

const basePost = (router, data) => {
    return post(preUrl + router, data)
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
    option: (data) => {
        return basePost('/option', data)
    },
    getRoles: (data) => {
        return basePost('/getRoles', data)
    },
}

export default userApi