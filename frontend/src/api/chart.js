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
    all: (data, type_) => {
        return basePost('/all/' + type_, data)
    },
}

export default chartApi