<template>
        <div class="member-profile">
        <el-row>
            <el-col :span="12">
                <span class="title">个人信息</span>
                <el-divider></el-divider>
                <el-form
                    :model="userInfo"
                    status-icon
                    ref="userInfo"
                    :rules="rules"
                    label-width="100px"
                >
                    <el-form-item label="用户名" prop="name">
                        <el-input
                            type="text"
                            v-model="userInfo.name"
                        ></el-input>
                    </el-form-item>
                    <el-form-item v-for="item, index in formInfo" :key="index" :label="item.label" :prop="item.prop">
                        <el-input
                            :type="item.type"
                            v-model="userInfo[item.prop]"
                        ></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submit">提交</el-button>
                        <el-button @click="resetForm('userInfo')">重置</el-button>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import check from '@/utils/check'
import utils from '../../utils/utils'
import userService from '../../service/user'
import {getByKey} from '../../utils/cookie'

export default {
    data() {
        let passReg = check.passwordReg
        let validateRePwd = (rule, value, callback) => {
            if (this.userInfo.password != value) {
                callback(new Error('两次输入密码不一致！'))
            } else {
                callback()
            }
        }
        return {
            userInfo: {
                mid: '',
                name: '',
                password: '',
                checkPass: '',
                email: '',
                mobile: '',
                wechat: ''
            },
            formInfo: [
                {label: '新密码', prop: 'password', type: 'password'},
                {label: '确认密码', prop: 'checkPass', type: 'password'},
                {label: 'Email', prop: 'email', type: 'text'},
                {label: '手机', prop: 'mobile', type: 'text'},
                {label: '微信', prop: 'wechat', type: 'text'},
            ],
            rules: {
                password: [
                    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' },
                    {
                        pattern: passReg,
                        message: '密码应是6-12位数字、字母或字符！',
                        trigger: 'blur'
                    }
                ],
                checkPass: [
                    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' },
                    { validator: validateRePwd, trigger: 'blur' }
                ],
                email: [
                    {pattern: check.emailReg, message: '邮箱格式错误', trigger: 'blur'}
                ],
                mobile: [
                    {pattern: check.mobileReg, message: '手机号码格式错误', trigger: 'blur'}
                ]
            },
            roles: this.$store.state.roles,
        }
    },
    mounted() {
        this.$nextTick(() => {
            utils.refreshRoles()
            this.userInfo.name = getByKey('user')
            this.getUserInfo()
        })
    },
    methods: {
        getUserInfo() {
            userService.getUserInfo({user: this.userInfo.name}, this.userInfo)
        },
        resetForm(formName) {
            this.$refs[formName].resetFields()
        },
        submit() {
            this.$refs.userInfo.validate((valid) => {
                if (valid) {
                    userService.update(this.userInfo)
                }
            })
        }
    }
}
</script>

<style>
.member-profile .el-form-item {
    width: 60%;
    text-align: left;
}
</style>