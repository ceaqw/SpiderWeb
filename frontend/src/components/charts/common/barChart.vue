<template>
  <div :class="className" :style="{height:height, width:width}" />
</template>

<script>
import * as echarts from 'echarts'
// import theme from '@/conf/theme'

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
        options: {
            type: Object,
            default: {
                tooltip: {
                    trigger: 'item',
                    backgroundColor: 'rgba(255,255,255,0.8)'
                },
                legend: {},
                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: [{type: 'category', data: []}],
                yAxis: [{type: 'value'}],
                series: 
                [
                    {
                        name: 'Success',
                        type: 'bar',
                        stack: 'Task',
                        emphasis: {
                            focus: 'series'
                        },
                        data: []
                    },
                    {
                        name: 'Fail',
                        type: 'bar',
                        stack: 'Task',
                        emphasis: {
                            focus: 'series'
                        },
                        data: []
                    },
                    {
                        name: 'Undone',
                        type: 'bar',
                        stack: 'Task',
                        emphasis: {
                            focus: 'series'
                        },
                        data: []
                    }
                ]
            }
        },
        chartDatas: {
            type: Function,
            required: true
        }
    },
    data() {
        return {
            chart: null,
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.initChart()
            window.addEventListener('resize', this.resize)
        })
    },
    beforeDestroy() {
        if (!this.chart) {
            return
        }
        this.chart.dispose()
        this.chart = null
        window.removeEventListener('resize', this.resize)
    },
    methods: {
        initChart() {
            this.chart = echarts.init(this.$el)
            this.chartDatas(this.options, 'bar')
            this.chart.setOption(this.options)
        },
        resize() {
            this.chart.resize()
        },
        refreshDatas() {
            this.chartDatas(this.options, 'bar')
            this.chart.setOption(this.options)
        }
    }
}
</script>
