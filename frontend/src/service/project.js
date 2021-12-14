/*
 * @Date: 2021-12-10 16:06:29
 * @LastEditTime: 2021-12-14 13:52:04
 * @Author: ceaqw
 */
import projectApi from '../api/project'
import store from '@/store'
import { ElMessage } from 'element-plus'

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

const createProject = (formData) => {
    projectApi.createProject(formData).then((result) => {
        if (result.msg == 'ok') ElMessage({message: formData.id == 0 ? '创建成功' : '更新成功',type: 'success', customClass:'zZindex'})
        else ElMessage({message: formData.id == 0 ? '创建失败' : '更新失败' , type: 'error', customClass:'zZindex'})
    }).catch((err) => {
        console.log(err)
    })
}

export {
    projectList,
    getAllPlatformAndProjectMap,
    getProjectInfos,
    createProject
}