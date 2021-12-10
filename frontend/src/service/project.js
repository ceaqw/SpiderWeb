/*
 * @Date: 2021-12-10 16:06:29
 * @LastEditTime: 2021-12-10 17:08:07
 * @Author: ceaqw
 */
import projectApi from '../api/project'
import store from '@/store'

const projectList = () => {
    projectApi.projectList().then((result) => {
        store.state.projectList = result.data
        for (const key in result.data) {
            store.state.projectList[key].unshift('all')
            // store.state.projectList.all = store.state.projectList.all.concat(result.data[key])
        }
        store.state.projectList['all'] = ['all']
    }).catch((err) => {
        console.log(err)
    })
}

const getAllPlatformAndProjectMap = () => {
    projectApi.getAllPlatformAndProjectMap().then((result) => {
        store.state.allPlatformAndProjectMap = result.data
    }).catch((err) => {
        console.log(err)
    })
}

const getProjectInfos = (filter, tableData) => {
    projectApi.getProjectInfos(filter).then((result) => {
        tableData.datas = result.data.datas
        tableData.totalNumber = result.data.totalNumber
    }).catch((err) => {
        console.log(err)
    })
}

export {
    projectList,
    getAllPlatformAndProjectMap,
    getProjectInfos
}