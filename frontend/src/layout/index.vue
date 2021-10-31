
<template>
<div class="common-layout">
  <el-container style="height: 100%; border: 0px">
        <el-header>
            <el-row>
                <el-col :span="12" style="text-align: left;">
                    <el-row>
                        <el-col :span="3" style="text-align: right"><img id="logo" :src="url"></el-col>
                        <el-col :span="1"></el-col>
                        <el-col :span="20" id="title">{{ title }}</el-col>
                    </el-row>
                </el-col>
                <el-col :span="12" style="text-align: right;">
                    <el-dropdown>
                        <span style="font-size: 18px">{{ user }}</span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item>个人中心</el-dropdown-item>
                                <el-dropdown-item>登出</el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </el-col>
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
import SideMenu from './sideMenu.vue'

export default {
    props: ['title'],
    components: {
        SideMenu,
    },
    data () {
        return {
            url: logo + '?' + +new Date(),
            user: 'chen.chen'
        }
    }
}
</script>

<style>
    .el-header {
        position: relative;
        width: 100%;
        background-color: #b3c0d1;
        color: var(--el-text-color-primary);
        text-align: center;
        line-height: 60px;
    }

    .el-aside {
        background-color: #d3dce6;
        color: var(--el-text-color-primary);
        text-align: center;
        display: block;
        position: absolute;
        left: 0;
        top: 60px;
        bottom: 0;
        width: auto;
    }

    .el-main {
        background-color: #e9eef3;
        color: var(--el-text-color-primary);
        text-align: center;
        position: absolute;
        left: 200px;
        right: 0;
        top: 60px;
        bottom: 0;
        overflow-y: scroll;
        margin: 0px;
        width: auto;
    }

    body > .el-container {
        margin-bottom: 5px;
    }

    .el-container:nth-child(5) .el-aside,
    .el-container:nth-child(6) .el-aside {
        line-height: 260px;
    }

    .el-container:nth-child(7) .el-aside {
        line-height: 320px;
    }
</style>
