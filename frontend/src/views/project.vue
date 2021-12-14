<!--
 * @Date: 2021-11-29 09:13:10
 * @LastEditTime: 2021-12-14 16:41:09
 * @Author: ceaqw
-->
<template>
    <div id="project">
        <Filter :filter="filterForm" parentName="project"/> 
        <ProjectChart 
            :filter="filterForm" 
            v-for="item in paginationData.data" 
            :key="item+Date.parse(new Date())"
            :project="item"
        />
        <el-pagination background layout="prev, pager, next" :total="paginationData.totalNumber" :page-size="paginationData.pageSize" @current-change="changePage">
    </el-pagination>
    </div>
</template>

<script>
import Filter from '@/components/filter.vue'
import ProjectChart from '../components/charts/project/chartGroup.vue'
import { getFilterData } from '../service/datas/base'
import baseConf from '../conf/baseConf'

// 模块有关pageSize操作均+1，避开无效参数all
let pageSize = baseConf.pageSize/2

export default {
    beforeRouteLeave(to, form, next) {
        // 清空刷新队列
        if (this.$store.state.flushQueue['project'].length) {
            this.$store.state.flushQueue['project'] = []
        }
        next()
    },
    components: {
        Filter,
        ProjectChart
    },
    beforeMount() {
        getFilterData()
    },
    mounted() {
        this.initPagation()
    },
    data() {
        return {
            filterForm: this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.$store.state.projectFilter,
            paginationData: {
                data: [],
                totalNumber: 0,
                pageSize: pageSize
            },
        }
    },
    watch: {
        // 监听平台变化
        'filterForm.platForm': function () {
            // console.log(this.$store.state.flushQueue['project'])
            this.$store.state.flushQueue['project'] = {}
            this.paginationData.data = this.$store.state.projectList[this.filterForm.platForm].slice(1, pageSize)
            this.paginationData.totalNumber = this.$store.state.projectList[this.filterForm.platForm].length - 1
        },
        // 监听project变化
        'filterForm.project': function () {
            // console.log(this.$store.state.flushQueue['project'])
            this.$store.state.flushQueue['project'] = {}
            if (this.filterForm.project != 'all') {
                this.paginationData.data = [this.filterForm.project]
                this.paginationData.totalNumber = 1
            } else {
                this.initPagation()
            }
        }
    },
    methods: {
        changePage(page) {
            this.paginationData.data = this.$store.state.projectList[this.filterForm.platForm].slice((page-1)*pageSize+1, page*pageSize)
        },
        initPagation() {
            this.paginationData.data = this.$store.state.projectList[this.filterForm.platForm].slice(1, pageSize)
            this.paginationData.totalNumber = this.$store.state.projectList[this.filterForm.platForm].length - 1
        },
    }
}
</script>
