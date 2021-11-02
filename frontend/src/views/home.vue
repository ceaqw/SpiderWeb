<template>
    <div class="home">
        <!-- 筛选条件布局 -->
        <Filter /> 
        <el-row>
            <line-chart :chart-data="lineChartData" />
        </el-row>
        <el-row>
            <el-col :xs="24" :sm="24" :lg="8">
                <div class="chart-wrapper">
                    <raddar-chart />
                </div>
            </el-col>
            <el-col :xs="24" :sm="24" :lg="8">
                <div class="chart-wrapper">
                    <pie-chart />
                </div>
            </el-col>
            <el-col :xs="24" :sm="24" :lg="8">
                <div class="chart-wrapper">
                    <bar-chart />
                </div>
            </el-col>
        </el-row>
  </div>
</template>

<script>
import theme from '@/conf/theme'
import Filter from '@/components/filter'

import LineChart from './dashboard/LineChart'
import RaddarChart from './dashboard/RaddarChart'
import PieChart from './dashboard/PieChart'
import BarChart from './dashboard/BarChart'

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
        LineChart,
        RaddarChart,
        PieChart,
        BarChart
    },
    data() {
        return {
            lineChartData: lineChartData.newVisitis,
            theme
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

.home .el-row {
    padding: 10px;
    margin-bottom: 15px;
    border-radius: 10px;
    box-shadow: 0px 3px 7px #D3DCE6;
    background-color: #F2F5F5;
}
</style>
