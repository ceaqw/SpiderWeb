<template>
    <div class="member-list">
        <el-row>
            <span class="title">用户列表</span>
            <el-divider></el-divider>
            <el-table :data="tableData.datas" stripe style="width: 100%">
                <el-table-column prop="mid" label="Mid"> </el-table-column>
                <el-table-column prop="name" label="Name"> </el-table-column>
                <el-table-column label="Status">
                    <template #default="scope">
                        <span style="margin-left: 10px">{{ memberStatus[scope.row.status] }}</span>
                    </template>
                </el-table-column>
                <el-table-column label="Options">
                    <template #default="scope">
                        <el-button type="warning" v-if="scope.row.status == 1" icon="el-icon-refresh-left" size="medium" title="恢复" @click.prevent="option(scope.row, 'recover')"></el-button>
                        <el-button type="danger" v-else-if="scope.row.status == 0" icon="el-icon-delete" size="medium" title="删除" @click.prevent="option(scope.row, 'del')"></el-button>
                        <el-select v-if="scope.row.status == 0" placeholder="选择权限" style="width: 40%; marginLeft: 10px;" v-model="scope.row.role" @change="option(scope.row, 'role')">
                            <el-option
                                v-for="item in roles"
                                :label="item.name"
                                :key="item.id"
                                :value="item.id"
                                >
                            </el-option>
                        </el-select>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination background layout="prev, pager, next" :total="tableData.totalNumber" @current-change="changePage" :page-size="tableData.pageSize">
            </el-pagination>
        </el-row>
    </div>
</template>

<script>
import userService from '@/service/user'
import baseConf from '@/conf/baseConf'
import utils from '@/utils/utils'

export default {
    data() {
        return {
            memberStatus: [
                '正常', '已失效', 
            ],
            roles: this.$store.state.roles,
            tableData: {
                datas: [
                    // {
                //     mid: '1000',
                //     name: 'cai.hongliang',
                //     status: 1,
                //     role: 1,
                // },
                ],
                totalNumber: 0,
                pageSize: baseConf.pageSize,
            },
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.changePage(1)
            utils.refreshRoles()
        })
    },
    methods: {
        option(data, type) {
            userService.option({
                    mid: data.mid, 
                    name: data.name, 
                    option: type, 
                    role: data.role
                }, 
                this.tableData
            )
        },
        changePage(page) {
            userService.userList({page: page, pageSize: this.tableData.pageSize}, this.tableData)
            // TODO: 对应页码更新
        },
    }
}
</script>

<style>
.member-list .el-row {
    text-align: left;
}
.member-list .title {
    font-size: 140%;
}
</style>
