// 图表数据请求api模块
import { post } from '@/utils/request'
import baseConf from '@/conf/baseConf'

const preUrl = baseConf.apiVersion + '/chart'

const basePost = (router, data) => {
    return post(preUrl + router, data)
}

// const baseGet = (router, data) => {
//     return get(preUrl + router, data)
// }


const chartApi = {
    all: (data) => {
        return basePost('/all', data)
    },
    analyse: (data) => {
        return basePost('/analyse', data)
    },
    rakuten: (data) => {
        return basePost('/all/rakuten', data)
    },
    yahoo: (data) => {
        return basePost('/all/yahoo', data)
    },
    amazon: (data) => {
        return basePost('/all/amazon', data)
    },
    topError: (data) => {
        return basePost('/topError', data)
    },
}

export default chartApi