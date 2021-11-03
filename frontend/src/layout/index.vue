
<template>
<div class="common-layout">
    <el-container style="height: 100%; border: 0px">
        <el-header>
            <el-row>
                <el-col :span="15" style="text-align: left;">
                    <el-row>
                        <el-col :span="3" style="text-align: right"><img id="logo" :src="url"></el-col>
                        <el-col :span="1"></el-col>
                        <el-col :span="20" id="title">{{ title }}</el-col>
                    </el-row>
                </el-col>
                <el-col :span="6" style="text-align: right;">
                    <el-button-group id="time">
                        <el-button v-for="(rate, index) in rateSelect"
                            :key="index"
                            :id="rateFalg[index]"
                            size="medium"
                            @click="changeRate(index)" plain>
                            {{ rate }}
                        </el-button>
                    </el-button-group>
                </el-col>
                <el-col :span="2" style="text-align: right;">
                    <el-dropdown v-if="user">
                        <span style="font-size: 18px">{{ user }}</span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item>个人中心</el-dropdown-item>
                                <el-dropdown-item @click="logout">登出</el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                    <router-link to="/login" style="font-size: 16px;" v-else>登录</router-link>
                </el-col>
                <el-col :span="1"></el-col>
            </el-row>
        </el-header>
        <el-container>
            <el-aside>
                <side-menu></side-menu>
            </el-aside>
            <el-main>
                <router-view></router-view>
            </el-main>
        </el-container>
  </el-container>
</div>
</template>

<script>
import logo from '@/assets/logo.png'
import userService from '@/service/user'
import SideMenu from './sideMenu.vue'
import { getByKey } from '@/utils/cookie'

export default {
    props: ['title'],
    components: {
        SideMenu,
    },
    data () {
        return {
            url: logo + '?' + +new Date(),
            user: getByKey('user'),
            rateFalg: ['', '', 'primary'],
            rateSelect: ['30S', '60S', 'OFF']
        }
    },
    methods: {
        changeRate(index) {
            this.rateFalg[this.rateFalg.indexOf('primary')] = 'info'
            this.rateFalg[index] = 'primary'
            // TODO: 更新轮询机制
            this.$store.state.pollingInterval = index
            alert(this.$store.state.pollingInterval)
        },
        logout() {
            userService.logout()
        }
    }
}
</script>

<style>
    .el-header, .el-footer {
        width: 100%;
        background-color: #b3c0d1;
        color: var(--el-text-color-primary);
        text-align: center;
        line-height: 50px;
    }
    /* .el-header {
        position: relative;
    } */

    .el-aside {
        background-color: #d3dce6;
        color: var(--el-text-color-primary);
        text-align: center;
        display: block;
        position: absolute;
        left: 0;
        top: 50px;
        bottom: 0;
        width: 200px;
    }

    .el-main {
        padding: 20px 20px 0px !important;
        background-color: #e9eef3;
        color: var(--el-text-color-primary);
        text-align: center;
        position: absolute;
        right: 0;
        left: 200px;
        top: 50px;
        bottom: 0;
        overflow-y: scroll;
    }

    .el-container:nth-child(5) .el-aside,
    .el-container:nth-child(6) .el-aside {
        line-height: 260px;
    }

    .el-container:nth-child(7) .el-aside {
        line-height: 320px;
    }

    #time .el-button {
        background-color: #b3c0d1;
        border: 0px;
        color: black;
    }

    #time #primary {
        background-color: #409EFF;
    }
</style>
