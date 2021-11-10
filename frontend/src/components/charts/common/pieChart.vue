<template>
  <div :class="className" :style="{height:height, width:width}" />
</template>

<script>
import * as echarts from 'echarts'
import theme from '@/conf/theme'

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
                title: {
                    text: '',
                    left: 'center',
                    textStyle: {
                        color: theme.color.black
                    }
                },
                tooltip: {
                    trigger: 'item',
                    backgroundColor: 'rgba(255,255,255,0.8)'
                },
                legend: {
                    orient: 'vertical',
                    left: 'left'
                },
                series: [
                    {
                        type: 'pie',
                        radius: ['40%', '70%'],
                        avoidLabelOverlap: false,
                        startAngle: 0,
                        endAngle: 180,
                        label: {
                            show: false,
                            position: 'center'
                        },
                        labelLine: {
                            show: false
                        },
                        emphasis: {
                            label: {
                                show: true,
                                fontSize: '40',
                                fontWeight: 'bold'
                            }
                        },
                        data: []
                    }
                ]
            }
        },
        chartDatas: {
            type: Function
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
            // window.addEventListener('resize', this.resize)
        })
    },
    beforeDestroy() {
        if (!this.chart) {
            return
        }
        this.chart.dispose()
        this.chart = null
        // window.removeEventListener('resize', this.resize)
    },
    methods: {
        initChart() {
            this.chart = echarts.init(this.$el)
            this.chartDatas(this.options, 'pie')
            this.chart.setOption(this.options)
        },
        resize() {
            this.chart.resize()
        },
        refreshDatas() {
            this.chartDatas(this.options, 'pie')
            this.chart.setOption(this.options)
        }
    }
}
</script>
