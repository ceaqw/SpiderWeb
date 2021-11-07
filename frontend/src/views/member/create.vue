<template>
    <div class="member-create">
        <el-row>
            <el-col :span="12">
                <span class="title">创建用户</span>
                <el-divider></el-divider>
                <el-form
                    :model="userInfo"
                    status-icon
                    :rules="rules"
                    ref="userInfo"
                    label-width="100px"
                >
                    <el-form-item label="用户名" prop="name">
                        <el-input
                            type="text"
                            v-model="userInfo.name"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="password">
                        <el-input
                            type="password"
                            v-model="userInfo.password"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="确认密码" prop="checkPass">
                        <el-input
                            type="password"
                            v-model="userInfo.checkPass"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="Email" prop="email">
                        <el-input
                            type="text"
                            v-model="userInfo.email"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="手机" prop="mobile">
                        <el-input
                            type="text"
                            v-model="userInfo.mobile"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="微信" prop="wechat">
                        <el-input
                            type="text"
                            v-model="userInfo.wechat"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="权限" prop="role">
                        <el-select placeholder="选择权限" v-model="userInfo.role">
                            <el-option
                                v-for="item in roles"
                                :label="item.name"
                                :key="item.id"
                                :value="item.id"
                            >
                        </el-option>
                    </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
                        <el-button @click="resetForm('ruleForm')">重置</el-button>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>
    </div>
</template>

<script>
import baseConf from '@/conf/baseConf'
import utils from '@/utils/utils'
import check from '@/utils/check'
import userService from '@/service/user'

export default {
    data() {
        let passReg = check.passwordReg
        let validateRePwd = (rule, value, callback) => {
            if (!passReg.test(value)) {
                callback(new Error('密码应是6-12位数字、字母或字符！'))
            } else if (this.userInfo.password != value) {
                callback(new Error('两次输入密码不一致！'))
            } else {
                callback()
            }
        }
        return {
            userInfo: {
                name: '',
                password: '',
                checkPass: '',
                role: baseConf.role,
                email: '',
                mobile: '',
                wechat: ''
            },
            rules: {
                name: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' },
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' },
                    {
                        pattern: passReg,
                        message: '密码应是6-12位数字、字母或字符！',
                        trigger: 'blur'
                    }
                ],
                checkPass: [
                    { required: true, message: '请输入确认密码', trigger: 'blur' },
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
        })
    },
    methods: {
        submitForm() {
            this.$refs.userInfo.validate((valid) => {
                if (valid) {
                    this.loading = true
                    userService.register(this.userInfo)
                } else {
                    this.$message.error('请检查输入')
                    return false
                }
            })
        },
        resetForm() {
            this.$refs.userInfo.resetFields()
        }
    }
}
</script>

<style>
.member-create .el-row{
    text-align: left;
    
}

.member-create .title {
    font-size: 140%;
}

.member-create .el-form-item {
    width: 60%;
    text-align: left;
}

</style>