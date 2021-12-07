/*
 * @Date: 2021-11-16 09:08:27
 * @LastEditTime: 2021-12-07 12:45:12
 * @Author: ceaqw
 */
// 用户相关服务模块
import utilsApi from '@/api/utils'
import store from '@/store'

const baseService = {
    projectList: () => {
        utilsApi.projectList().then((result) => {
            store.state.projectList = result.data
            for (const key in result.data) {
                store.state.projectList[key].unshift('all')
                // store.state.projectList.all = store.state.projectList.all.concat(result.data[key])
            }
            store.state.projectList['all'] = ['all']
        }).catch((err) => {
            console.log(err)
        })
    },
}

export default baseService