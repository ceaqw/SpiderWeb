/*
 * @Date: 2021-11-29 09:13:10
 * @LastEditTime: 2021-12-10 17:07:49
 * @Author: ceaqw
 */
// 全局共享变量储存模块

import baseConf from '@/conf/baseConf'

import { createStore } from 'vuex'

export default createStore({
    state () {
        return {
            pollingInterval: 2,
            shareFilter: Object.assign({}, baseConf.filterForm),
            dashboardFilter: Object.assign({}, baseConf.filterForm),
            projectFilter: Object.assign({}, baseConf.filterForm),
            kpiFilter: Object.assign({}, baseConf.filterForm),
            failFilter: Object.assign({}, baseConf.filterForm),
            roles: [],
            FilterSharing: true,
            projectList: {},
            allPlatformAndProjectMap: {},
            flushQueue: {dashboard: {}, project: {}, kpi: {}, fail: {}},
            count: 1,
            pageTaskQueue: {},
            cacheChart: {},
        }
    }
})
