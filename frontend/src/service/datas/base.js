// 基本数据后端同步
import store from '@/store'
import baseService from '../base'

const getFilterData = () => {
    // 需要刷新
    if (!Object.keys(store.state.projectList).length) {
        baseService.projectList()
    }
}

export {
    getFilterData,
}