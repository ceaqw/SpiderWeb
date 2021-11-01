<template>
    <div>
        <el-table :data="tableData" stripe style="width: 100%">
            <el-table-column prop="mid" label="Mid"> </el-table-column>
            <el-table-column prop="name" label="Name"> </el-table-column>
            <el-table-column label="Status">
                <template #default="scope">
                    <span style="margin-left: 10px">{{ memberStatus[scope.row.status] }}</span>
                </template>
            </el-table-column>
            <el-table-column label="Options">
                <template #default="scope">
                    <el-button type="warning" v-if="scope.row.status == 1" icon="el-icon-refresh-left" size="medium" title="恢复" @click.prevent="option(scope.row.mid, 'recover')"></el-button>
                    <el-button type="danger" v-else-if="scope.row.status == 0" icon="el-icon-delete" size="medium" title="删除" @click.prevent="option(scope.row.mid, 'del')"></el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination background layout="prev, pager, next" :total="totalNumber" @current-change="changePage">
        </el-pagination>
    </div>
</template>

<script>
export default {
    data() {
        return {
            totalNumber: 500,
            memberStatus: [
                '正常', '已失效', 
            ],
            tableData: [
                {
                    mid: '1000',
                    name: 'cai.hongliang',
                    status: 1,
                },
                {
                    mid: '1001',
                    name: 'caicaicai',
                    status: 1,
                },
                {
                    mid: '1002',
                    name: 'feng.teng',
                    status: 0,
                },
                {
                    mid: '1003',
                    name: 'wang.chengyao',
                    status: 1,
                },
            ],
        }
    },
    methods: {
        option(mid, type) {
            alert(mid + ' ' + type)
            // TODO: 对应用户操作
        },
        changePage(page) {
            alert(page)
            // TODO: 对应页码更新
        }
    }
}
</script>

<style>

</style>
