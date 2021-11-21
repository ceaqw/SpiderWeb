<template>
    <div id="project">
        <Filter :filter="filterForm"/> 
        <h2>Project</h2>
        <el-pagination background layout="prev, pager, next" :total="paginationData.totalNumber" :page-size="paginationData.pageSize">
    </el-pagination>
    </div>
</template>

<script>
import Filter from '@/components/filter.vue'
import { getFilterData } from '@/service/datas/base.js'

export default {
    components: {
        Filter,
    },
    beforeMount() {
        getFilterData()
    },
    mounted() {

    },
    data() {
        return {
            filterForm: this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.$store.state.projectFilter,
            paginationData: {
                data: [],
                totalNumber: this.$store.state.projectList['all'].length,
                pageSize: this.$conf.pageSize
            }
        }
    },
}
</script>
