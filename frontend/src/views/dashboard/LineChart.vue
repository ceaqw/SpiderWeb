<template>
  <div :class="className" :style="{height:height, width:width}" />
</template>

<script>
import * as echarts from 'echarts'
// require('echarts/theme/macarons') // echarts theme

export default {
    props: {
        className: {
            type: String,
            default: 'chart'
        },
        width: {
            type: String,
            default: '100%'
        },
        height: {
            type: String,
            default: '350px'
        },
        autoResize: {
            type: Boolean,
            default: true
        },
        chartData: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            chart: null
        }
    },
    watch: {
        chartData: {
            deep: true,
            handler(val) {
                this.setOptions(val)
            }
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.initChart()
        })
    },
    beforeDestroy() {
        if (!this.chart) {
        return
        }
        this.chart.dispose()
        this.chart = null
    },
    methods: {
        initChart() {
            this.chart = echarts.init(this.$el)
            this.chart.setOption({
                grid: {
                    top: 30,
                    bottom: 30
                },
                xAxis: {
                    type: 'category',
                    data: this.chartData.expectedData
                },
                yAxis: {
                    type: 'value'
                },
                series: [
                    {
                        data: this.chartData.actualData,
                        type: 'line',
                        smooth: true
                    }
                ]
            })
        }
    }
}
</script>
