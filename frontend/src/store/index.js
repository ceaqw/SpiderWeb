// 全局共享变量储存模块

import { createStore } from 'vuex'

export default createStore({
    state () {
        return {
            pollingInterval: 2,
            dateRange: 3,
            filterForm: {
                dateRange: ['', ''],
                platForm: 'all',
                project: 'all',
                showType: 1
            }
        }
    },
    mutations: {
        set(state, arg) {
            state[arg[0]] = arg[1]
        }
    },
})
