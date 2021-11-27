<template>
    <div id="project">
        <Filter :filter="filterForm" parentName="project"/> 
        <ProjectChart 
            :filter="filterForm" 
            v-for="item, index in paginationData.data" 
            :key="index"
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
        console.log(this.$store.state.flushQueue['project'])
    },
    data() {
        return {
            filterForm: this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.$store.state.projectFilter,
            paginationData: {
                data: this.$store.state.projectList['all'].slice(0, pageSize),
                totalNumber: this.$store.state.projectList['all'].length,
                pageSize: pageSize
            },
        }
    },
    methods: {
        changePage(page) {
            this.paginationData.data = this.$store.state.projectList['all'].slice((page-1)*this.$conf.pageSize, page*this.$conf.pageSize)
        },
    }
}
</script>
