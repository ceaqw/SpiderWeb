<template>
    <div>
        <el-menu default-active="1-4-1" class="el-menu-vertical-demo" background-color="#D3DCE6">
            <span v-for="(item, index) in sideRoutes" :key="index">
                <div v-if="item.auth == userAuth">
                    <el-menu-item v-if="item.path != 'group'" :index="index" @click="to('/home/'+item.path)">
                        <i :class="item.icon"></i>
                        <template #title>{{ item.path }}</template>   
                    </el-menu-item>
                    <el-sub-menu v-else-if="item.path == 'group'" :index="index">
                        <template #title>
                            <i :class="item.icon"></i>
                            <span>{{ item.name }}</span>
                        </template>
                        <el-menu-item 
                            v-for="(subItem, subIndex) in item.routes" 
                            :key="subIndex" 
                            :index="index+'-'+subIndex"
                            @click="to('/' + item.name + '/' + subItem.path)">
                            <i :class="subItem.icon"></i>
                            <template #title>{{ subItem.path }}</template>
                        </el-menu-item>
                    </el-sub-menu> 
                </div>
            </span>
        </el-menu>
    </div>
</template>

<script>
import { sideItem } from '@/store/routerAuth'
export default {
    data () {
        return {
            sideRoutes: sideItem.routers,
            userAuth: 3
        }
    },
    methods: {
        to(url) {
            this.$router.push(url)
        }
    }
}
</script>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    min-height: 400px;
}
.el-menu {
    text-align: left;
}
</style>
