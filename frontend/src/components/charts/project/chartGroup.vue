<template>
    <div id="charts-project-chartGroup">
        <div class="row">
            <div class="title">总览图表</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="10">
                    <PieChart 
                        :chartDatas="this.chartDatas.projectChartDatas" 
                        :filter="filterForm" 
                        :type_="'pie-left-' + project" 
                        parentName="project"
                    />
                </el-col>
                <el-col :span="14">
                    <PieChart 
                        :chartDatas="this.chartDatas.projectChartDatas" 
                        :filter="filterForm" 
                        :type_="'pie-right-' + project" 
                        parentName="project"
                    />
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
import PieChart from '@/components/charts/common/pieChart'

import { projectChartDatas } from '../../../service/datas/projectChart'
export default {
    props: {
        filter: {
            type: Object,
            required: true
        },
        project: {
            type: String,
            required: true
        }
    },
    components: {
        PieChart,
    },
    data() {
        return {
            chartDatas: {
                projectChartDatas
            },
            filterForm: this.filter,
        }
    },
}
</script>

<style>

</style>