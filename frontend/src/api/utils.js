import { post } from '@/utils/request'

const preUrl = '/api/v1/user'

const basePost = (router, data) => {
    return post(preUrl + router, data)
}

const utilsApi = {
    login: (data) => {
        return basePost('/login', data)
    },
}

export default utilsApi