<template>
  <div class="login-wrap">
    <div class="ms-login">
      <div class="ms-title">{{ title }}</div>
      <el-form :model="param"
               :rules="rules"
               ref="login"
               label-width="0px"
               class="ms-content">
        <el-form-item prop="name">
          <el-input v-model="param.name"
                    placeholder="用户名"
                    @keyup.enter="submitForm">
            <el-button icon="el-icon-lx-people"></el-button>
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input type="password"
                    placeholder="输入密码"
                    v-model="param.password"
                    @keyup.enter="submitForm">
            <el-button icon="el-icon-lx-lock"></el-button>
          </el-input>
        </el-form-item>
        <el-form-item prop="re_password">
          <el-input type="password"
                    placeholder="确认密码"
                    v-model="param.re_password"
                    @keyup.enter="submitForm">
            <el-button icon="el-icon-lx-lock"></el-button>
          </el-input>
        </el-form-item>
        <div class="login-reg">
          <el-button type="primary" @click="submitForm">
                <span>注 册 <i class="el-icon-loading" v-if="param.loading"></i></span>
          </el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
// import Cookie from 'js-cookie'
import userService from '@/service/user'
import check from '@/utils/check'

export default {
    props: ['title'],
    data() {
        let regPass = check.passwordReg
        let validateRePwd = (rule, value, callback) => {
            if (!regPass.test(value)) {
                callback(new Error('密码应是6-12位数字、字母或字符！'))
            } else if (this.param.password != value) {
                callback(new Error('两次输入密码不一致！'))
            } else {
                callback()
            }
        }
        return {
            param: {
                name: '',
                password: '',
                re_password: '',
                login: true,
                loading: false
            },
            rules: {
                name: [
                    { required: true, message: '请输入用户名', trigger: 'blur' },
                    { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' },
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' },
                    { min: 6, max: 12, message: '长度在 6 到 12 个字符', trigger: 'blur' },
                    {
                        required: true,
                        pattern: regPass,
                        message: '密码应是6-12位数字、字母或字符！',
                        trigger: 'blur'
                    }
                ],
                re_password: [
                    { required: true, message: '请输入确认密码', trigger: 'blur' },
                    { validator: validateRePwd, trigger: 'blur' }
                ]
            },
        }
    },
    methods: {
        submitForm() {
            this.$refs.login.validate((valid) => {
                if (valid) {
                    this.param.loading = true
                    userService.register(this.param)
                } else {
                    this.$message.error('请检查输入')
                    return false
                }
            })
        }
    }
}
</script>

<style scoped>
.login-wrap {
  position:absolute;
  width: 100%;
  height: 100%;
  /* background: url('../assets/loginBg.gif') no-repeat center center; */
  background-color: #E9EEF3;
  background-size: 100%;
}
.ms-title {
  width: 100%;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
  /* color: rgb(55, 66, 59); */
  color: #252526;
  border-bottom: 1px solid #ddd;
}
.ms-login {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 350px;
  margin: -190px 0 0 -175px;
  border-radius: 5px;
  background: rgba(255, 255, 255, 0.7);
  overflow: hidden;
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
.login-reg button {
  background-color: deepskyblue;
}
</style>
