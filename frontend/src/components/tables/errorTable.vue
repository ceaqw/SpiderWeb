
<template>
<div class="errorTable">
    <el-table
        :data="tableData.data"
        size="mini"
        :highlight-current-row="true"
        @cell-click="rowClick"
        @sort-change="sortChange"
    >
        <el-table-column prop="rank" label="Id" :formatter="formatter" width="40%"></el-table-column>
        <el-table-column prop="project" label="Project Name" width="140%"></el-table-column>
        <el-table-column prop="platform" label="Platform" ></el-table-column>
        <el-table-column prop="fullName" label="Fulla Name" width="180%"></el-table-column>
        <el-table-column prop="fail" label="Fail" sortable="custom" width="80%"></el-table-column>
        <el-table-column prop="undone" label="Undone" sortable="custom"></el-table-column>
        <el-table-column prop="success" label="Success" sortable="custom"></el-table-column>
    </el-table>
    <el-pagination background layout="prev, pager, next" :total="tableData.totalNumber" @current-change="changePage" :page-size="tableData.pageSize">
    </el-pagination>
</div>
</template>

<script>
import { topError } from '@/service/datas/localChart'
 
export default {
    props: {
        Id: {
            type: String,
            required: true
        },
        parentName: {type: String},
        tableDatas: {
            type: Function
        },
        filter: {
            type: Object
        }
    },
    data() {
        return {
            tableData: {
                data: [],
                pageSize: this.$conf.pageSize/2,
                totalNumber: 0,
            },
            cacheData: [],
            currentPage: 1
        }
    },
    mounted() {
        new Promise(()=>{this.initChart()})
    },
    methods: {
        formatter(row, column, val, index) {
            return index + 1
        },
        changePage(page) {
            this.currentPage = page
            const pageSize = this.$conf.pageSize/2
            this.tableData.data = this.cacheData.slice(pageSize*(page-1), pageSize*page)
        },
        initChart() {
            this.tableDatas(this.tableData, this.filter, this.cacheData)
            this.$store.state.flushQueue[this.parentName][this.Id] = this.refreshDatas
        },
        refreshDatas() {
            this.tableDatas(this.tableData, this.filter, this.cacheData)
        },
        rowClick(row) {
            topError(row)
        },
        sortChange(params) {
            const sortKey = params.prop
            if (!sortKey) return
            if (params.order == 'ascending') {
                this.cacheData.sort((a, b)=>{return a[sortKey] - b[sortKey]})
            } else if (params.order == 'descending') {
                this.cacheData.sort((a, b)=>{return b[sortKey] - a[sortKey]})
            }
            this.changePage(this.currentPage)
        }
    },
}
</script>

<style>

</style>
