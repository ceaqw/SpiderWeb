// 基本工具包
import store from '@/store'
import userApi from '@/api/user'

const utils = {
    refreshRoles: () => {
        // 监测是否需要刷新
        if (store.state.roles.length == 0) {
            userApi.getRoles().then((result) => {
                const roles = result.data.roles
                for (const index in roles) {
                    store.state.roles.push(roles[index])
                }
            }).catch((err) => {
                console.log(err)
            })
        }
    },
}

export default utils