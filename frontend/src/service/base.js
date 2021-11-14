// 用户相关服务模块
import utilsApi from '@/api/utils'
import store from '@/store'

const baseService = {
    projectList: () => {
        utilsApi.projectList().then((result) => {
            store.state.projectList = result.data
            store.state.projectList['all'] = []
            for (const key in result.data) {
                store.state.projectList.all = store.state.projectList.all.concat(result.data[key])
            }
        }).catch((err) => {
            console.log(err)
        })
    },
}

export default baseService