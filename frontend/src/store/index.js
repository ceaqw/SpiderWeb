// 全局共享变量储存模块

import { createStore } from 'vuex'

export default createStore({
    state () {
        return {
            pollingInterval: 2,
            dateRangeType: 3,
            filterForm: {
                startDate: '',
                endDate: '',
                platForm: 'all',
                project: 'all',
                showType: 1
            },
            roles: [],
        }
    },
    mutations: {
        set(state, arg) {
            state[arg[0]] = arg[1]
        }
    },
})
