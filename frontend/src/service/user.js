// 用户相关服务模块
import router from '@/route'
import userApi from '../api/user'
import { setToken } from '@/utils/auth'
import { setByKey, getByKey } from '@/utils/cookie'
import { removeByKey } from '@/utils/cookie'
import { ElMessage } from 'element-plus'
// import store from '@/store'

const userService = {
    login: (data) => {
        userApi.login(data).then((result) => {
            setToken(result.data.token)
            setByKey('userAuth', result.data.userAuth)
            setByKey('userId', result.data.userId)
            setByKey('user', data.username)
            router.push('/')
        }).catch((err) => {
            data.loading = false
            console.log(err)
        })
    },
    logout: (data) => {
        userApi.logout(data).then(() => {
            removeByKey('user')
            removeByKey('userAuth')
            removeByKey('userId')
            router.replace('/login')
        }).catch((err) => {
            ElMessage({
                message: '退出失败，请稍后重试',
                type: 'error'
            })
            console.log(err)
        })
    },
    userList: (data, resp) => {
        userApi.userList(data).then((result) => {
            resp.datas = []
            const userList = result.data.userList
            for (const index in userList) {
                resp.datas.push(userList[index])
            }
            resp.totalNumber = result.data.total
        }).catch((err) => {
            console.log(err)
        })
    },
    getUserInfo: (data, resp) => {
        userApi.getUserInfo(data).then((result) => {
            if ('mid' in result.data) {
                resp.mid = result.data.mid
                resp.email = result.data.email
                resp.mobile = result.data.mobile
                resp.wechat = result.data.wechat
            } else {
                ElMessage({
                    message: '获取用户数据无效',
                    type: 'warning'
                })
            }
        }).catch((err) => {
            console.log(err)
        })
    },
    option: (data, resp) => {
        userApi.option(data).then((result) => {
            if (data.option != 'role') {
                for (const index in resp.datas) {
                    if (resp.datas[index].mid == data.mid) {
                        resp.datas[index].status ^= 1
                        break
                    }
                }   
            } else {
                if (getByKey('user') == data.name) {
                    setByKey('userAuth', data.role)
                }
            }
            if (result.status == 200) {
                ElMessage({
                    message: '操作成功',
                    type: 'success'
                })
            }
        }).catch((err) => {
            console.log(err)
        })
    },
    register: (data) => {
        userApi.register(data).then((result) => {
            if (result.status == 200) {
                ElMessage({
                    message: '注册成功',
                    type: 'success'
                })
                if (data.login) {
                    router.replace('/login')
                } 
            }
        }).catch((err) => {
            data.loading = false
            console.log(err)
        })
    },
    update: (data) => {
        userApi.update(data).then((result) => {
            if (result.msg == "ok") ElMessage({message: '更新成功', type: 'success'})
            else ElMessage({message: '更新失败', type: 'error'})
        }).catch((err) => {
            console.log(err)
        })
    },
    validateName: (data) => {
        userApi.validator({name: data.name, type: 'name'}).then((result) => {
            if (result.status == 200 && result.msg == 'ok') {
                data.allow = true
            } else if (result.msg == 'exits') {
                ElMessage({
                    message: '用户名已存在',
                    type: 'error'
                })
            }
        }).catch((err) => {
            console.log(err)
        })
        console.log(data)
    }
}

export default userService