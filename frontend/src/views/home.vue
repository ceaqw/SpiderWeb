<template>
    <div class="home">
        <!-- 筛选条件布局 -->
        <Filter /> 
        <div class="row">
            <div class="title">总览图表</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="10">
                    <PieChart />
                </el-col>
                <!-- <el-col :span="1">
                    <el-divider direction="vertical" style="height: 100%"></el-divider>
                </el-col> -->
                <el-col :span="14">
                    <BarChart />
                </el-col>
            </el-row>
        </div>
        <div class="row">
            <div class="title">Analyse监控</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="10">
                    <PieChart />
                </el-col>
                <!-- <el-col :span="1">
                    <el-divider direction="vertical" style="height: 100%"></el-divider>
                </el-col> -->
                <el-col :span="14">
                    <BarChart />
                </el-col>
            </el-row>
        </div>
        <el-row>
            <el-col :span="8">
                <div class="row">
                    <div class="title">Rakuten</div>
                    <el-divider></el-divider>
                    <el-row>
                        <PieChart />
                    </el-row>
                    <el-divider></el-divider>
                    <el-row>
                        <BarChart />
                    </el-row>
                </div>
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="7">
                <div class="row">
                    <div class="title">Yahoo</div>
                    <el-divider></el-divider>
                    <el-row>
                        <PieChart />
                    </el-row>
                    <el-divider></el-divider>
                    <el-row>
                        <BarChart />
                    </el-row>
                </div>
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="7">
                <div class="row">
                    <div class="title">Amazon</div>
                    <el-divider></el-divider>
                    <el-row>
                        <PieChart />
                    </el-row>
                    <el-divider></el-divider>
                    <el-row>
                        <BarChart />
                    </el-row>
                </div>
            </el-col>
        </el-row>
        <div class="row">
            <div class="title">TOP ERROR</div>
            <el-divider></el-divider>
            <el-row>
                <el-col :span="10">
                    <PieChart />
                </el-col>
                <el-col :span="14"></el-col>
            </el-row>
        </div>
  </div>
</template>

<script>
import theme from '@/conf/theme'
import Filter from '@/components/filter'

import PieChart from '@/components/charts/common/pieChart'
import BarChart from '@/components/charts/common/barChart'

const lineChartData = {
  newVisitis: {
    expectedData: [100, 120, 161, 134, 105, 160, 165],
    actualData: [120, 82, 91, 154, 162, 140, 145]
  },
  messages: {
    expectedData: [200, 192, 120, 144, 160, 130, 140],
    actualData: [180, 160, 151, 106, 145, 150, 130]
  },
  purchases: {
    expectedData: [80, 100, 121, 104, 105, 90, 100],
    actualData: [120, 90, 100, 138, 142, 130, 130]
  },
  shoppings: {
    expectedData: [130, 140, 141, 142, 145, 150, 160],
    actualData: [120, 82, 91, 154, 162, 140, 130]
  }
}

export default {
    components: {
        Filter,
        PieChart,
        BarChart,
    },
    data() {
        return {
            lineChartData: lineChartData.newVisitis,
            theme,
            commonInfos: {
                title: ['all_platform', 'all_project']
            }
        }
    },
    methods: {
        handleSetLineChartData(type) {
        this.lineChartData = lineChartData[type]
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
