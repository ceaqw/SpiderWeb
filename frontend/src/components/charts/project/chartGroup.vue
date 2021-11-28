<template>
    <div id="charts-project-chartGroup">
        <div class="row">
            <div class="title">{{ project }}</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="8">
                    <PieChart 
                        :chartDatas="this.chartDatas.projectChartDatas" 
                        :filter="filterForm" 
                        :type_="'pie-left-' + project" 
                        :Id="'project-pie-left-' + project"
                        parentName="project"
                    />
                </el-col>
                <el-col :span="1">
                    <el-divider direction="vertical" style="height: 90%;"></el-divider>
                </el-col>
                <el-col :span="11">
                    <PieChart 
                        :chartDatas="this.chartDatas.projectChartDatas" 
                        :filter="filterForm" 
                        :type_="'pie-right-' + project" 
                        :Id="'project-pie-right-' + project"
                        parentName="project"
                    />
                </el-col>
                <el-col :span="4">
                    kpi data
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
    mounted() {
        console.log(this.project)
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
#charts-project-chartGroup .title {
    font-size: 130%;
    text-align: left;
    margin: 10px 5px 0px;
    padding-top: 10px;
    font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB',
    'Microsoft YaHei', '微软雅黑', Arial, sans-serif;
}

#charts-project-chartGroup .row{
    padding: 0 20px;
    margin-bottom: 15px;
    border-radius: 10px;
    box-shadow: 0px 3px 7px #D3DCE6;
    background-color: #F2F5F5;
    text-align: center;
}
</style>