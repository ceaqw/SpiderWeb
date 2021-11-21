<template>
    <div class="home">
        <!-- 筛选条件布局 -->
        <Filter :filter="filterForm" parentName="dashboard"/> 
        <div class="row">
            <div class="title">总览图表</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="10">
                    <PieChart :chartDatas="this.chartDatas.allChartDatas" :filter="filterForm" parentName="dashboard"/>
                </el-col>
                <!-- <el-col :span="1">
                    <el-divider direction="vertical" style="height: 100%"></el-divider>
                </el-col> -->
                <el-col :span="14">
                    <BarChart :chartDatas="this.chartDatas.allChartDatas" :filter="filterForm" parentName="dashboard"/>
                </el-col>
            </el-row>
        </div>
        <div class="row">
            <div class="title">Analyse监控</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="10">
                    <PieChart :chartDatas="this.chartDatas.analyseDatas" :filter="filterForm" parentName="dashboard"/>
                </el-col>
                <!-- <el-col :span="1">
                    <el-divider direction="vertical" style="height: 100%"></el-divider>
                </el-col> -->
                <el-col :span="14">
                    <BarChart :chartDatas="this.chartDatas.analyseDatas" :filter="filterForm" parentName="dashboard"/>
                </el-col>
            </el-row>
        </div>
        <el-row>
            <el-col :span="8">
                <div class="row">
                    <div class="title">Rakuten</div>
                    <el-divider></el-divider>
                    <el-row>
                        <PieChart :chartDatas="this.chartDatas.rakutenDatas" :filter="filterForm" parentName="dashboard"/>
                    </el-row>
                    <el-divider></el-divider>
                    <el-row>
                        <BarChart :chartDatas="this.chartDatas.rakutenDatas" :filter="filterForm" parentName="dashboard"/>
                    </el-row>
                </div>
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="7">
                <div class="row">
                    <div class="title">Yahoo</div>
                    <el-divider></el-divider>
                    <el-row>
                        <PieChart :chartDatas="this.chartDatas.yahooDatas" :filter="filterForm" parentName="dashboard"/>
                    </el-row>
                    <el-divider></el-divider>
                    <el-row>
                        <BarChart :chartDatas="this.chartDatas.yahooDatas" :filter="filterForm" parentName="dashboard"/>
                    </el-row>
                </div>
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="7">
                <div class="row">
                    <div class="title">Amazon</div>
                    <el-divider></el-divider>
                    <el-row>
                        <PieChart :chartDatas="this.chartDatas.amazonDatas"  :filter="filterForm" parentName="dashboard"/>
                    </el-row>
                    <el-divider></el-divider>
                    <el-row>
                        <BarChart :chartDatas="this.chartDatas.amazonDatas" :filter="filterForm" parentName="dashboard"/>
                    </el-row>
                </div>
            </el-col>
        </el-row>
        <div class="row">
            <div class="title">TOP ERROR</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="9">
                    <PieChart :filter="filterForm" cacheChartKey="topError" parentName="dashboard"/>
                </el-col>
                <el-col :span="15">
                    <!-- <PieChart :chartDatas="this.chartDatas.topErrorDatas" :filter="filterForm" parentName="dashboard"/> -->
                    <ErrorTable :tableDatas="this.chartDatas.topErrorDatas" :filter="filterForm" parentName="dashboard"/>
                </el-col>
            </el-row>
        </div>
        <div class="row">
            <div class="title">关键 KPI</div>
            <el-divider></el-divider>
            <el-table :data="kpiDatas.datas" stripe style="width: 100%">
                <el-table-column prop="project" label="项目"> </el-table-column>
                <el-table-column prop="kpi" label="Kpi"> </el-table-column>
                <el-table-column label="占比">
                    <template #default="scope">
                        <span>{{ scope.row.rate }}%</span>
                    </template>
                </el-table-column>
                <el-table-column label="查看">
                    <template #default="scope">
                        <el-button icon="el-icon-view" size="medium" type="primary" title="查看" @click.prevent="view(scope.row)"></el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
  </div>
</template>

<script>
import theme from '@/conf/theme'
import Filter from '@/components/filter'

import PieChart from '@/components/charts/common/pieChart'
import BarChart from '@/components/charts/common/barChart'
import ErrorTable from '@/components/tables/errorTable.vue'

import { 
        allChartDatas, analyseDatas, rakutenDatas, yahooDatas, amazonDatas, topErrorDatas
    } from '@/service/datas/chart'

export default {
    beforeRouteLeave(to, form, next) {
        // 清空刷新队列
        if (this.$store.state.flushQueue['dashboard'].length) {
            this.$store.state.flushQueue['dashboard'] = []
        }
        next()
    },
    components: {
        Filter,
        PieChart,
        BarChart,
        ErrorTable
    },
    data() {
        return {
            theme,
            kpiDatas: {
                datas: [
                    {project: 'item_category', kpi: 'content_empty', rate: '55.37'},
                    {project: 'lohaco_item', kpi: 'single_tax', rate: '100'},
                    {project: 'lohaco_item', kpi: '	single_price', rate: '100'},
                    {project: 'lohaco_item', kpi: 'coupon', rate: '100'},
                    {project: 'lohaco_item', kpi: 'rank_cid', rate: '100'},
                ]
            },
            chartDatas: {
                allChartDatas,
                analyseDatas,
                rakutenDatas,
                yahooDatas,
                amazonDatas,
                topErrorDatas
            },
            filterForm: this.$store.state.FilterSharing ? this.$store.state.shareFilter : this.$store.state.dashboardFilter,
        }
    },
    mouted() {
        // 清空刷新队列
        if (this.$store.state.flushQueue['dashboard'].length) {
            this.$store.state.flushQueue['dashboard'] = []
        }
    },
    methods: {
        view(rowData) {
            alert(rowData.project)
        }
  }
}
</script>

<style lang="scss" scoped>
.home {
    width: auto;
    position: relative;

    .chart-wrapper {
        // background: #fff;
        // padding: 16px 16px 0;
        margin-bottom: 32px;
    }
}

.home .row{
    padding: 0 20px;
    margin-bottom: 15px;
    border-radius: 10px;
    box-shadow: 0px 3px 7px #D3DCE6;
    background-color: #F2F5F5;
    text-align: center;
}

.home .title {
    font-size: 130%;
    text-align: left;
    margin: 10px 5px 5px;
    font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB',
    'Microsoft YaHei', '微软雅黑', Arial, sans-serif;
}

.home .el-divider {
    margin: 5px;
}
</style>
