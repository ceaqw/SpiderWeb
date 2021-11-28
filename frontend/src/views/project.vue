<template>
    <div id="project">
        <Filter :filter="filterForm" parentName="project"/> 
        <ProjectChart 
            :filter="filterForm" 
            v-for="item in paginationData.data" 
            :key="item"
            :project="item"
        />
        <el-pagination background layout="prev, pager, next" :total="paginationData.totalNumber" :page-size="paginationData.pageSize" @current-change="changePage">
    </el-pagination>
    </div>
</template>

<script>
import Filter from '@/components/filter.vue'
import ProjectChart from '../components/charts/project/chartGroup.vue'
import { getFilterData } from '@/service/datas/base.js'
import baseConf from '../conf/baseConf'

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
            console.log(this.$store.state.flushQueue['project'])
            this.$store.state.flushQueue['project'] = {}
            this.paginationData.data = this.$store.state.projectList[this.filterForm.platForm].slice(0, pageSize)
            this.paginationData.totalNumber = this.$store.state.projectList[this.filterForm.platForm].length
        },
        // 监听project变化
        'filterForm.project': function () {
            console.log(this.$store.state.flushQueue['project'])
            if (this.filterForm.project != 'all') {
                this.$store.state.flushQueue['project'] = {}
                this.paginationData.data = [this.filterForm.project]
                this.paginationData.totalNumber = 1
            }
        }
    },
    methods: {
        changePage(page) {
            this.paginationData.data = this.$store.state.projectList['all'].slice((page-1)*pageSize, page*pageSize)
        },
        initPagation() {
            this.paginationData.data = this.$store.state.projectList[this.filterForm.platForm].slice(0, pageSize)
            this.paginationData.totalNumber = this.$store.state.projectList[this.filterForm.platForm].length
        },
    }
}
</script>
