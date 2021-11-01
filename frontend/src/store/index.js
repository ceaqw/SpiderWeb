// 变量储存共享模块

import { createStore } from 'vuex'

export default createStore({
    state () {
        return {
            userAuth: 1
        }
    },
    mutations: {
        set(state, arg) {
            state[arg[0]] = arg[1]
        }
    },
})
