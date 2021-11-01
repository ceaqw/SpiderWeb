<template>
    <div class="login-wrap">
        <div class="ms-login">
            <div class="ms-title">{{ title }}</div>
            <el-form :model="form"
                :rules="rules"
                ref="login"
                label-width="0px"
                class="ms-content">
                <el-form-item prop="username">
                    <el-input v-model="form.username" placeholder="username">
                        <el-button icon="el-icon-lx-people"></el-button>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input type="password"
                            placeholder="password"
                            v-model="form.password"
                            @keyup.enter="submitForm">
                        <el-button icon="el-icon-lx-lock"></el-button>
                    </el-input>
                </el-form-item>
                <div class="login-btn">
                    <el-button type="primary" @click="submitForm">登录</el-button>
                </div>
                <p class="login-tips">没有账号？请先注册</p>
                <div class="login-reg">
                    <el-button type="primary" @click="registerMethod">注册</el-button>
                </div>
                <router-link to="/">暂不登录，游客预览</router-link>
            </el-form>
        </div>
    </div>
</template>

<script>
import userService from '@/service/user'
import { passwordReg } from '@/utils/check'
import { setByKey } from '@/utils/cookie'

export default {
    props: ['title'],
    data() {
        let reg = passwordReg
        return {
            form: {
                username: '',
                password: ''
            },
            rules: {
                username: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' },
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' },
                    {
                        required: true,
                        pattern: reg,
                        message: '密码应是6-12位数字、字母或字符！',
                        trigger: 'blur'
                    }
                ]
            },
            loading: false
        }
    },
    methods: {
        submitForm() {
            this.$refs.login.validate((valid) => {
                if (valid) {
                    this.loading = true
                    userService.login(this.form)
                    setByKey('user', this.form.username)
                } else {
                    this.$message.error('请检查输入')
                    return false
                }
            })
        },
        // 注册
        registerMethod() {
            this.$router.push('/register')
        }
    }
}
</script>

<style scoped>
.login-wrap {
  position: relative;
  width: 100%;
  height: 100%;
  background: url('../assets/loginBg.gif') no-repeat center center;
  background-size: 100%;
}

.ms-login {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 350px;
  margin: -190px 0 0 -175px;
  border-radius: 5px;
  background: rgba(255, 255, 255, 0.3);
  overflow: hidden;
}

.ms-title {
  width: 100%;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
  color: #fff;
  border-bottom: 1px solid #ddd;
}

.ms-content {
  padding: 30px 30px;
}
.login-btn,
.login-reg {
  text-align: center;
}
.login-btn button,
.login-reg button {
  width: 100%;
  height: 36px;
  color: black;
  margin-bottom: 10px;
}
.login-tips {
  text-align: center;
  font-size: 12px;
  line-height: 30px;
  color: rgb(3, 70, 31);
}

.login-reg button {
  background-color: deepskyblue;
}
</style>
